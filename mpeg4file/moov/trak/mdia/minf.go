package moov

//This box contains all the objects that declare
//characteristic information of the media in the track

type Minf struct {
	//extends Box
	size		uint32
	largesize	uint64
	boxtype		uint32
}