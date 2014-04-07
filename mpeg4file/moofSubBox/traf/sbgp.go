package traf

import "strconv"

type sbgp struct{
	size uint32
	largeSize uint64
	boxType uint32
	// The version of the box. either 0 or 1
	version uint8
	flags [3]byte
	// sbgp box specific elements
	groupingType uint32
	groupingTypeParameter uint32
	entryCount uint32
	sampleCount uint32
	groupDescriptionIndex uint32
}

func (s *sbgp) SetSize(size uint64) {
	if size==0{
		s.size = 0
	}else if size > 4294967295 {
		s.size = 1
		s.largeSize = s
	} else {
		s.size = uint32(s)
	}
}

func (sb *sbgp) String() string{
	return strconv.FormatUint(uint64(sb.size),10)
}