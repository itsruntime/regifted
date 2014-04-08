package ts

import (
	"flag"
	"fmt"
	// "io/ioutil"
	"log"
	"os"
	"regifted/data"
)

const TS_PACKET_SIZE = 188
const PACKET_TYPE_ERROR = 0
const PACKET_TYPE_PAT = 2
const PACKET_TYPE_PCR = 3
const PACKET_TYPE_PES = 4
const PACKET_TYPE_ES = 1
const PACKET_TYPE_PMT = 5
const PACKET_TYPE_PROGRAM = 6
const PACKET_TYPE_TS = 7

type TSState struct {
	globals_initialized bool
	// the keys in these maps are pids
	pesCollector           map[uint]Pes
	pmtConstructors        map[uint]Pmt
	entryConstructors      map[uint]PmtEntry
	elementaryConstructors map[uint]ElementaryStreamPacket
	types                  map[uint]uint
	pat                    Pat

	bytes  []byte
	reader *data.Reader
	pcr    uint

	// pes.streamtype -> pes[]
	pesMap map[uint][]Pes
}

// this is still global state - it's a temporary step in-between
var state TSState

func Load(bytes []byte) *TSState {
	fmt.Println("load()")

	state = TSState{}
	var state2 *TSState
	state2 = &state
	state2.bytes = bytes
	state2.reader = data.NewReader(bytes)
	state2.main()
	return state2
}

func (state *TSState) main() {
	reader := state.reader
	bytes := state.bytes

	rc := state.Init()
	if rc != true {
		log.Printf("could not initialize state\n")
		os.Exit(71)
	}
	fmt.Println("Size: ", len(bytes))
	s := uint64(len(bytes))

	for reader.Cursor < s {
		var pesData *Pes
		byteChunk := reader.ReadBytes(TS_PACKET_SIZE)
		tsPacket := TsPacket{}
		tsPacket.byteChunk = byteChunk
		packetType, packetReader := tsPacket.Read()

		switch {
		case packetType == PACKET_TYPE_PAT:
			readPat(&tsPacket, packetReader)

		case packetType == PACKET_TYPE_PMT:
			readPMT(&tsPacket, packetReader)

		case packetType == PACKET_TYPE_ES:
			pesData = readES(&tsPacket, packetReader)

			if pesData != nil {
				if state.pesMap[pesData.streamType] != nil {
					state.pesMap[pesData.streamType] = make([]Pes, 1, 1)

				}

				state.pesMap[pesData.streamType] = append(state.pesMap[pesData.streamType], *pesData)

			}
		}

		if tsPacket.hasAdaptation && tsPacket.adaptation.hasPCR {
			state.pcr = tsPacket.adaptation.pcr.pcr
		}
	}

	for key := range state.pesCollector {
		state.CreateAndDispensePes(key, state.types[key])
	}

}

//CreateAndDispensePes
//Dump the remaining PES
func (state *TSState) CreateAndDispensePes(pid uint, streamType uint) {
	pes := state.pesCollector[pid]
	pes.pid = pid
	pes.streamType = streamType
	pes.Read()
	pes.Print()
}

func readPat(tsPacket *TsPacket, reader *data.Reader) {
	state.pat.byteChunk = reader.ReadBytes(reader.Size - reader.Cursor)
	state.pat.unitStart = tsPacket.unitStart
	state.pat.Read()

	state.loadPAT(&state.pat)
}

func readPMT(tsPacket *TsPacket, reader *data.Reader) {
	pmt, _ := state.pmtConstructors[tsPacket.pid]
	pmt.unitStart = tsPacket.unitStart
	pmt.byteChunk = reader.ReadBytes(reader.Size - reader.Cursor)
	pmt.Read()

	state.loadPMT(&pmt)
}

func readES(tsPacket *TsPacket, reader *data.Reader) *Pes {
	var pesData *Pes
	elementaryStreamPacket, _ := state.elementaryConstructors[tsPacket.pid]
	elementaryStreamPacket.pid = tsPacket.pid
	elementaryStreamPacket.unitStart = tsPacket.unitStart

	if tsPacket.hasAdaptation {
		elementaryStreamPacket.payload = tsPacket.adaptation.payload
	} else {
		elementaryStreamPacket.payload = reader.ReadBytes(reader.Size - reader.Cursor)
	}

	pesData = state.dispatch(&elementaryStreamPacket)
	elementaryStreamPacket.Print()
	return pesData
}

// todo( mathew guest ) I think golang wants to use error as return codes but
// it's a little slow so I'm cheating
func getFilepath() (string, int) {
	flag.Parse()
	argc := flag.NArg()
	if argc < 1 {
		log.Printf("Usage: " + os.Args[0] + " [input ts file]\n")
		return "", 66
	}
	if argc > 1 {
		log.Printf("Ignoring all but first argument.\n")
		os.Exit(1)
	}
	fileName := os.Args[1]
	return fileName, 0
}

//Init
//Initialize the constructors
func (state *TSState) Init() bool {
	if state.globals_initialized == true {
		log.Printf("EE attempted to initialize globals twice\n")
		return false
	}
	state.pmtConstructors = make(map[uint]Pmt)
	state.entryConstructors = make(map[uint]PmtEntry)
	state.types = make(map[uint]uint)
	state.pesCollector = make(map[uint]Pes)
	state.elementaryConstructors = make(map[uint]ElementaryStreamPacket)
	state.pat = Pat{}
	state.pat.tableId = 0
	state.pesMap = make(map[uint][]Pes)
	state.globals_initialized = true
	return true
}

func (state *TSState) DeleteState() {
	if state.globals_initialized == false {
		return
	}
	state.globals_initialized = false
	state.Init()
	state.globals_initialized = false
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
