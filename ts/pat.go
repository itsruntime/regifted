package ts

import (
	"regifted/data"

	"fmt"
	"log"
)

type Pat struct {
	SKIP_BYTES int
	CRC_SIZE   int

	byteChunk []byte

	unitStart bool

	pointerField           bool
	tableId                uint
	flags                  uint
	sectionSyntaxIndicator bool
	sectionLength          uint
	transportStreamId      uint
	versionNumber          uint
	currentNext            uint
	sectionNumber          uint
	lastSectionNumber      uint
	count                  uint

	programs []Program
}

//Pat Read
//table_id – This is an 8-bit field, which shall be set to 0x00 a
//
//sectionSyntaxIndicator – The section_syntax_indicator is a 1-bit field which shall be set to '1'.
//
//sectionLength – This is a 12-bit field, the first two bits of which shall be '00'. The remaining 10 bits specify the number
//of bytes of the section, starting immediately following the section_length field, and including the CRC. The value in this
//field shall not exceed 1021 (0x3FD).
//
//transportStreamId – This is a 16-bit field which serves as a label to identify this Transport Stream from any other
//multiplex within a network. Its value is defined by the user.
//
//versionNumber – This 5-bit field is the version number of the whole Program Association Table. The version number
//shall be incremented by 1 modulo 32 whenever the definition of the Program Association Table changes. When the
//current_next_indicator is set to '1', then the version_number shall be that of the currently applicable Program Association
//Table. When the current_next_indicator is set to '0', then the version_number shall be that of the next applicable Program
//Association Table.
//
//currentNext – A 1-bit indicator, which when set to '1' indicates that the Program Association Table sent is
//currently applicable. When the bit is set to '0', it indicates that the table sent is not yet applicable and shall be the next
//table to become valid.
//
//sectionNumber – This 8-bit field gives the number of this section. The section_number of the first section in the
//Program Association Table shall be 0x00. It shall be incremented by 1 with each additional section in the Program
//Association Table.
//
//lastSectionNumber – This 8-bit field specifies the number of the last section (that is, the section with the highest
//section_number) of the complete Program Association Table.
func (pat *Pat) Read() {
	if pat.byteChunk == nil {
		log.Printf("attempted to read from nil pointer: byteChunk\n")
		return
	}

	var SKIP_BYTES uint = 5
	var CRC_SIZE uint = 4
	var PROGRAM_SIZE uint = 4
	var flags uint = 0
	var flag bool = false

	reader := data.NewReader(pat.byteChunk)
	if reader.Read(1) == 1 {
		flag = true
	}

	pat.pointerField = (pat.unitStart && flag) || false
	pat.tableId = reader.Read(1)
	flags = reader.Read(2)
	pat.sectionSyntaxIndicator = flags&0x8000 > 0
	pat.sectionLength = flags & 0x3ff
	pat.transportStreamId = reader.Read(2)
	flags = reader.Read(1)
	pat.versionNumber = flags & 0x3ffe
	pat.currentNext = flags & 0x0001
	pat.sectionNumber = reader.Read(1)
	pat.lastSectionNumber = reader.Read(1)
	pat.count = pat.sectionLength - SKIP_BYTES

	pat.Print()

	for pat.count > CRC_SIZE {
		program := Program{}
		program.Read(reader)
		pat.programs = append(pat.programs, program)
		pat.count = pat.count - PROGRAM_SIZE
	}
}

// loads a PAT into a TS State object
func (state *TSState) loadPAT(pat *Pat) {
	for idx, program := range pat.programs {
		_ = idx
		state.pmtConstructors[program.pid] = Pmt{}
	}
}

func (pat *Pat) Print() {
	fmt.Println("\n:::Pat:::\n")
	fmt.Println("tableId = ", pat.tableId)
	fmt.Println("pointerField = ", pat.pointerField)
	fmt.Println("sectionSyntaxIndicator = ", pat.sectionSyntaxIndicator)
	fmt.Println("sectionLength = ", pat.sectionLength)
	fmt.Println("transportStreamId = ", pat.transportStreamId)
	fmt.Println("versionNumber = ", pat.versionNumber)
	fmt.Println("currentNext = ", pat.currentNext)
	fmt.Println("sectionNumber = ", pat.sectionNumber)
	fmt.Println("lastSectionNumber = ", pat.lastSectionNumber)
	fmt.Println("count = ", pat.count)

}
