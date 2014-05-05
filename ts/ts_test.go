package ts

import (
	"fmt"
	// "io/ioutil"
	"log"
	//"os"
	"regifted/util"
	"testing"
)


// const TESTFP = "test1.ts"

// func TestReader(t *testing.T) {
// 	// compare file size from os and then counted from buffer
// 	fp, err := os.Open(TESTFP)
// 	if err != nil {
// 		t.Fatalf("did not open test file\n")
// 	}
// 	fi, err := fp.Stat()
// 	if err != nil {
// 		// Could not obtain stat, handle error
// 	}
// 	fileSizeOS := fi.Size()
// 	fp.Close()

// 	bytes, err := ioutil.ReadFile(TESTFP)
// 	if err != nil {
// 		t.Fatalf("did not open test file\n")
// 	}
// 	fileSizeBuff := len(bytes)
// 	if uint64(fileSizeOS) != uint64(fileSizeBuff) {
// 		t.Error("file size calculated from buffer does not match size returned from OS\n")
// 	}

// 	// t.Error( "print" )
// }

func TestTransportPacketRead(t *testing.T) {
	InitLogger()
	var rc bool
	state := &TSState{}
	rc = state.init()
	_ = rc

	var packetBytes []byte
	var packetString string
	var err error
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
	err = util.GenerateBytesFromString(&packetBytes, &packetString)
	if err != nil {
		log.Printf("EE problem in test suite")
	}
	fmt.Println(packetBytes)
	packet := TsPacket{byteChunk: packetBytes}
	_ = packet

	var type_ int
	type_, _ = state.ReadTSPacket(&packet)
	// if type_ != PACKET_TYPE_PAT {
	// 	t.Error( "packet type read incorrectly" )
	// }
	_ = type_

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

	// PAT
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
		err = util.GenerateBytesFromString(&packetBytes, &packetString)
		if err != nil {
			log.Printf("EE problem in test suite")
		}
		packet = TsPacket{byteChunk: packetBytes}
		type_, _ = state.ReadTSPacket(&packet)
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
	err = util.GenerateBytesFromString(&packetBytes, &packetString)
	if err != nil {
		log.Printf("EE problem in test suite")
	}
	packet = TsPacket{byteChunk: packetBytes}
	type_, _ = state.ReadTSPacket(&packet)
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
		var state *TSState

		state = &TSState{}
		rc = state.init()
		if rc == false {
			t.Error("initial init() failed")
		}
}
