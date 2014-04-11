package mp4box

import (
		"encoding/binary"
		"fmt"
		"bytes"
		)

type mdat struct{
	size uint32
	largeSize uint64
	boxType uint32
	data []byte
}

func NewMdat (s uint64, payload []byte) *mdat{
	newMdat:=new(mdat)
	newMdat.SetSize(s)
	newMdat.data = payload
	return newMdat
}

func (m *mdat) SetSize (s uint64){
	if s==0{
		m.size=0
	} else {
		if s>4294967295 {
			m.size = uint32(s)
		}else{
			m.size = 1
			m.largeSize = s
		}
	}
}

func (m *mdat) Write() []byte{
	buf := new(bytes.Buffer)
	var err error
	// Size
	err=binary.Write(buf, binary.BigEndian, m.size)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf,binary.BigEndian,m.boxType)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	buf.Write(m.data)
	return buf.Bytes()
}