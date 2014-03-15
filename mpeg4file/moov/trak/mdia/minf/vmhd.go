package moov

//There is a different media information header for each //track type (corresponding to the media handler-type); 
//the matching header shall be present, which may be one //of those defined here, or one defined in a derived 
//specification.

//One type of mediaheader must be present
//'vmhd', 'smhd', 'hmhd', 'nmhd'

//The video media header contains general presentation //information, independent of the coding, for video 
//media. Note that the flags field has the value 1

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