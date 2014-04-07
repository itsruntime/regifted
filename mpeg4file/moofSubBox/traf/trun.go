package traf

import "strconv"

type trun struct{
	size uint32
	largeSize uint64
	boxType uint32
}

func NewTrun(s uint64){
	newTrun := new(trun)
	newTrun.SetSize(s)
	newTrun.boxType = 0x7472756E
	return newTrun
}

func (t *trun) SetSize(s uint64) {
	if s==0{
		t.size = 0
	}else if s > 4294967295 {
		t.size = 1
		t.largeSize = s
	} else {
		t.size = uint32(s)
	}
}

func (t *trun) String() string{
	return strconv.FormatUint(uint64(t.size),10)
}

func (t *trun) Write(f *File) {
	// Size
	// BoxType
	// Contained samples write
}

// Make struct for Sample information