 package mp4box

import (
		"strconv"
		"encoding/binary"
		"fmt"
		"bytes"
		)

type TfhdLevel3 interface{
	String() string
}

type tfhd struct{
	size uint32
	largeSize uint64
	boxType uint32
	version uint8
	flags [3]byte
	trackID uint32
	//optional fields
	baseDataOffset uint64
	sampleDescriptionIndex uint32
	defaultSampleDuration uint32
	defaultSampleSize uint32
	defaultSampleFlags uint32
}

func NewTfhd(s uint64, 
				flag [3]byte ,
				trackNum uint32) *tfhd{
	newTfhd := new(tfhd)
	newTfhd.SetSize(s)
	newTfhd.boxType = 0x74666864
	newTfhd.flags = flag
	newTfhd.trackID = trackNum
	return newTfhd
}

func (t *tfhd) SetSize(s uint64){
	if s == 0 {
		t.size = 0 
	} else if s > 4294967295 {
		t.size = 1
		t.largeSize = s
	}else {
		t.size = uint32(s)
	}
}

func (t *tfhd) String() string{
	return strconv.FormatUint(uint64(t.size),10)
}

func (m *tfhd) Write() []byte{
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
	//version
	err = binary.Write(buf,binary.BigEndian,m.version)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	//flags
	err = binary.Write(buf,binary.BigEndian,m.flags)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	//sequnce
	err = binary.Write(buf,binary.BigEndian,m.trackID)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	if m.baseDataOffset != 0{
		err = binary.Write(buf,binary.BigEndian,m.baseDataOffset)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	}
	if m.sampleDescriptionIndex!= 0{
		err = binary.Write(buf,binary.BigEndian,m.sampleDescriptionIndex)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	}
	if m.defaultSampleDuration!= 0{
		err = binary.Write(buf,binary.BigEndian,m.defaultSampleDuration)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	}
	if m.defaultSampleSize!= 0{
		err = binary.Write(buf,binary.BigEndian,m.defaultSampleSize)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	}
	if m.defaultSampleFlags!= 0{
		err = binary.Write(buf,binary.BigEndian,m.defaultSampleFlags)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	}
	return buf.Bytes()
}

















