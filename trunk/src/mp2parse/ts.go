package main

import "mp2parse/data"
import "os"
import "fmt"

type Packet interface {
	Read()
	Print()
}

type PesDispenser struct {
	pesCollectors map[int]Writer
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
	splice int

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

type Writer struct {
	chunks []byte
	size   int
}

type Program struct {
	byteChunk []byte

	pid    uint32
	number uint32
}

var pmtConstructors map[uint32]Pmt
var entryConstructors map[uint32]PmtEntry
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

	fmt.Println("Size: ", len(bytes))

	for curser < len(bytes) {
		byteChunk := data.ReadBytes(curser, size, bytes)
		curser = curser + size

		tsPacket := TsPacket{}

		tsPacket.byteChunk = byteChunk

		tsPacket.Read()
	}

}

func (tsPacket TsPacket) Read() {

	var curser int = 0
	var flags uint32 = 0
	byteChunk := tsPacket.byteChunk

	tsPacket.sync = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	if tsPacket.sync == 71 {
		fmt.Println("\nG Found, Packet Start/////////////////////////")

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

		if pmtEntry, ok := entryConstructors[tsPacket.pid]; ok {
			pmtEntry.byteChunk = data.TruncateBytes(curser, byteChunk)
			pmtEntry.Read()
		}

	}

}

func (adaptation *Adaptation) Read() {

	var flags uint32 = 0
	var curser int = 0
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
		//TODO
	}

	if adaptation.hasOPCR {
		//TODO
	}

	if adaptation.hasSplice {
		//TODO
	}

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

	pat.Print()

}

func (pmt *Pmt) Read() {

	var CRC_SIZE uint32 = 4
	var SKIP_BYTES uint32 = 5
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

	pmt.count = pmt.sectionLength - SKIP_BYTES - pmt.programInfoLength

	for pmt.count > CRC_SIZE {

		pmtEntry := PmtEntry{}

		pmtEntry.byteChunk = data.TruncateBytes(curser, byteChunk)

		pmtEntry.Read()

		pmt.entries = append(pmt.entries, pmtEntry)
		types[pmtEntry.pid] = pmtEntry.streamType
		entryConstructors[pmtEntry.pid] = pmtEntry
		pmt.count -= SKIP_BYTES + pmtEntry.infoLength

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

	pmtEntry.Print()

}

func (pcr Pcr) Read() {

}

func (pcr Pcr) Print() {
}

func (pat Pat) Print() {
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

	fmt.Println("\nPacket End////////////////////////////")

}

func (pmt Pmt) Print() {

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

	fmt.Println("\nPacket End////////////////////////////")
}

func (pmtEntry PmtEntry) Print() {
	fmt.Println("\n:::PmtEntry:::\n")
	fmt.Println("pid = ", pmtEntry.pid)
	fmt.Println("streamType = ", pmtEntry.streamType)
	fmt.Println("infoLength = ", pmtEntry.infoLength)
	fmt.Println("descriptor = ", pmtEntry.descriptor)

	fmt.Println("\nPacket End////////////////////////////")

}

func (program Program) Print() {

	fmt.Println("\n:::Program:::\n")
	fmt.Println("pid = ", program.pid)
	fmt.Println("number = ", program.number)

}

func (adaptation Adaptation) Print() {
}

func (tsPacket TsPacket) Print() {

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
