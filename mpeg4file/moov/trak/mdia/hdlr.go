package moov

//This box within a Media Box declares the process by //which the media-data in the track is presented, and thus, 
//the nature of the media in a track. For example, a //video track would be handled by a video handler.

type Hdlr struct {
	extends FullBox
	size			uint32
	largeSize		uint64 //used if size==1
	boxtype			uint32
	flags			uint32 //bit(24) in spec
	version			uint8 //an integer that specifies the version of the box; seems like it should always be 0
	pre_defined		uint32 //set to zero in the spec
	//an integer containing one of the following values:
	//‘vide’  Video track 
	//‘soun’  Audio track 
	//‘hint’	Hint track 
	//‘meta’ 	Timed Metadata track 
	//‘auxv’ 	Auxiliary Video track
	handlerType		uint32 
	const reserved 	[3]uint32
	//human readable name for the track type; utf-8 chars
	name			string
}