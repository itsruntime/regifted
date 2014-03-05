package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regifted/data"
)

const TS_PACKET_SIZE = 188
const PACKET_TYPE_ERROR = 0
const PACKET_TYPE_PAT = 2
const PACKET_TYPE_PCR = 3
const PACKET_TYPE_PES = 4
const PACKET_TYPE_PMT = 5
const PACKET_TYPE_PROGRAM = 6
const PACKET_TYPE_TS = 7

type Writer struct {
	chunks []byte
	size   int
}

type TSState struct {
	globals_initialized    bool
	pesCollector           map[uint]Pes
	pmtConstructors        map[uint]Pmt
	entryConstructors      map[uint]PmtEntry
	elementaryConstructors map[uint]ElementaryStreamPacket
	types                  map[uint]uint
	pat                    Pat
}

// this is still global state - it's a temporary step in-between
var state TSState

func main() {
	state = TSState{}
	state.main()
}

func (state *TSState) main() {
	fileName, rv := getFilepath()
	if rv != 0 {
		os.Exit(rv)
	}
	fmt.Printf("Attempting to read file, Run 7 " + fileName + "\n")

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("did not open file\n")
		// os.Exit(66)
		// seems like panic is better?
		panic(err)
	}
	reader := data.NewReader(bytes)
	_ = reader

	rc := Init()
	if rc != true {
		log.Printf("could not initialize global state\n")
		os.Exit(71)
	}
	fmt.Println("Size: ", len(bytes))

	s := uint64(len(bytes))
	for reader.Cursor < s {
		byteChunk := reader.ReadBytes(TS_PACKET_SIZE)
		tsPacket := TsPacket{}
		tsPacket.byteChunk = byteChunk
		packetType, packetReader := tsPacket.Read()

		// depending on what kind of packet it is, process it for that packet
		if packetType == PACKET_TYPE_PAT {
			state.pat.byteChunk = packetReader.ReadBytes(packetReader.Size - packetReader.Cursor)
			state.pat.unitStart = tsPacket.unitStart
			state.pat.Read()

		} else if packetType == PACKET_TYPE_PMT {
			if pmt, ok := state.pmtConstructors[tsPacket.pid]; ok {
				pmt.unitStart = tsPacket.unitStart
				pmt.byteChunk = packetReader.ReadBytes(packetReader.Size - packetReader.Cursor)
				pmt.Read()
			}

		} else if packetType == PACKET_TYPE_PES {
			if elementaryStreamPacket, ok := state.elementaryConstructors[tsPacket.pid]; ok {
				elementaryStreamPacket.pid = tsPacket.pid
				elementaryStreamPacket.unitStart = tsPacket.unitStart

				if tsPacket.hasAdaptation {
					elementaryStreamPacket.payload = tsPacket.adaptation.payload
				} else {
					elementaryStreamPacket.payload = packetReader.ReadBytes(packetReader.Size - packetReader.Cursor)
				}

				elementaryStreamPacket.Dispatch()
				elementaryStreamPacket.Print()
			}

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
func Init() bool {
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
	state.globals_initialized = true
	return true
}

func DeleteState() {
	if state.globals_initialized == false {
		return
	}
	state.globals_initialized = false
	Init()
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
