package main

import (
	// "fmt"
	"log"
	"testing"
)

func awfulStateSetup() {
	DeleteState()
	Init()
}

func TestTransportPacketRead(t *testing.T) {
  // ts header
  // n_bits value
	// 8			sync byte is always 'G' 0x47
	// 1			transport error flag
	// 1			payload unit start indicator
	// 1			transport priority indicator
	// 13			pid
	// 2			scrambling control
	// 1			adapaption field exists
	// 1			contains payload
	// 4			continuity counter
	//
	// bars are byte marks. header is 4 bytes total; i.e. 8 hex chars.
	//                     |                    |                    |
  // [sync]                [tei] [psi] [tp] [         pid        ]   [sc] [af] [pf] [cont]

  // bin  	hex
  // 0000 	0
  // 0001 	1
  // 0010 	2
  // 0011 	3
  // 0100 	4
  // 0101 	5
  // 0110 	6
  // 0111 	7
  // 1000 	8
  // 1001 	9
  // 1010 	A
  // 1011 	B
  // 1100 	C
  // 1101 	D
  // 1110 	E
  // 1111 	F

	// could endian-ness matter theoretically?
	var packetBytes []byte
	var packetString string
	var err error

	awfulStateSetup()

	packetString = `
		47 02 62 17 80 70 24 F1 31 D4 33 F3 88 8B 23 5A
		94 13 D4 1D DD CD 61 D8 73 08 82 88 B1 13 85 BB
		D0 93 C4 41 54 BE 68 70 70 EA 3D CF 70 1F 50 CD
		4E 26 A5 00 D2 1C C4 4C EE 71 C4 68 38 2B 91 45
		44 40 F7 50 85 5D 7F DD 62 58 28 72 D6 4F 39 4A
		A1 86 83 28 00 9E 2E 01 F0 AE 79 CC 16 F1 3C 19
		7F FE A3 25 F0 1A 5B 0E CC 45 B6 1C D8 EC 2C E3
		48 E7 78 9F F1 E1 62 30 78 1C 80 D5 D0 CF DA C0
		6A A0 01 F4 03 FA AF 80 B2 34 DE 01 DC 14 46 72
		0F 23 CB 70 21 DE AE 70 1D A0 F1 22 E1 10 79 1C
		58 EA D4 4F 50 07 D1 3E D8 77 E4 63 65 2C E6 D0
		9A 11 82 26 CC 62 D6 2E 00 1F DA C3`
	err = generateBytesFromString(&packetBytes, &packetString)
	if err != nil {
		log.Printf( "EE problem in test suite" )
	}
	packet := TsPacket{byteChunk: packetBytes}
	packet.Read()
	if packet.sync != 0X47 { // 71 = 'G' = 0x47
		t.Error("Transport Stream Packet read " +
			"sync byte incorrectly.")
	}
	if packet.transportError != false {
		t.Error("Transport Stream Packet read " +
			"transport error bit incorrectly.")
	}
	if packet.unitStart != false {
		t.Error("Transport Stream Packet read " +
			"unit start incorrectly.")
	}
	if packet.priority != false {
		t.Error("Transport Stream Packet read " +
			"priority incorrectly.")
	}
	if packet.pid != 610 {
		t.Error("Transport Stream Packet read " +
			"packet id incorrectly.")
	}
	if packet.scramble != 0 {
		t.Error("Transport Stream Packet read " +
			"scrambling control incorrectly.")
	}
	if packet.hasAdaptation != false {
		t.Error("Transport Stream Packet read " +
			"adaptation indiciator incorrectly.")
	}
	if packet.hasPayload != true {
		t.Error("Transport Stream Packet read " +
			"payload indicator incorrectly.")
	}
	if packet.continuity != 7 {
		t.Error("Transport Stream Packet read " +
			"continuity incorrectly.")
	}

	awfulStateSetup()
  packetString = `
   47 00 01 17 80 70 24 F1 31 D4 33 F3 88 8B 23 5A
   94 13 D4 1D DD CD 61 D8 73 08 82 88 B1 13 85 BB
   D0 93 C4 41 54 BE 68 70 70 EA 3D CF 70 1F 50 CD
   4E 26 A5 00 D2 1C C4 4C EE 71 C4 68 38 2B 91 45
   44 40 F7 50 85 5D 7F DD 62 58 28 72 D6 4F 39 4A
   A1 86 83 28 00 9E 2E 01 F0 AE 79 CC 16 F1 3C 19
   7F FE A3 25 F0 1A 5B 0E CC 45 B6 1C D8 EC 2C E3
   48 E7 78 9F F1 E1 62 30 78 1C 80 D5 D0 CF DA C0
   6A A0 01 F4 03 FA AF 80 B2 34 DE 01 DC 14 46 72
   0F 23 CB 70 21 DE AE 70 1D A0 F1 22 E1 10 79 1C
   58 EA D4 4F 50 07 D1 3E D8 77 E4 63 65 2C E6 D0
   9A 11 82 26 CC 62 D6 2E 00 1F DA C3`
  err = generateBytesFromString(&packetBytes, &packetString)
  if err != nil {
		log.Printf( "EE problem in test suite" )
	}
  packet = TsPacket{byteChunk: packetBytes}
  packet.Read()
  if packet.sync != 0X47 { // 71 = 'G' = 0x47
		t.Error("Transport Stream Packet read " +
			"sync byte incorrectly.")
	}
	if packet.transportError != false {
		t.Error("Transport Stream Packet read " +
			"transport error bit incorrectly.")
	}
	if packet.unitStart != false {
		t.Error("Transport Stream Packet read " +
			"unit start incorrectly.")
	}
	if packet.priority != false {
		t.Error("Transport Stream Packet read " +
			"priority incorrectly.")
	}
	if packet.pid != 1 {
		t.Error("Transport Stream Packet read " +
			"packet id incorrectly.")
	}
	if packet.scramble != 0 {
		t.Error("Transport Stream Packet read " +
			"scrambling control incorrectly.")
	}
	if packet.hasAdaptation != false {
		t.Error("Transport Stream Packet read " +
			"adaptation indiciator incorrectly.")
	}
	if packet.hasPayload != true {
		t.Error("Transport Stream Packet read " +
			"payload indicator incorrectly.")
	}
	if packet.continuity != 7 {
		t.Error("Transport Stream Packet read " +
			"continuity incorrectly.")
	}

	// PAT
  awfulStateSetup()
  packetString = `
  4740 0010 0000 b00d 0001 c100 0000 01f0
  002a b104 b2ff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff`
  err = generateBytesFromString(&packetBytes, &packetString)
  if err != nil {
		log.Printf( "EE problem in test suite" )
	}
  packet = TsPacket{byteChunk: packetBytes}
  packet.Read()
  if packet.sync != 0X47 { // 71 = 'G' = 0x47
		t.Error("Transport Stream Packet read " +
			"sync byte incorrectly.")
	}
  if packet.transportError != false {
		t.Error("Transport Stream Packet read " +
			"transport error bit incorrectly.")
	}
	if packet.unitStart != true {
		t.Error("Transport Stream Packet read " +
			"unit start incorrectly.")
	}
	if packet.priority != false {
		t.Error("Transport Stream Packet read " +
			"priority incorrectly.")
	}
	if packet.pid != 0 {
		t.Error("Transport Stream Packet read " +
			"packet id incorrectly.")
	}
	if packet.scramble != 0 {
		t.Error("Transport Stream Packet read " +
			"scrambling control incorrectly.")
	}
	if packet.hasAdaptation != false {
		t.Error("Transport Stream Packet read " +
			"adaptation indiciator incorrectly.")
	}
	if packet.hasPayload != true {
		t.Error("Transport Stream Packet read " +
			"payload indicator incorrectly.")
	}
	if packet.continuity != 0 {
		t.Error("Transport Stream Packet read " +
			"continuity incorrectly.")
	}

	// incorrect sync byte on PAT
  packetString = `
  4640 0010 0000 b00d 0001 c100 0000 01f0
  002a b104 b2ff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff ffff ffff
  ffff ffff ffff ffff ffff ffff`
  err = generateBytesFromString(&packetBytes, &packetString)
  if err != nil {
		log.Printf( "EE problem in test suite" )
	}
  packet = TsPacket{byteChunk: packetBytes}
  packet.Read()
  if packet.sync == 0X47 { // 71 = 'G' = 0x47
		t.Error("Transport Stream Packet read " +
			"sync byte incorrectly.")
	}
  if packet.transportError != false {
		t.Error("Transport Stream Packet read " +
			"transport error bit incorrectly.")
	}
	if packet.unitStart != false {
		t.Error("Transport Stream Packet read " +
			"unit start incorrectly.")
	}
	if packet.priority != false {
		t.Error("Transport Stream Packet read " +
			"priority incorrectly.")
	}
	if packet.pid != 0 {
		t.Error("Transport Stream Packet read " +
			"packet id incorrectly.")
	}
	if packet.scramble != 0 {
		t.Error("Transport Stream Packet read " +
			"scrambling control incorrectly.")
	}
	if packet.hasAdaptation != false {
		t.Error("Transport Stream Packet read " +
			"adaptation indiciator incorrectly.")
	}
	if packet.hasPayload != false {
		t.Error("Transport Stream Packet read " +
			"payload indicator incorrectly.")
	}
	if packet.continuity != 0 {
		t.Error("Transport Stream Packet read " +
			"continuity incorrectly.")
	}
}

func TestInit(t *testing.T) {
	var rc bool
	globals_initialized = false
	// todo( mathew guest ) assert objects empty
  rc = Init()
  if rc == false {
  	t.Error("initial Init() failed")
  }
  rc = Init()
  // todo( mathew guest ) assert objects unchanged
  if rc == true {
  	t.Error("secondary Init() returned success when it should have failed")
  }
}
