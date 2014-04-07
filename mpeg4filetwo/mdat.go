package mpeg4file

import (
	"mpeg4file/box"
)

type Mdat struct{
	*box.Box
	data []byte
}

func NewMdat (s uint64, box uint32, payload []byte) Mdat {
	Mdat:= {&box.Box{0,0,box},
			payload}
	Mdat.SetSize(s)
	return Mdat
}

func (m *Mdat) SetSize (s uint64){
	if s==0{
		m.Size=0
	} else {
		if s>4294967295 {
			m.Size = uint32(s)
		}else{
			m.Size = 1
			m.LargeSize = s
		}
	}
}