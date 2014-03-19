package traf

import "strconv"

type subs struct{
	// box

	// The size of this box. If size = 1 then the
	// size of this box is stored in largeBox.
	size uint32

	// The size of this box. Used when uint32 is 
	// inadequate space for the size. 
	largeSize uint64

	// The name of the box (subs) stored as a hex
	// decimal number.
	boxType uint32
	// fullbox

	// The version of this box either 0 or 1
	version uint8

	// Unused by this box. Is inharted from fullbox
	flags [3]byte
	// traf

	
	entryCount uint32
	sampleDelta uint32
	subabmpleCount uint32
	largeSubsampleSize []uint32
	subsampleSize []uint16
	subsamplePriority []uint8
	discardable []uint8
	resvered []uint32
}

func (s *subs) String() string{
	return strconv.FormatUint(uint64(s.size),10)
}