package moov

//MovieExtendsBox is abox that indicates there might be Movie Fragment
//Boxes in the file. To know of all the samples in the tracks, these Movie
//Fragment Boxes must be found and scanned in order, and their information
//logically added to that found in the Movie Box. It sounds like the
//Fragments are constructed into a Moov and its Traks?
type Mvex struct {
	//extends Box('mvex')
	size			uint32
	largeSize		uint64
	boxtype			uint32
}