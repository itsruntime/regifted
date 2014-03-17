package moov

//Contains all the boxes that have information about
//a track's media data.

type Mdia struct {
	//extends Box
	size		uint32
	largesize	uint64
	boxtype		uint32
}