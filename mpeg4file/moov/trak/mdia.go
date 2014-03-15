package moov
//The media declaration container contains all the objects that //declare information about the media data within a 
//track. 

type Mdia struct {
	//extends Box
	size		uint32
	largesize	uint64
	boxtype		uint32
}