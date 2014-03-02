package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regifted/data"
)

type Pes struct {
	byteChunk []byte

	pid          uint
	streamType   uint
	streamId     uint
	packetLength uint
	flags        uint
	pts          uint
	dts          uint
	payload      []byte
	nal          Nal
}

type Nal struct {
}

type Pat struct {
	SKIP_BYTES int
	CRC_SIZE   int

	byteChunk []byte

	unitStart bool

	pointerField           bool
	tableId                uint
	flags                  uint
	sectionSyntaxIndicator bool
	sectionLength          uint
	transportStreamId      uint
	versionNumber          uint
	currentNext            uint
	sectionNumber          uint
	lastSectionNumber      uint
	count                  uint

	programs []Program
}

type TsPacket struct {
	byteChunk []byte

	sync  uint
	flags uint

	payload []byte

	transportError bool
	unitStart      bool
	priority       bool
	pid            uint
	scramble       uint
	hasAdaptation  bool
	hasPayload     bool
	continuity     uint
	adaptation     Adaptation
}

type Adaptation struct {
	payload   []byte
	byteChunk []byte

	size     uint
	pcr      Pcr
	opcr     Pcr
	splice   uint
	stuffing int

	discontinuity bool
	random        bool
	priority      bool
	hasPCR        bool
	hasOPCR       bool
	hasSplice     bool
	hasPrivate    bool
	hasExtension  bool
}

type Pcr struct {
	pcr       uint
	ext       uint
	byteChunk []byte

	a  uint
	b  uint
	c  uint
	d  uint
	ef uint
	e  uint
	f  uint
	g  uint
}

type Pmt struct {
	byteChunk              []byte
	pointerField           bool
	unitStart              bool
	tableId                uint
	sectionSyntaxIndicator bool
	sectionLength          uint
	programNumber          uint
	versionNumber          uint
	sectionNumber          uint
	lastSectionNumber      uint

	pcrPid            uint
	programInfoLength uint
	descriptor        []byte
	count             uint
	entries           []PmtEntry

	types map[uint32]uint
}

type PmtEntry struct {
	byteChunk []byte

	streamType uint
	pid        uint
	infoLength uint
	descriptor []byte
}

type ElementaryStreamPacket struct {
	byteChunk []byte
	payload   []byte

	unitStart bool

	pid           uint
	hasAdaptation bool
}

type Writer struct {
	chunks []byte
	size   int
}

type Program struct {
	byteChunk []byte

	pid    uint
	number uint
}

var pesCollector map[uint]Pes
var pmtConstructors map[uint]Pmt
var entryConstructors map[uint]PmtEntry
var elementaryConstructors map[uint]ElementaryStreamPacket
var types map[uint]uint
var pat Pat

func main() {

	fileName := os.Args[1]

	fmt.Printf("Attempting to read file, Run 7" + fileName + "\n")

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	reader := data.NewReader(bytes)

	pat = Pat{}
	pat.tableId = 0

	Init()

	fmt.Println("Size: ", len(bytes))

	for reader.Cursor < uint64(len(bytes)) {

		byteChunk := reader.ReadBytes(188)

		tsPacket := TsPacket{}

		tsPacket.byteChunk = byteChunk

		tsPacket.Read()

	}

	for key := range pesCollector {

		CreateAndDispensePes(key, types[key])

	}

}

//CreateAndDispensePes
//Dump the remaining PES
func CreateAndDispensePes(pid uint, streamType uint) {

	pes := pesCollector[pid]

	pes.pid = pid

	pes.streamType = streamType

	pes.Read()

	pes.Print()

}

//Init
//Initialize the constructors
func Init() {
	pmtConstructors = make(map[uint]Pmt)
	entryConstructors = make(map[uint]PmtEntry)
	types = make(map[uint]uint)
	pesCollector = make(map[uint]Pes)
	elementaryConstructors = make(map[uint]ElementaryStreamPacket)
}

//TsPacket Read
//TsPacket
//The first read of every packet
//flags:
//asserts 'G' - 8 bits
//transportError - Set by demodulator if can't correct errors in the stream, to tell the demultiplexer that the packet has an uncorrectable error (1 bit)
//unitStart - Boolean flag with a value of true meaning the start of PES data or PSI otherwise zero only. (1 bit)
//priority - Boolean flag with a value of true meaning the current packet has a higher priority than other packets with the same PID. (1 bit)
//pid - Packet Identifier (13 bit)
//scramble - '00' = Not scrambled, '01' = Reserved for future use, '10' = Scrambled with even key, '11' = Scrambled with odd key (2 bits)
//hasAdaptation - If adaption field exist value is true (0 or more bits)
//hasPayload - 	If contains payload value is true (0 or more bits)
//continuity - Sequence number of payload packets, Incremented only when a payload is present (i.e., payload value is true) (4 bits)
func (tsPacket *TsPacket) Read() {

	var flags uint = 0

	reader := data.NewReader(tsPacket.byteChunk)

	tsPacket.sync = reader.Read(1)

	if tsPacket.sync == 71 {

		flags = reader.Read(2)

		tsPacket.transportError = flags&0x8000 > 0
		tsPacket.unitStart = flags&0x4000 > 0
		tsPacket.priority = flags&0x2000 > 0
		tsPacket.pid = flags & 0x1fff
		fmt.Println("pid", tsPacket.pid)

		flags = reader.Read(1)

		tsPacket.scramble = flags & 0xc0 >> 6
		tsPacket.hasAdaptation = flags&0x20 > 0
		tsPacket.hasPayload = flags&0x10 > 0
		tsPacket.continuity = flags & 0x0f

		tsPacket.Print()

		if tsPacket.hasAdaptation {
			tsPacket.adaptation.byteChunk = reader.ReadBytes(reader.Size - reader.Cursor)
			tsPacket.adaptation.Read()
		}

		if tsPacket.pid == 0 {
			pat.byteChunk = reader.ReadBytes(reader.Size - reader.Cursor)

			pat.unitStart = tsPacket.unitStart
			pat.Read()
		}

		if pmt, ok := pmtConstructors[tsPacket.pid]; ok {
			pmt.unitStart = tsPacket.unitStart
			pmt.byteChunk = reader.ReadBytes(reader.Size - reader.Cursor)
			pmt.Read()
		}

		if elementaryStreamPacket, ok := elementaryConstructors[tsPacket.pid]; ok {

			elementaryStreamPacket.pid = tsPacket.pid
			elementaryStreamPacket.unitStart = tsPacket.unitStart

			if tsPacket.hasAdaptation {
				elementaryStreamPacket.payload = tsPacket.adaptation.payload
			} else {
				elementaryStreamPacket.payload = reader.ReadBytes(reader.Size - reader.Cursor)
			}

			elementaryStreamPacket.Dispatch()
			elementaryStreamPacket.Print()
		}

	}

}

//Pat Read
//table_id – This is an 8-bit field, which shall be set to 0x00 a
//
//sectionSyntaxIndicator – The section_syntax_indicator is a 1-bit field which shall be set to '1'.
//
//sectionLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the section, starting immediately following the section_length field, and including the CRC. The value in this
//field shall not exceed 1021 (0x3FD).
//
//transportStreamId – This is a 16-bit field which serves as a label to identify this Transport Stream from any other
//multiplex within a network. Its value is defined by the user.
//
//versionNumber – This 5-bit field is the version number of the whole Program Association Table. The version number
//shall be incremented by 1 modulo 32 whenever the definition of the Program Association Table changes. When the
//current_next_indicator is set to '1', then the version_number shall be that of the currently applicable Program Association
//Table. When the current_next_indicator is set to '0', then the version_number shall be that of the next applicable Program
//Association Table.
//
//currentNext – A 1-bit indicator, which when set to '1' indicates that the Program Association Table sent is
//currently applicable. When the bit is set to '0', it indicates that the table sent is not yet applicable and shall be the next
//table to become valid.
//
//sectionNumber – This 8-bit field gives the number of this section. The section_number of the first section in the
//Program Association Table shall be 0x00. It shall be incremented by 1 with each additional section in the Program
//Association Table.
//
//lastSectionNumber – This 8-bit field specifies the number of the last section (that is, the section with the highest
//section_number) of the complete Program Association Table.
func (pat *Pat) Read() {

	var SKIP_BYTES uint = 5
	var CRC_SIZE uint = 4
	var PROGRAM_SIZE uint = 4
	var flags uint = 0

	var flag bool = false

	reader := data.NewReader(pat.byteChunk)

	if reader.Read(1) == 1 {
		flag = true
	}

	pat.pointerField = (pat.unitStart && flag) || false

	pat.tableId = reader.Read(1)

	flags = reader.Read(2)

	pat.sectionSyntaxIndicator = flags&0x8000 > 0

	pat.sectionLength = flags & 0x3ff

	pat.transportStreamId = reader.Read(2)

	flags = reader.Read(1)

	pat.versionNumber = flags & 0x3ffe
	pat.currentNext = flags & 0x0001

	pat.sectionNumber = reader.Read(1)

	pat.lastSectionNumber = reader.Read(1)

	pat.count = pat.sectionLength - SKIP_BYTES

	pat.Print()

	for pat.count > CRC_SIZE {
		program := Program{}
		pmt := Pmt{}

		program.Read(reader)

		pat.programs = append(pat.programs, program)
		pmtConstructors[program.pid] = pmt

		pat.count = pat.count - PROGRAM_SIZE
	}

}

//Pmt Read
//tableId – This is an 8-bit field, which in the case of a TS_program_map_section shall be always set to 0x02
//
//sectionSyntaxIndicator – The section_syntax_indicator is a 1-bit field which shall be set to '1'.
//
//sectionLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the section starting immediately following the section_length field, and including the CRC. The value in this
//field shall not exceed 1021 (0x3FD).
//
//programNumber – program_number is a 16-bit field. It specifies the program to which the program_map_PID is
//applicable. One program definition shall be carried within only one TS_program_map_section. This implies that a
//program definition is never longer than 1016 (0x3F8). See Informative Annex C for ways to deal with the cases when
//that length is not sufficient. The program_number may be used as a designation for a broadcast channel, for example. By
//describing the different program elements belonging to a program, data from different sources (e.g. sequential events)
//can be concatenated together to form a continuous set of streams using a program_number.
//
//versionNumber – This 5-bit field is the version number of the TS_program_map_section. The version number shall be
//incremented by 1 modulo 32 when a change in the information carried within the section occurs. Version number refers
//to the definition of a single program, and therefore to a single section. When the current_next_indicator is set to '1', then
//the version_number shall be that of the currently applicable TS_program_map_section. When the current_next_indicator
//is set to '0', then the version_number shall be that of the next applicable TS_program_map_section.
//
//sectionNumber – The value of this 8-bit field shall be 0x00.
//
//lastSectionNumber – The value of this 8-bit field shall be 0x00.
//
//pcrPid – This is a 13-bit field indicating the PID of the Transport Stream packets which shall contain the PCR fields
//valid for the program specified by program_number. If no PCR is associated with a program definition for private
//streams, then this field shall take the value of 0x1FFF.
//
//programInfoLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the
//number of bytes of the descriptors immediately following the program_info_length field.
func (pmt *Pmt) Read() {

	var CRC_SIZE uint = 4
	var SKIP_BYTES uint = 9
	var flags uint = 0

	var flag bool = false

	reader := data.NewReader(pmt.byteChunk)

	if reader.Read(1) == 1 {
		flag = true
	}

	pmt.pointerField = (pmt.unitStart && flag) || false

	pmt.tableId = reader.Read(1)

	flags = reader.Read(2)

	pmt.sectionSyntaxIndicator = flags&0x8000 > 0
	pmt.sectionLength = flags & 0x3ff

	pmt.programNumber = reader.Read(2)

	pmt.versionNumber = reader.Read(1)

	pmt.sectionNumber = reader.Read(1)

	pmt.lastSectionNumber = reader.Read(1)

	pmt.pcrPid = reader.Read(2) & 0x1fff

	pmt.programInfoLength = reader.Read(2) & 0x3ff

	pmt.descriptor = reader.ReadBytes(uint64(pmt.programInfoLength))

	pmt.count = pmt.sectionLength - SKIP_BYTES - pmt.programInfoLength

	pmt.Print()

	for pmt.count > CRC_SIZE {

		pmtEntry := PmtEntry{}

		pmtEntry.Read(reader)

		pmt.entries = append(pmt.entries, pmtEntry)
		types[pmtEntry.pid] = pmtEntry.streamType
		elementaryConstructors[pmtEntry.pid] = ElementaryStreamPacket{}

		pmt.count -= (5 + pmtEntry.infoLength)

	}

}

//Adaptation Read
//size - Number of bytes in the adaptation field immediately following this byte (8 bits)
//
//discontinuity - Set to 1 if current TS packet is in a discontinuity state with respect to either the continuity counter or the program clock reference (1 bit)
//
//random - Set to 1 if the PES packet in this TS packet starts a video/audio sequence (1 bit)
//
//priority - 1 = higher priority
//
//hasPCR - 1 means adaptation field does contain a PCR field
//
//hasOPCR - 1 means adaptation field does contain an OPCR field
//
//hasSplice - 1 means presence of splice countdown field in adaptation field
//
//hasPrivate - 1 means presence of private data bytes in adaptation field
//
//hasExtension - 1 means presence of adaptation field extension
func (adaptation *Adaptation) Read() {

	var flags uint = 0
	var spliceFlag int = 0
	var pcrFlag int = 0
	var opcrFlag int = 0

	reader := data.NewReader(adaptation.byteChunk)

	adaptation.size = reader.Read(1)

	flags = reader.Read(1)

	adaptation.discontinuity = flags&0x80 > 0
	adaptation.random = flags&0x40 > 0
	adaptation.priority = flags&0x20 > 0
	adaptation.hasPCR = flags&0x10 > 0

	adaptation.hasOPCR = flags&0x08 > 0

	adaptation.hasSplice = flags&0x04 > 0

	adaptation.hasPrivate = flags&0x02 > 0
	adaptation.hasExtension = flags&0x01 > 1

	if adaptation.hasPCR {

		pcrFlag = 6

		adaptation.pcr.byteChunk = reader.ReadBytes(6)

		adaptation.pcr.Read()
	}

	if adaptation.hasOPCR {

		opcrFlag = 6

		adaptation.pcr.byteChunk = reader.ReadBytes(6)
		adaptation.opcr.Read()

	}

	if adaptation.hasSplice {

		spliceFlag = 1
		adaptation.splice = reader.Read(1)

	}

	adaptation.stuffing = int(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)

	reader.Cursor += uint64(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)

	payload := reader.ReadBytes(reader.Size - reader.Cursor)

	adaptation.payload = payload

	adaptation.Print()

}

//Program Read
//number – Program_number is a 16-bit field. It specifies the program to which the program_map_PID is
//applicable. When set to 0x0000, then the following PID reference shall be the network PID. For all other cases the value
//of this field is user defined. This field shall not take any single value more than once within one version of the Program
//Association Table
//
//pid – The program_map_PID is a 13-bit field specifying the PID of the Transport Stream packets
//which shall contain the program_map_section applicable for the program as specified by the program_number. No
//program_number shall have more than one program_map_PID assignment. The value of the program_map_PID is
//defined by the user, but shall only take values
func (program *Program) Read(reader *data.Reader) {

	//reader := data.NewReader(program.byteChunk)

	program.number = reader.Read(2)

	program.pid = reader.Read(2) & 0x1fff

	program.Print()

}

//PCR Read
//PCR fields
//valid for the program specified by program_number
func (pcr *Pcr) Read() {

	reader := data.NewReader(pcr.byteChunk)

	pcr.a = reader.Read(1)

	pcr.b = reader.Read(1)

	pcr.c = reader.Read(1)

	pcr.d = reader.Read(1)

	pcr.ef = reader.Read(1)

	pcr.e = pcr.ef & 0x80
	if pcr.e > 0 {
		pcr.e = 1
	} else {
		pcr.e = 0
	}

	pcr.f = pcr.ef & 0x01

	pcr.g = reader.Read(1)

	pcr.pcr = (pcr.a << 25) | (pcr.b << 17) | (pcr.c << 9) | (pcr.d << 1) | (pcr.e&1 | 0)
	pcr.ext = (pcr.f << 8) | pcr.g

	pcr.Print()

}

//PmtEntry Read
//streamType – This is an 8-bit field specifying the type of program element carried within the packets with the PID
//whose value is specified by the elementary_PID.
//
//pid – This is a 13-bit field specifying the PID of the Transport Stream packets which carry the associated
//program element.
//
//infoLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the descriptors of the associated program element immediately following the ES_info_length field.
func (pmtEntry *PmtEntry) Read(reader *data.Reader) {

	pmtEntry.streamType = reader.Read(1)

	pmtEntry.pid = reader.Read(2) & 0x1fff

	pmtEntry.infoLength = reader.Read(2) & 0x3ff

	pmtEntry.descriptor = reader.ReadBytes(uint64(pmtEntry.infoLength))

	pmtEntry.Print()

}

//Pes Read
//prefix – The packet_start_code_prefix is a 24-bit code. Together with the stream_id that follows, it
//constitutes a packet start code that identifies the beginning of a packet. The packet_start_code_prefix is the bit string
//'0000 0000 0000 0000 0000 0001' (0x000001 in hexadecimal).
//
//streamId – This 8-bit field shall have a value '1111 1111' (0xFF).
//
//packetLength – The PES_packet_length is a 16-bit field indicating the total number of bytes in the program_stream_directory immediately following this field
func (pes *Pes) Read() {

	reader := data.NewReader(pes.byteChunk)

	var prefix uint

	var headerLength uint

	var headerData []byte

	var flags uint

	prefix = reader.Read(3)

	if prefix == uint(0x000001) {

		pes.streamId = reader.Read(1)

		pes.packetLength = reader.Read(2)

		flags = reader.Read(2)

		headerLength = reader.Read(1)

		headerData = reader.ReadBytes(uint64(headerLength))

		if (flags & 0x0080) == 1 {
			pes.pts = ReadHeaderData(headerData)
		}
		if (flags & 0x0040) == 1 {
			pes.dts = ReadHeaderData(headerData)
		}
		pes.payload = reader.ReadBytes(reader.Size - reader.Cursor)

	}

}

//ElementaryStreamPacket Dispatch
//if unitstart, dump current PES and construct a new one,
//else append the es payload
func (elementaryStreamPacket *ElementaryStreamPacket) Dispatch() {

	var pesData Pes

	pesData = pesCollector[elementaryStreamPacket.pid]

	if elementaryStreamPacket.unitStart {

		if pesData, ok := pesCollector[elementaryStreamPacket.pid]; ok {

			pesData.pid = elementaryStreamPacket.pid
			pesData.streamType = types[elementaryStreamPacket.pid]
			pesData.Read()
			pesData.Print()

		}
		pesData = Pes{}

	}

	pesData.byteChunk = append(pesData.byteChunk, elementaryStreamPacket.payload...)

	pesCollector[elementaryStreamPacket.pid] = pesData

}

func ReadHeaderData(bytes []byte) uint {

	reader := data.NewReader(bytes)

	var a uint = (reader.Read(1) >> 1) & 0x07

	var b uint = reader.Read(1)

	var c uint = (reader.Read(1) >> 1) & 0x7f

	var d uint = reader.Read(1)

	var e uint = (reader.Read(1) >> 1) & 0x7f

	var timestamp uint = (a << 30) | (b << 22) | (c << 15) | (d << 7) | e

	return timestamp

}

func (elementaryStreamPacket *ElementaryStreamPacket) Print() {
	fmt.Println("\n:::ES:::\n")
	fmt.Println("payload = ", elementaryStreamPacket.payload)
}

func (program *Program) Print() {

	fmt.Println("\n:::Program:::\n")
	fmt.Println("pid = ", program.pid)
	fmt.Println("number = ", program.number)

}

func (pmtEntry *PmtEntry) Print() {
	fmt.Println("\n:::PmtEntry:::\n")
	fmt.Println("pid = ", pmtEntry.pid)
	fmt.Println("streamType = ", pmtEntry.streamType)
	fmt.Println("infoLength = ", pmtEntry.infoLength)
	fmt.Println("descriptor = ", pmtEntry.descriptor)

}

func (pmt *Pmt) Print() {

	fmt.Println("\n:::Pmt65435:::\n")
	fmt.Println("tableId = ", pmt.tableId)
	fmt.Println("pointerField = ", pmt.pointerField)
	fmt.Println("sectionSyntaxIndicator = ", pmt.sectionSyntaxIndicator)
	fmt.Println("sectionLength = ", pmt.sectionLength)
	fmt.Println("programNumber = ", pmt.programNumber)
	fmt.Println("versionNumber = ", pmt.versionNumber)
	fmt.Println("sectionNumber = ", pmt.sectionNumber)
	fmt.Println("lastSectionNumber = ", pmt.lastSectionNumber)
	fmt.Println("pcrPid = ", pmt.pcrPid)
	fmt.Println("programInfoLength = ", pmt.programInfoLength)
	fmt.Println("count = ", pmt.count)

	fmt.Println("descriptor = ", pmt.descriptor)

}

func (pes *Pes) Print() {

	fmt.Println("\n:::PES:::\n")
	fmt.Println("////////////////////////////////")
	fmt.Println("//pid = ", pes.pid)
	fmt.Println("//streamType = ", pes.streamType)
	fmt.Println("//streamId = ", pes.streamId)
	fmt.Println("//packetLength = ", pes.packetLength)
	fmt.Println("//pts = ", pes.pts)
	fmt.Println("//dts = ", pes.dts)
	fmt.Println("//payload length= ", len(pes.payload))
	fmt.Println("//nal = ", pes.nal)
	fmt.Println("////////////////////////////////")

}

func (pcr *Pcr) Print() {

	fmt.Println("\n:::Pcr:::\n")
	fmt.Println("a = ", pcr.a)
	fmt.Println("b = ", pcr.b)
	fmt.Println("c = ", pcr.c)
	fmt.Println("d = ", pcr.d)
	fmt.Println("ef = ", pcr.ef)
	fmt.Println("e = ", pcr.e)
	fmt.Println("f = ", pcr.f)
	fmt.Println("g = ", pcr.g)
	fmt.Println("pcr = ", pcr.pcr)
	fmt.Println("ext = ", pcr.ext)
}

func (adaptation *Adaptation) Print() {

	fmt.Println("\n:::Adaptation:::\n")
	fmt.Println("size = ", adaptation.size)
	fmt.Println("discontinuity = ", adaptation.discontinuity)
	fmt.Println("random = ", adaptation.random)
	fmt.Println("priority = ", adaptation.priority)
	fmt.Println("hasPCR = ", adaptation.hasPCR)
	fmt.Println("hasOPCR = ", adaptation.hasOPCR)
	fmt.Println("hasSplice = ", adaptation.hasSplice)
	fmt.Println("hasPrivate = ", adaptation.hasPrivate)
	fmt.Println("hasExtension = ", adaptation.hasExtension)

	fmt.Println("stuffing = ", adaptation.stuffing)

	fmt.Println("payload = ", adaptation.payload)

}

func (tsPacket *TsPacket) Print() {

	fmt.Println("\n:::TsRead:::\n")
	fmt.Println("sync = ", tsPacket.sync)
	fmt.Println("transportError = ", tsPacket.transportError)
	fmt.Println("unitStart = ", tsPacket.unitStart)
	fmt.Println("priority = ", tsPacket.priority)
	fmt.Println("pid = ", tsPacket.pid)
	fmt.Println("scramble = ", tsPacket.scramble)
	fmt.Println("hasAdaptation = ", tsPacket.hasAdaptation)
	fmt.Println("hasPayload = ", tsPacket.hasPayload)
	fmt.Println("continuity = ", tsPacket.continuity)
}

func (pat *Pat) Print() {
	fmt.Println("\n:::Pat:::\n")
	fmt.Println("tableId = ", pat.tableId)
	fmt.Println("pointerField = ", pat.pointerField)
	fmt.Println("sectionSyntaxIndicator = ", pat.sectionSyntaxIndicator)
	fmt.Println("sectionLength = ", pat.sectionLength)
	fmt.Println("transportStreamId = ", pat.transportStreamId)
	fmt.Println("versionNumber = ", pat.versionNumber)
	fmt.Println("currentNext = ", pat.currentNext)
	fmt.Println("sectionNumber = ", pat.sectionNumber)
	fmt.Println("lastSectionNumber = ", pat.lastSectionNumber)
	fmt.Println("count = ", pat.count)

}
