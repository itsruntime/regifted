package moov

//There are four track types: video, sound, hint, and null. 
//Each type has its own media header. Each header corresponds
//to one media handler type.

//The video media header contains general presentation 
//information.

type Vmhd struct {
	//extends FullBox
	size		uint32
	largeSize	uint64 //size==1 if used
	boxtype		uint32
	flags		uint32 //flags all set to 1
	version		uint8 //seems to be set to 0 in the spec
	//specifies a composition mode for this video track
	//graphicsMode		uint16 //template, = 0 in spec
	//set of 3 color values(rgb) available for graphics mode
	//opcolor	[3]uint16 //template, = {0,0,0} in spec
}