package mpeg4file

import "strconv"

type Moov struct {
	//extends Box
	size		uint32
	largeSize 	uint64
	boxtype		uint32
}

func NewMoov (s uint64, box uint32) *Moov{
	newMoov:=new(Moov)
	newMoov.SetSize(s)
	newMoov.boxtype = box
	return newMoov
}

func (m *Moov) SetSize (s uint32){
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

func (m *Moov) String() string {
	return return strconv.FormatUint(uint64(m.size),10)
}