package ts

import (
	"regifted/data"

	"fmt"
)

type Program struct {
	byteChunk []byte

	pid    uint
	number uint
}

//Program Read
//number – Program_number is a 16-bit field. It specifies the program to which the program_map_PID is
//applicable. When set to 0x0000, then the following PID reference shall be the network PID. For all other cases the value
//of this field is user defined. This field shall not take any single value more than once within one version of the Program
//Association Table
//
//pid – The program_map_PID is a 13-bit field specifying the PID of the Transport Stream packets
//which shall contain the program_map_section applicable for the program as specified by the program_number. No
//program_number shall have more than one program_map_PID assignment. The value of the program_map_PID is
//defined by the user, but shall only take values
func (program *Program) Read(reader *data.Reader) {
	logger.Debug("ts - Program.Read()")
	program.number = reader.Read(2)
	program.pid = reader.Read(2) & 0x1fff
	program.Print()
}

func (program *Program) Print() {
	fmt.Println("\n:::Program:::\n")
	fmt.Println("pid = ", program.pid)
	fmt.Println("number = ", program.number)
}
