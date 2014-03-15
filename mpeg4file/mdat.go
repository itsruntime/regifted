package mpeg4file

type mdat struct{
	size uint32
	largeSize uint64
	boxtype uint32
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