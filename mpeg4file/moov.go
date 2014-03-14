package moov

type Moov struct {
	//extends Box
	size		uint32
	largeSize 	uint64
	boxtype		uint32
}

func NewMoov (s uint32, box uint32) *Moov{
	newMoov:=new(Moov)
	newMoov.size = s
	newMoov.boxtype = box
	return newMoov
}

func NewMoovWithLargeSize (s uint64, box uint32) *Moov{
	newMoov:=new(Moov)
	newMoov.size = 1
	newMoov.largeSize = s
	newMoov.boxtype = box
	return newMoov
}

func (moov *Moov) SetSize (s uint32){
	moov.size = s
}

func (moov *Moov) SetLargeSize (s uint64){
	moov.size = 1
	moov.largeSize = s
}