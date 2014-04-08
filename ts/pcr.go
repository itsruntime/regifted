package ts

import (
	"regifted/data"

	"fmt"
	"log"
)

type Pcr struct {
	pcr       uint
	ext       uint
	byteChunk []byte

	a  uint
	b  uint
	c  uint
	d  uint
	ef uint
	e  uint
	f  uint
	g  uint
}

//PCR Read
//PCR fields
//valid for the program specified by program_number
func (pcr *Pcr) Read() {
	if pcr.byteChunk == nil {
		log.Printf("attempted to read from nil pointer: byteChunk\n")
		return
	}
	reader := data.NewReader(pcr.byteChunk)

	pcr.a = reader.Read(1)

	pcr.b = reader.Read(1)

	pcr.c = reader.Read(1)

	pcr.d = reader.Read(1)

	pcr.ef = reader.Read(1)

	pcr.e = pcr.ef & 0x80
	if pcr.e > 0 {
		pcr.e = 1
	} else {
		pcr.e = 0
	}

	pcr.f = pcr.ef & 0x01

	pcr.g = reader.Read(1)

	pcr.pcr = (pcr.a << 25) | (pcr.b << 17) | (pcr.c << 9) | (pcr.d << 1) | (pcr.e&1 | 0)
	pcr.ext = (pcr.f << 8) | pcr.g

	pcr.Print()

}

func (pcr *Pcr) Print() {

	fmt.Println("\n:::Pcr:::\n")
	fmt.Println("a = ", pcr.a)
	fmt.Println("b = ", pcr.b)
	fmt.Println("c = ", pcr.c)
	fmt.Println("d = ", pcr.d)
	fmt.Println("ef = ", pcr.ef)
	fmt.Println("e = ", pcr.e)
	fmt.Println("f = ", pcr.f)
	fmt.Println("g = ", pcr.g)
	fmt.Println("pcr = ", pcr.pcr)
	fmt.Println("ext = ", pcr.ext)
}
