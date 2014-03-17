package moov

//The data information box points to the track's media data.

type Dinf struct {
	//extends Box
	size		uint32
	largesize	uint64 //size==1 if used
	boxtype		uint32
}