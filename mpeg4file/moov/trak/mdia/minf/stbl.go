package moov

//Contains all the time and data
//indexing of the media samples in a track: locate
//samples in time, determine their type, determine 
//their size, their container, and their offset into 
//that container.

type Stbl struct {
	//extends Box
	size		uint32
	largesize	uint64 //size==1 if used
	boxtype		uint32
}