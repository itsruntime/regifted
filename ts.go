package main

import "data"
import "os"
import "fmt"

type Packet interface {
	Read()
	Print()
	Dispatch()
}

type Pes struct {
	//pesCollectors map[int]Writer

	byteChunk []byte

	pid          uint32
	streamType   uint32
	streamId     uint32
	packetLength uint32
	flags        uint32
	pts          uint32
	dts          uint32
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
	tableId                uint32
	flags                  uint32
	sectionSyntaxIndicator bool
	sectionLength          uint32
	transportStreamId      uint32
	versionNumber          uint32
	currentNext            uint32
	sectionNumber          uint32
	lastSectionNumber      uint32
	count                  uint32

	programs []Program
}

type TsPacket struct {
	byteChunk []byte

	sync  uint32
	flags uint32

	payload []byte

	transportError bool
	unitStart      bool
	priority       bool
	pid            uint32
	scramble       uint32
	hasAdaptation  bool
	hasPayload     bool
	continuity     uint32
	adaptation     Adaptation
}

type Adaptation struct {
	payload   *[]byte
	byteChunk []byte

	size   uint32
	pcr    Pcr
	opcr   Pcr
	splice uint32

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
	pcr       uint32
	ext       uint32
	byteChunk []byte

	a  uint32
	b  uint32
	c  uint32
	d  uint32
	ef uint32
	e  uint32
	f  uint32
	g  uint32
}

type Pmt struct {
	byteChunk              []byte
	pointerField           bool
	unitStart              bool
	tableId                uint32
	sectionSyntaxIndicator bool
	sectionLength          uint32
	programNumber          uint32
	versionNumber          uint32
	sectionNumber          uint32
	lastSectionNumber      uint32

	pcrPid            uint32
	programInfoLength uint32
	descriptor        []byte
	count             uint32
	entries           []PmtEntry

	types map[uint32]uint32
}

type PmtEntry struct {
	byteChunk []byte

	streamType uint32
	pid        uint32
	infoLength uint32
	descriptor []byte
}

type ElementaryStreamPacket struct {
	byteChunk []byte
	payload   []byte

	unitStart bool

	pid           uint32
	hasAdaptation bool
}

type Writer struct {
	chunks []byte
	size   int
}

type Program struct {
	byteChunk []byte

	pid    uint32
	number uint32
}

var pesCollector map[uint32]Pes
var pmtConstructors map[uint32]Pmt
var entryConstructors map[uint32]PmtEntry
var elementaryConstructors map[uint32]ElementaryStreamPacket
var types map[uint32]uint32
var pat Pat

func main() {

	fileName := os.Args[1]

	var curser int = 0
	var size int = 188

	fmt.Printf("Attempting to read file:" + fileName + "\n")

	bytes := data.Read(fileName, 0)

	pat = Pat{}
	pat.tableId = 0

	pmtConstructors = make(map[uint32]Pmt)
	entryConstructors = make(map[uint32]PmtEntry)
	types = make(map[uint32]uint32)
	pesCollector = make(map[uint32]Pes)
	elementaryConstructors = make(map[uint32]ElementaryStreamPacket)

	fmt.Println("Size: ", len(bytes))

	for curser < len(bytes) {
		byteChunk := data.ReadBytes(curser, size, bytes)
		curser = curser + size

		tsPacket := TsPacket{}

		tsPacket.byteChunk = byteChunk

		tsPacket.Read()

	}

	for key := range pesCollector {

		CreateAndDispensePes(key, types[key])

	}

}

func CreateAndDispensePes(pid uint32, streamType uint32) {

	pes := pesCollector[pid]

	pes.pid = pid

	pes.streamType = streamType

	pes.Read()

	pes.Print()

}

func (tsPacket *TsPacket) Read() {

	var curser int = 0
	var flags uint32 = 0
	byteChunk := tsPacket.byteChunk

	tsPacket.sync = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	if tsPacket.sync == 71 {
		//fmt.Println("\nG Found, Packet Start/////////////////////////")

		flags = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
		curser += 2

		tsPacket.transportError = flags&0x8000 > 0
		tsPacket.unitStart = flags&0x4000 > 0
		tsPacket.priority = flags&0x2000 > 0
		tsPacket.pid = flags & 0x1fff

		flags = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
		curser++

		tsPacket.scramble = flags & 0xc0 >> 6
		tsPacket.hasAdaptation = flags&0x20 > 0
		tsPacket.hasPayload = flags&0x10 > 0
		tsPacket.continuity = flags & 0x0f

		//tsPacket.Print()

		if tsPacket.hasAdaptation {
			tsPacket.adaptation.byteChunk = data.TruncateBytes(curser, byteChunk)
			tsPacket.adaptation.payload = &tsPacket.payload
			tsPacket.adaptation.Read()
		}

		if tsPacket.pid == 0 {
			pat.byteChunk = data.TruncateBytes(curser, byteChunk)

			pat.unitStart = tsPacket.unitStart
			pat.Read()
		}

		if pmt, ok := pmtConstructors[tsPacket.pid]; ok {
			pmt.unitStart = tsPacket.unitStart
			pmt.byteChunk = data.TruncateBytes(curser, byteChunk)
			pmt.Read()
		}

		if elementaryStreamPacket, ok := elementaryConstructors[tsPacket.pid]; ok {
			elementaryStreamPacket.pid = tsPacket.pid
			elementaryStreamPacket.unitStart = tsPacket.unitStart
			elementaryStreamPacket.byteChunk = data.TruncateBytes(curser, byteChunk)
			elementaryStreamPacket.payload = tsPacket.payload
			elementaryStreamPacket.hasAdaptation = tsPacket.hasAdaptation
			elementaryStreamPacket.Read()
			elementaryStreamPacket.Dispatch()

			//fmt.Println("PID = ", tsPacket.pid)
			//fmt.Println("UNITSTART = ", elementaryStreamPacket.unitStart)
			//fmt.Println("BYTECHUCNK1278909878345 = ", len(pesCollector[elementaryStreamPacket.pid].byteChunk))

			//elementaryStreamPacket.Print()
		}

		//if pmtEntry, ok := entryConstructors[tsPacket.pid]; ok {
		//	pmtEntry.byteChunk = data.TruncateBytes(curser, byteChunk)
		//	pmtEntry.Read()
		//}

	}

}

func (adaptation *Adaptation) Read() {

	var flags uint32 = 0
	var curser int = 0
	var spliceFlag int = 0
	var pcrFlag int = 0
	var opcrFlag int = 0

	byteChunk := adaptation.byteChunk

	adaptation.size = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	flags = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	adaptation.discontinuity = flags&0x80 > 0
	adaptation.random = flags&0x40 > 0
	adaptation.priority = flags&0x20 > 0
	adaptation.hasPCR = flags&0x10 > 0
	//pcrFlag = flags & 0x10
	adaptation.hasOPCR = flags&0x08 > 0
	//opcrFlag = flags & 0x08
	adaptation.hasSplice = flags&0x04 > 0
	//spliceFlag = flags & 0x04
	adaptation.hasPrivate = flags&0x02 > 0
	adaptation.hasExtension = flags&0x01 > 1

	if adaptation.hasPCR {
		fmt.Println("BIG HERE890769 = ", pcrFlag)
		pcrFlag = 6
		adaptation.pcr.byteChunk = data.TruncateBytes(curser, byteChunk)
		adaptation.pcr.Read()
		curser += 6
	}

	if adaptation.hasOPCR {
		fmt.Println("BIG HERE789087 = ", pcrFlag)
		opcrFlag = 6
		adaptation.pcr.byteChunk = data.TruncateBytes(curser, byteChunk)
		adaptation.opcr.Read()
		curser += 6

	}

	if adaptation.hasSplice {
		fmt.Println("BIG HERE7897890980 = ", pcrFlag)
		spliceFlag = 1
		adaptation.splice = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
		curser++
	}

	fmt.Println("BIG pcrFlag123432 = ", pcrFlag)
	fmt.Println("BIG opcrFlag = ", opcrFlag)
	fmt.Println("BIG spliceFlag = ", spliceFlag)

	fmt.Println("BIG pcrFlagAND = ", (pcrFlag & 6))
	fmt.Println("BIG opcrFlagAND = ", (opcrFlag & 6))
	fmt.Println("BIG spliceFlagAND = ", (spliceFlag & 1))

	//stuffing := data.ReadBytes(curser, (int(adaptation.size) - 1 - (pcrFlag& 6) - (opcrFlag & 6) - (spliceFlag & 1)),  byteChunk)
	//fmt.Println("BIG NUMNBER678976789876878909898788789 = ", (adaptation.size - 1 - (pcrFlag) - (opcrFlag) - (spliceFlag & 1)))



	curser += int(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)
	fmt.Println("SIZE 67898767890987898709 = ", adaptation.size)
	fmt.Println("BIG NUMNBER67897678987687890989878909878988789 = ", curser)

	payload := data.TruncateBytes(curser, byteChunk)

	adaptation.payload = &payload

}

func (program *Program) Read() {

	var curser int = 0
	byteChunk := program.byteChunk

	program.number = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	program.pid = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x1fff
	curser += 2

}

func (pat *Pat) Read() {

	var SKIP_BYTES uint32 = 5
	var CRC_SIZE uint32 = 4
	var PROGRAM_SIZE uint32 = 4
	var flags uint32 = 0

	var READ_SIZE int = 4
	var curser int = 0

	var flag bool = false

	byteChunk := pat.byteChunk

	if data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk)) == 1 {
		flag = true
	}
	curser++

	pat.pointerField = (pat.unitStart && flag) || false

	pat.tableId = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	flags = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	pat.sectionSyntaxIndicator = flags&0x8000 > 0

	pat.sectionLength = flags & 0x3ff

	pat.transportStreamId = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	flags = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pat.versionNumber = flags & 0x3ffe
	pat.currentNext = flags & 0x0001

	pat.sectionNumber = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pat.lastSectionNumber = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pat.count = pat.sectionLength - SKIP_BYTES

	for pat.count > CRC_SIZE {
		program := Program{}
		pmt := Pmt{}

		program.byteChunk = data.TruncateBytes(curser, byteChunk)
		curser = curser + READ_SIZE

		program.Read()

		pat.programs = append(pat.programs, program)
		pmtConstructors[program.pid] = pmt

		pat.count = pat.count - PROGRAM_SIZE
	}

	//pat.Print()

}

func (pmt *Pmt) Read() {

	var CRC_SIZE uint32 = 4
	var SKIP_BYTES uint32 = 9
	var flags uint32 = 0

	var curser int = 0

	var flag bool = false

	byteChunk := pmt.byteChunk

	if data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk)) == 1 {
		flag = true
	}
	curser++

	pmt.pointerField = (pmt.unitStart && flag) || false

	pmt.tableId = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	flags = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	pmt.sectionSyntaxIndicator = flags&0x8000 > 0
	pmt.sectionLength = flags & 0x3ff

	pmt.programNumber = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	pmt.versionNumber = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pmt.sectionNumber = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pmt.lastSectionNumber = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pmt.pcrPid = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x1fff
	curser += 2

	pmt.programInfoLength = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x3ff
	curser += 2

	pmt.descriptor = data.ReadBytes(curser, int(pmt.programInfoLength), byteChunk)
	curser += int(pmt.programInfoLength)

	pmt.count = pmt.sectionLength - SKIP_BYTES - pmt.programInfoLength

	//count = self.sectionLength - self.SKIP_BYTES - self.programInfoLength

	for pmt.count > CRC_SIZE {

		pmtEntry := PmtEntry{}

		pmtEntry.byteChunk = data.TruncateBytes(curser, byteChunk)

		pmtEntry.Read()
		curser += (int(pmtEntry.infoLength) + 5)

		pmt.entries = append(pmt.entries, pmtEntry)
		types[pmtEntry.pid] = pmtEntry.streamType
		elementaryConstructors[pmtEntry.pid] = ElementaryStreamPacket{}

		//entryConstructors[pmtEntry.pid] = pmtEntry
		pmt.count -= (5 + pmtEntry.infoLength)

	}

	pmt.Print()
}

func (pmtEntry *PmtEntry) Read() {

	var curser int = 0
	byteChunk := pmtEntry.byteChunk

	pmtEntry.streamType = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pmtEntry.pid = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x1fff
	curser += 2

	pmtEntry.infoLength = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x3ff
	curser += 2

	pmtEntry.descriptor = data.ReadBytes(curser, int(pmtEntry.infoLength), byteChunk)
	curser += int(pmtEntry.infoLength)

	//pmtEntry.Print()

}

func (pcr *Pcr) Read() {

	var curser int = 0
	byteChunk := pcr.byteChunk

	pcr.a = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++
	pcr.b = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++
	pcr.c = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++
	pcr.d = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++
	pcr.ef = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	pcr.e = pcr.ef & 0x80
	if pcr.e > 0 {
		pcr.e = 1
	} else {
		pcr.e = 0
	}

	pcr.f = pcr.ef & 0x01

	pcr.g = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++
	pcr.pcr = (pcr.a << 25) | (pcr.b << 17) | (pcr.c << 9) | (pcr.d << 1) | (pcr.e&1 | 0)
	pcr.ext = (pcr.f << 8) | pcr.g

}

func (elementaryStreamPacket *ElementaryStreamPacket) Read() {
	fmt.Println("ADAPTATION =  ", elementaryStreamPacket.hasAdaptation)
	if !elementaryStreamPacket.hasAdaptation {
		fmt.Println("PAYLOADSET TO BYTE CHUNK = ")
		elementaryStreamPacket.payload = elementaryStreamPacket.byteChunk

	} else {
		fmt.Println("PAYLOADSET = ", elementaryStreamPacket.payload)
	}

}

func (pes *Pes) Read() {

	var curser int = 0

	var prefix uint32

	var headerLength uint32

	var headerData []byte

	var flags uint32

	byteChunk := pes.byteChunk

	prefix = data.ReadSegemnt(data.ReadBytes(curser, 3, byteChunk))
	curser += 3

	//fmt.Println("prefix12343232 = ", byteChunk)

	if prefix == uint32(0x000001) {

		pes.streamId = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
		curser++

		pes.packetLength = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
		curser += 2

		flags = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
		curser += 2

		headerLength = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
		curser++

		headerData = data.ReadBytes(curser, int(headerLength), byteChunk)
		curser += int(headerLength)

		if (flags & 0x0080) == 1 {
			pes.pts = data.ReadHeaderData(headerData)
		}
		if (flags & 0x0040) == 1 {
			pes.dts = data.ReadHeaderData(headerData)
		}
		pes.payload = data.TruncateBytes(curser, byteChunk)

		//pes.nal =
	}

}

func (elementaryStreamPacket *ElementaryStreamPacket) Dispatch() {

	var pesData Pes

	pesData = pesCollector[elementaryStreamPacket.pid]

	if elementaryStreamPacket.unitStart {

		if pesData, ok := pesCollector[elementaryStreamPacket.pid]; ok {
			//fmt.Println("BYTTTEEESSS = ", pesData.byteChunk)
			if len(pesData.byteChunk) != 0 {
				//fmt.Println("HERE = ", pesData.byteChunk[0])
				pesData.pid = elementaryStreamPacket.pid
				pesData.streamType = types[elementaryStreamPacket.pid]
				pesData.Read()
				pesData.Print()

			}
			pesData = Pes{}

		}

	}

	//fmt.Println("6797678909879 = ", elementaryStreamPacket.payload)

	pesData.byteChunk = append(pesData.byteChunk, elementaryStreamPacket.payload...)

	pesCollector[elementaryStreamPacket.pid] = pesData

}

func (pcr *Pcr) Print() {
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

	for i := 0; i < len(pat.programs); i++ {

		pat.programs[i].Print()
	}

	//fmt.Println("\nPacket End////////////////////////////")

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

	for i := 0; i < len(pmt.entries); i++ {

		pmt.entries[i].Print()
	}

	//fmt.Println("\nPacket End////////////////////////////")
}

func (pmtEntry *PmtEntry) Print() {
	fmt.Println("\n:::PmtEntry:::\n")
	fmt.Println("pid = ", pmtEntry.pid)
	fmt.Println("streamType = ", pmtEntry.streamType)
	fmt.Println("infoLength = ", pmtEntry.infoLength)
	fmt.Println("descriptor = ", pmtEntry.descriptor)

	//fmt.Println("\nPacket End////////////////////////////")

}

func (program *Program) Print() {

	fmt.Println("\n:::Program:::\n")
	fmt.Println("pid = ", program.pid)
	fmt.Println("number = ", program.number)

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

func (elementaryStreamPacket *ElementaryStreamPacket) Print() {
	//fmt.Println("\n:::ES:::\n")
	//fmt.Println("//payload = ", elementaryStreamPacket.payload)
}

func (adaptation *Adaptation) Print() {
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
