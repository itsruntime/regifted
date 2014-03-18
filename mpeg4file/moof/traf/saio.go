package traf

import "strconv"

type saio struct{
	//box
	size uint32
	largeSize uint64
	boxType uint32
	// Fullbox
	version uint8
	//flags [3]byte doesn't work with line 27
	flags uint8
	//saio box
	auxInfoType uint32
	auxInfoTypeParameter uint32
	entryCount uint32
	offset *uint32
	largeOffset *uint64 
}

func NewSaio(s uint64, flag uint8, ver uint8, entry uint32){
	//Need to finallize this constructor 
	newSaio:=new(saio)
	newSaio.SetSize(s)
	newSaio.boxType = 0x7361696F
	newSaio.flags = flag
	if (newSaio.flags & 1)>1 {
		//aux_info
	}
	newSaio.version = ver
	newSaio.entryCount = entry
	if newSaio.version == 0 {
		offset := make([]uint32,0,newSaio.entryCount)
	} else {
		largeOffset := make([]uint64,0,newSaio.entryCount)
	}
}

func (s *saio) SetSize(size uint64) {
	if size==0{
		s.size = 0
	}else if size > 4294967295 {
		s.size = 1
		s.largeSize = size
	} else {
		s.size = uint32(size)
	}
}

func (s *saio) String() string{
	return strconv.FormatUint(uint64(s.size),10)
}