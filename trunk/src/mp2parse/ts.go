package main

import "mp2parse/data"
import "os"
import "fmt"

type Reader struct {
	curser int
	size   int
	data   []byte
}

type PesDispenser struct {
	pesCollectors map[int]*Writer
}

type TransportStream struct {
	pat          Pat
	constructors map[int]string
	types        map[int]Type
}

type Type struct {
}

type Constructor struct {
}

type Pat struct {
	SKIP_BYTES int
	CRC_SIZE   int

	//self.packet = packet
	pointerField           int
	tableId                int
	flags                  int
	sectionSyntaxIndicator int
	sectionLength          int
	transportStreamId      int
	versionNumber          int
	currentNext            int
	sectionNumber          int
	lastSectionNumber      int
	programs               []Program
	count                  int
}

type TsPacket struct {
	sync  int
	flags int

	transportError int
	unitStart      int
	priority       int
	pid            int
	scramble       int
	hasAdaptation  int
	hasPayload     int
	continuity     int
}

type Writer struct {
	chunks []byte
	size   int
}

type Program struct {
}

func main() {

	fileName := os.Args[1]

	fmt.Printf("Attempting to read file:" + fileName + "\n")

	bytes := data.Read(fileName, 0)

	reader := Reader{0, 188, bytes}

	transport := TransportStream{Pat{}, map[int]string{}, map[int]Type{}}

	transport.pat.tableId = 0

	fmt.Println("Size: ", len(bytes))

	for reader.curser < len(bytes) {
		byteChunk := data.ReadBytes(reader.curser, reader.size, bytes)
		reader.curser = reader.curser + reader.size
		dispense(byteChunk, reader)
	}

}

func dispense(byteChunk []byte, reader Reader) {

	tsPacket := TsPacket{}

	TsRead(tsPacket, byteChunk)

}

func patRead(pat Pat, byteChunk []byte) {

}

func TsRead(tsPacket TsPacket, byteChunk []byte) {

	if byteChunk[0] == 71 {
		fmt.Println("\nG found, packet contents: \n", byteChunk)

	}

}
