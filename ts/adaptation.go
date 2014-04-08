package ts

import (
	"regifted/data"
	"regifted/util"
	"regifted/util/mylog"

	"fmt"
)

type Adaptation struct {
	payload   []byte
	byteChunk []byte

	size     uint
	pcr      Pcr
	opcr     Pcr
	splice   uint
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

//Adaptation Read
//size - Number of bytes in the adaptation field immediately following this byte (8 bits)
//
//discontinuity - Set to 1 if current TS packet is in a discontinuity state with respect to either the continuity counter or the program clock reference (1 bit)
//
//random - Set to 1 if the PES packet in this TS packet starts a video/audio sequence (1 bit)
//
//priority - 1 = higher priority
//
//hasPCR - 1 means adaptation field does contain a PCR field
//
//hasOPCR - 1 means adaptation field does contain an OPCR field
//
//hasSplice - 1 means presence of splice countdown field in adaptation field
//
//hasPrivate - 1 means presence of private data bytes in adaptation field
//
//hasExtension - 1 means presence of adaptation field extension
func (adaptation *Adaptation) Read() {
	if adaptation.byteChunk == nil {
		logger.Error("Adaptation.Read() was called with a nil payload")
		return
	}
	logger.Debug("Adaptation.Read() - attempting to process Adaptation data that's already loaded")
	if logger.IsWithinSeverity(mylog.SEV_TRACE) {
		logger.Trace("Adaptation.Read() - Adaptation payload: %s", util.SprintfHex(adaptation.byteChunk))
	}

	var flags uint = 0
	var spliceFlag int = 0
	var pcrFlag int = 0
	var opcrFlag int = 0

	reader := data.NewReader(adaptation.byteChunk)

	adaptation.size = reader.Read(1)

	flags = reader.Read(1)

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
		adaptation.pcr.byteChunk = reader.ReadBytes(6)
		adaptation.pcr.Read()
	}

	if adaptation.hasOPCR {
		opcrFlag = 6
		adaptation.pcr.byteChunk = reader.ReadBytes(6)
		adaptation.opcr.Read()

	}

	if adaptation.hasSplice {
		spliceFlag = 1
		adaptation.splice = reader.Read(1)

	}
	adaptation.stuffing = int(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)
	reader.Cursor += uint64(int(adaptation.size) - 1 - pcrFlag - opcrFlag - spliceFlag)
	payload := reader.ReadBytes(reader.Size - reader.Cursor)
	adaptation.payload = payload
	adaptation.Print()
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
