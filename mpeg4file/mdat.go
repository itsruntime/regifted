package mpeg4file

type mdat struct{
	size uint32
	largeSize uint64
	boxtype uint32
	data []byte
}

func NewMdat (s uint32, payload []byte) *mdat{
	newMdat:=new(mdat)
	newMdat.size=s
	newMdat.data = payload
	return newMdat
}

func NewMdatLargeSize (s uint64, payload []byte) *mdat{
	newMdat:=new(mdat)
	newMdat.size=1
	newMdat.largeSize = s
	newMdat.data = payload
	return newMdat
}

func (m *mdat) SetSize (s uint64){
	if s>4294967295 {
		m.size = uint32(s)
	}else{
		m.size = 1
		m.largeSize = s
	}
}