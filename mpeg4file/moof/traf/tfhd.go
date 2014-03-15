 package traf

import "strconv"

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