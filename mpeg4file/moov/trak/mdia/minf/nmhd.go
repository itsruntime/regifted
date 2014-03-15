package moov

//There is a different media information header for each //track type (corresponding to the media handler-type); 
//the matching header shall be present, which may be one //of those defined here, or one defined in a derived 
//specification.

//One type of mediaheader must be present
//'vmhd', 'smhd', 'hmhd', 'nmhd'

//Streams other than visual and audio (e.g., timed
//metadata streams) may use a null Media Header Box, as 
//defined here

type Smhd struct {
	//extends FullBox
	size		uint32
	largeSize	uint64 //size==1 if used
	boxtype		uint32
	flags		uint32 //set all to 0
	version		uint8 //seems to be set to 0 in the spec
}