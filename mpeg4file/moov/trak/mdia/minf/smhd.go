package moov

//There is a different media information header for each //track type (corresponding to the media handler-type); 
//the matching header shall be present, which may be one //of those defined here, or one defined in a derived 
//specification.

//One type of mediaheader must be present
//'vmhd', 'smhd', 'hmhd', 'nmhd'

//The sound media header contains general presentation
//information, independent of the coding, for audio 
//media. This header is used for all tracks containing
//audio.

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