package ts

import (
	"regifted/data"
	"regifted/util"
	"regifted/util/mylog"

	"fmt"
)

type Pes struct {
	byteChunk []byte

	pid          uint
	streamType   uint
	streamId     uint
	packetLength uint
	flags        uint
	pts          uint
	dts          uint
	Payload      []byte
}

//Pes Read
//prefix – The packet_start_code_prefix is a 24-bit code. Together with the stream_id that follows, it
//constitutes a packet start code that identifies the beginning of a packet. The packet_start_code_prefix is the bit string
//'0000 0000 0000 0000 0000 0001' (0x000001 in hexadecimal).
//
//streamId – This 8-bit field shall have a value '1111 1111' (0xFF).
//
//packetLength – The PES_packet_length is a 16-bit field indicating the total number of bytes in the program_stream_directory immediately following this field
func (pes *Pes) Read() {
	if pes.byteChunk == nil {
		logger.Error("PES.Read() was called with a nil payload")
		return
	}
	logger.Debug("PES.Read() - attempting to process PES data that's already loaded")
	if logger.IsWithinSeverity(mylog.SEV_TRACE) {
		logger.Trace("PES.Read() - PES payload: %s", util.SprintfHex(pes.byteChunk))
	}

	reader := data.NewReader(pes.byteChunk)

	var prefix uint
	var headerLength uint
	var headerData []byte
	var flags uint

	prefix = reader.Read(3)
	if prefix == uint(0x000001) {
		pes.streamId = reader.Read(1)
		pes.packetLength = reader.Read(2)
		flags = reader.Read(2)
		headerLength = reader.Read(1)
		headerData = reader.ReadBytes(uint64(headerLength))
		if (flags & 0x0080) == 1 {
			pes.pts = ReadHeaderData(headerData)
		}
		if (flags & 0x0040) == 1 {
			pes.dts = ReadHeaderData(headerData)
		}
		pes.Payload = reader.ReadBytes(reader.Size - reader.Cursor)
	}
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
	fmt.Println("//payload length= ", len(pes.Payload))
	fmt.Println("//nal =  {}") // DELETE
	fmt.Println("////////////////////////////////")

}
