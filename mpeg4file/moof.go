package mpeg4file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"regifted/box"
	//"strconv"
)

type moof struct {
	*box.Box
}

func NewMoof(s uint32) *moof {
	//newMoof := &moof{&box.Box{}} // creates and empty struct
	// to keep this the same, it must get the box data from Box
	// and assign it to the moof variables struct; then the handling
	// of common Box fields is outside of the moof box's control
	newMoof := &moof{&box.Box{Size: s, BoxType: 0x6d6f6f66}}
	return newMoof
}

func (m *moof) SetSize(s uint32) {
	m.Size = s
}

func (m *moof) Write() []byte {
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
	return buf.Bytes()
}
