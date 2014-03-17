package moov

//There are four track types: video, sound, hint, and null. 
//Each type has its own media header. Each header corresponds
//to one media handler type.

//For streams other than visual and audio.

type Smhd struct {
	//extends FullBox
	size		uint32
	largeSize	uint64 //size==1 if used
	boxtype		uint32
	flags		uint32 //set all to 0
	version		uint8 //seems to be set to 0 in the spec
}