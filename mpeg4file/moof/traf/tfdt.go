package traf

import "strconv"

// 
type tfdt struct{
	// box

	// The size of this box. If size = 1 then the
	// size of this box is stored in largeBox.
	size uint32

	// The size of this box. Used when uint32 is 
	// inadequate space for the size. 
	largeSize uint64

	// The name of the box (tfdt) stored as a hex
	// decimal number. 
	boxType uint32
	
	// fullbox

	// The version of this box either 0 or 1
	version uint8

	// Unused by this box. Is inharted from fullbox
	flags [3]byte
	
	// tfdt components 

	// Equals the sum of the decode durations of 
	// all earlier samples in the media, expressed
	// in the media's timescale. 
	baseMediaDecodeTime uint32

	// Equals the sum of the decode durations of 
	// all earlier samples in the media, expressed
	// in the media's timescale.
	largeBaseMediaDecodeTime uint64 
}

func (t *tfdt) String() string{
	return strconv.FormatUint(uint64(t.size),10)
}