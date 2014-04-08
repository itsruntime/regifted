package mpeg4file

import (
		"strconv"
		"encoding/binary"
		"fmt"
		"bytes"
		)

type moof struct{
	size uint32
	largeSize uint64
	boxType uint32
}

func NewMoof(s uint64, box uint32) *moof{
	newMoof:=new(moof)
	newMoof.SetSize(s)
	newMoof.boxType = 0x6d6f6f66
	return newMoof
}

func (m *moof) SetSize(s uint64){
	if s==0 {
		m.size = 0
	}else if s>4294967295 {
		m.size = 1
		m.largeSize = s
	} else {
		m.size = uint32(s)
	}
}

func (m *moof) String() string{
	return strconv.FormatUint(uint64(m.size),10)
}

func (m *moof) Write() []byte{
	buf := new(bytes.Buffer)
	var err error
	// Size
	err=binary.Write(buf, binary.BigEndian, m.size)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf,binary.BigEndian,m.boxType)
	// Contained boxes write
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}