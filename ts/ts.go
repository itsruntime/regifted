package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regifted/data"
)

type Writer struct {
	chunks []byte
	size   int
}

var globals_initialized bool
var pesCollector map[uint]Pes
var pmtConstructors map[uint]Pmt
var entryConstructors map[uint]PmtEntry
var elementaryConstructors map[uint]ElementaryStreamPacket
var types map[uint]uint
var pat Pat

func main() {
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

	rc := Init()
	if rc != true {
		log.Printf("could not initialize global state\n")
		os.Exit(71)
	}

	fmt.Println("Size: ", len(bytes))

	s := uint64(len(bytes))
	for reader.Cursor < s {

		byteChunk := reader.ReadBytes(188)

		tsPacket := TsPacket{}

		tsPacket.byteChunk = byteChunk

		tsPacket.Read()

	}

	for key := range pesCollector {

		CreateAndDispensePes(key, types[key])

	}

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
func Init() bool {
	if globals_initialized == true {
		log.Printf("EE attempted to initialize globals twice\n")
		return false
	}
	pmtConstructors = make(map[uint]Pmt)
	entryConstructors = make(map[uint]PmtEntry)
	types = make(map[uint]uint)
	pesCollector = make(map[uint]Pes)
	elementaryConstructors = make(map[uint]ElementaryStreamPacket)
	pat = Pat{}
	pat.tableId = 0
	globals_initialized = true
	return true
}

func DeleteState() {
	if globals_initialized == false {
		return
	}
	globals_initialized = false
	Init()
	globals_initialized = false
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
