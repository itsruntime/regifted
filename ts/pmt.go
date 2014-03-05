package main

import (
	"fmt"
	"log"
	"regifted/data"
)

type Pmt struct {
	byteChunk              []byte
	pointerField           bool
	unitStart              bool
	tableId                uint
	sectionSyntaxIndicator bool
	sectionLength          uint
	programNumber          uint
	versionNumber          uint
	sectionNumber          uint
	lastSectionNumber      uint

	pcrPid            uint
	programInfoLength uint
	descriptor        []byte
	count             uint
	entries           []PmtEntry

	types map[uint32]uint
}

type PmtEntry struct {
	byteChunk []byte

	streamType uint
	pid        uint
	infoLength uint
	descriptor []byte
}

//Pmt Read
//tableId – This is an 8-bit field, which in the case of a TS_program_map_section shall be always set to 0x02
//
//sectionSyntaxIndicator – The section_syntax_indicator is a 1-bit field which shall be set to '1'.
//
//sectionLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the section starting immediately following the section_length field, and including the CRC. The value in this
//field shall not exceed 1021 (0x3FD).
//
//programNumber – program_number is a 16-bit field. It specifies the program to which the program_map_PID is
//applicable. One program definition shall be carried within only one TS_program_map_section. This implies that a
//program definition is never longer than 1016 (0x3F8). See Informative Annex C for ways to deal with the cases when
//that length is not sufficient. The program_number may be used as a designation for a broadcast channel, for example. By
//describing the different program elements belonging to a program, data from different sources (e.g. sequential events)
//can be concatenated together to form a continuous set of streams using a program_number.
//
//versionNumber – This 5-bit field is the version number of the TS_program_map_section. The version number shall be
//incremented by 1 modulo 32 when a change in the information carried within the section occurs. Version number refers
//to the definition of a single program, and therefore to a single section. When the current_next_indicator is set to '1', then
//the version_number shall be that of the currently applicable TS_program_map_section. When the current_next_indicator
//is set to '0', then the version_number shall be that of the next applicable TS_program_map_section.
//
//sectionNumber – The value of this 8-bit field shall be 0x00.
//
//lastSectionNumber – The value of this 8-bit field shall be 0x00.
//
//pcrPid – This is a 13-bit field indicating the PID of the Transport Stream packets which shall contain the PCR fields
//valid for the program specified by program_number. If no PCR is associated with a program definition for private
//streams, then this field shall take the value of 0x1FFF.
//
//programInfoLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the
//number of bytes of the descriptors immediately following the program_info_length field.
func (pmt *Pmt) Read() {
	if pmt.byteChunk == nil {
		log.Printf("attempted to read from nil pointer: byteChunk\n")
		return
	}

	var CRC_SIZE uint = 4
	var SKIP_BYTES uint = 9
	var flags uint = 0

	var flag bool = false

	reader := data.NewReader(pmt.byteChunk)

	if reader.Read(1) == 1 {
		flag = true
	}

	pmt.pointerField = (pmt.unitStart && flag) || false

	pmt.tableId = reader.Read(1)

	flags = reader.Read(2)

	pmt.sectionSyntaxIndicator = flags&0x8000 > 0
	pmt.sectionLength = flags & 0x3ff

	pmt.programNumber = reader.Read(2)

	pmt.versionNumber = reader.Read(1)

	pmt.sectionNumber = reader.Read(1)

	pmt.lastSectionNumber = reader.Read(1)

	pmt.pcrPid = reader.Read(2) & 0x1fff

	pmt.programInfoLength = reader.Read(2) & 0x3ff

	pmt.descriptor = reader.ReadBytes(uint64(pmt.programInfoLength))

	pmt.count = pmt.sectionLength - SKIP_BYTES - pmt.programInfoLength

	pmt.Print()

	for pmt.count > CRC_SIZE {

		pmtEntry := PmtEntry{}

		pmtEntry.Read(reader)

		pmt.entries = append(pmt.entries, pmtEntry)
		types[pmtEntry.pid] = pmtEntry.streamType
		elementaryConstructors[pmtEntry.pid] = ElementaryStreamPacket{}

		pmt.count -= (5 + pmtEntry.infoLength)

	}

}

//PmtEntry Read
//streamType – This is an 8-bit field specifying the type of program element carried within the packets with the PID
//whose value is specified by the elementary_PID.
//
//pid – This is a 13-bit field specifying the PID of the Transport Stream packets which carry the associated
//program element.
//
//infoLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the descriptors of the associated program element immediately following the ES_info_length field.
func (pmtEntry *PmtEntry) Read(reader *data.Reader) {

	pmtEntry.streamType = reader.Read(1)

	pmtEntry.pid = reader.Read(2) & 0x1fff

	pmtEntry.infoLength = reader.Read(2) & 0x3ff

	pmtEntry.descriptor = reader.ReadBytes(uint64(pmtEntry.infoLength))

	pmtEntry.Print()

}

func (pmt *Pmt) Print() {

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

}

func (pmtEntry *PmtEntry) Print() {
	fmt.Println("\n:::PmtEntry:::\n")
	fmt.Println("pid = ", pmtEntry.pid)
	fmt.Println("streamType = ", pmtEntry.streamType)
	fmt.Println("infoLength = ", pmtEntry.infoLength)
	fmt.Println("descriptor = ", pmtEntry.descriptor)

}
