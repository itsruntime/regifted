package moov

//There are four track types: video, sound, hint, and null. 
//Each type has its own media header. Each header corresponds
//to one media handler type.

//The sound media header contains general presentation
//information for audio media. For audio tracks.

type Smhd struct {
	//extends FullBox
	size		uint32
	largeSize	uint64 //size==1 if used
	boxtype		uint32
	flags		uint32 //no note on how to handle flags
	version		uint8 //seems to be set to 0 in the spec
	//is a fixed-point 8.8 number that places mono
	//audio tracks in a stereo space; 0 is centre (the
	//normal value); full left is -1.0 and full right is 1.0
	//balance		uint16 //template, = 0 in spec
	const reserved uint16 //set at 0 in spec
}