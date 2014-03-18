package traf

import "strconv"

type saiz struct{
	// box
	size uint32
	largeSize uint64
	boxType uint32
	// fullbox
	flags [3]byte
	version uint8
	// saiz box
	auxInfoType uint32
	auxInfoTypeParameter uint32
	defaultSampleInfoSize uint8
	sampleCount uint32
	sampleInfoSize []uint8
}

func (s *saiz) String() string{
	return strconv.FormatUint(uint64(s.size),10)
}