package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

type Mfhd struct {
	*FullBoxFields
	SequenceNumber uint32
}

func NewMfhd(s uint32, ver uint8, flag []byte, sqn uint32) *Mfhd {
	newMfhd := &Mfhd{&FullBoxFields{
		&BoxFields{s, 0x6d666864},
		ver,
		flag},
		sqn}
	return newMfhd
}

func (m *Mfhd) SetSize(s uint32) {
	m.Size = s
}

func (m *Mfhd) Write() []byte {
	buf := new(bytes.Buffer)
	var err error
	// Size
	err = binary.Write(buf, binary.BigEndian, m.Size)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf, binary.BigEndian, m.BoxType)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//version
	err = binary.Write(buf, binary.BigEndian, m.Version)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//flags
	err = binary.Write(buf, binary.BigEndian, m.Flags)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//sequence number
	err = binary.Write(buf, binary.BigEndian, m.SequenceNumber)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
