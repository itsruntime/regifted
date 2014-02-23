package main

import "mp2parse/data"
import "os"
import "fmt"

type Packet interface {
	Read()
}

type Reader struct {
	curser int
	size   int
	bytes  []byte
}

type PesDispenser struct {
	pesCollectors map[int]Writer
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

	byteChunk []byte

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
	pat *Pat

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

		tsPacket := TsPacket{}
		tsPacket.pat = &transport.pat

		tsPacket.byteChunk = byteChunk

		Dispense(reader, tsPacket)
	}

}

func PacketRead(packet Packet) {

	packet.Read()

}

func Dispense(reader Reader, packet Packet) {

	packet.Read()

}

func (tsPacket TsPacket) Read() {

	var curser int = 0
	var flags uint32 = 0
	byteChunk := tsPacket.byteChunk
	//var adaptation Adaptation

	tsPacket.sync = data.ReadSegemnt(data.ReadBytes(curser, 1, byteChunk))
	curser++

	if tsPacket.sync == 71 {
		fmt.Println("\nG found, packet contents: \n", byteChunk)

		flags = data.ReadSegemnt(data.ReadBytes(curser, 2, byteChunk))
		curser++

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

		if tsPacket.hasAdaptation {
			tsPacket.adaptation.byteChunk = data.ReadBytes(curser, 188-curser, byteChunk)
			tsPacket.adaptation.payload = &tsPacket.payload
			tsPacket.adaptation.Read()
		}

		if tsPacket.pid == 0 {
			tsPacket.pat.byteChunk = data.ReadBytes(curser, 188-curser, byteChunk)
			PacketRead(tsPacket.pat)
		}

		//fmt.Println("\n 67898767 tsPacket \n", tsPacket)

	}

}

func (adaptation Adaptation) Read() {

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

	payload := data.ReadBytes(curser, 188-curser, byteChunk)

	adaptation.payload = &payload

}

func (pat Pat) Read() {

	//var  SKIP_BYTES int = 5
	// var CRC_SIZE int = 4

	//TODO

}

func (pcr Pcr) Read() {

}
