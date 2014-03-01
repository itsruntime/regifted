package main

import (
	"fmt"
	"os"
	"regifted/data"
)

type Packet interface {
	Read()
	Print()
	Dispatch()
}

type Pes struct {
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
	payload   []byte
	byteChunk []byte

	size     uint32
	pcr      Pcr
	opcr     Pcr
	splice   uint32
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

	fmt.Printf("Attempting to read file, Run 6" + fileName + "\n")

	bytes := data.Read(fileName, 0)

	pat = Pat{}
	pat.tableId = 0

	Init()

	fmt.Println("Size: ", len(bytes))

	for curser < len(bytes) {
		byteChunk := data.ReadBytes(curser, size, bytes)
		curser = curser + size

		tsPacket := TsPacket{}

		tsPacket.byteChunk = byteChunk

    fmt.Println( "tsPacket.Read()" )
		tsPacket.Read()
		fmt.Println( "/tsPacket.Read()" )

	}

	for key := range pesCollector {

		CreateAndDispensePes(key, types[key])

	}

}

func Init() {
	pmtConstructors = make(map[uint32]Pmt)
	entryConstructors = make(map[uint32]PmtEntry)
	types = make(map[uint32]uint32)
	pesCollector = make(map[uint32]Pes)
	elementaryConstructors = make(map[uint32]ElementaryStreamPacket)
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
	//assert 
	if tsPacket.sync == 71 {

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

		tsPacket.Print()

		if tsPacket.hasAdaptation {
			tsPacket.adaptation.byteChunk = data.TruncateBytes(curser, byteChunk)
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
			elementaryStreamPacket.hasAdaptation = tsPacket.hasAdaptation

			if tsPacket.hasAdaptation {
				elementaryStreamPacket.payload = tsPacket.adaptation.payload
			} else {
				elementaryStreamPacket.payload = tsPacket.payload
			}

			elementaryStreamPacket.Read()
			elementaryStreamPacket.Dispatch()
			elementaryStreamPacket.Print()
		}

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

	adaptation.hasOPCR = flags&0x08 > 0

	adaptation.hasSplice = flags&0x04 > 0

	adaptation.hasPrivate = flags&0x02 > 0
	adaptation.hasExtension = flags&0x01 > 1

	if adaptation.hasPCR {

		pcrFlag = 6
		adaptation.pcr.byteChunk = data.TruncateBytes(curser, byteChunk)
		adaptation.pcr.Read()
		curser += 6
	}

	if adaptation.hasOPCR {

		opcrFlag = 6
		adaptation.pcr.byteChunk = data.TruncateBytes(curser, byteChunk)
		adaptation.opcr.Read()
		curser += 6

	}

	if adaptation.hasSplice {

		spliceFlag = 1
		adaptation.splice = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
		curser++
	}

	adaptation.stuffing = int(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)

	curser += int(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)

	payload := data.TruncateBytes(curser, byteChunk)

	adaptation.payload = payload

	adaptation.Print()

}

func (program *Program) Read() {

	var curser int = 0
	byteChunk := program.byteChunk

	program.number = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
	curser += 2

	program.pid = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk)) & 0x1fff
	curser += 2

	program.Print()

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

	pat.Print()

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

	pmt.Print()

	for pmt.count > CRC_SIZE {

		pmtEntry := PmtEntry{}

		pmtEntry.byteChunk = data.TruncateBytes(curser, byteChunk)

		pmtEntry.Read()
		curser += (int(pmtEntry.infoLength) + 5)

		pmt.entries = append(pmt.entries, pmtEntry)
		types[pmtEntry.pid] = pmtEntry.streamType
		elementaryConstructors[pmtEntry.pid] = ElementaryStreamPacket{}

		pmt.count -= (5 + pmtEntry.infoLength)

	}

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

	pmtEntry.Print()

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

	pcr.Print()

}

func (elementaryStreamPacket *ElementaryStreamPacket) Read() {

	if !elementaryStreamPacket.hasAdaptation {

		elementaryStreamPacket.payload = elementaryStreamPacket.byteChunk

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

	}

}

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

func (pmtEntry *PmtEntry) Print() {
	fmt.Println("\n:::PmtEntry:::\n")
	fmt.Println("pid = ", pmtEntry.pid)
	fmt.Println("streamType = ", pmtEntry.streamType)
	fmt.Println("infoLength = ", pmtEntry.infoLength)
	fmt.Println("descriptor = ", pmtEntry.descriptor)

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
	fmt.Println("\n:::ES:::\n")
	fmt.Println("payload = ", elementaryStreamPacket.payload)
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
