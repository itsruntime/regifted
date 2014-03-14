package moov

/*This is a container box for a single track of a presentation. A presentation consists of one or more tracks. 
Each track is independent of the other tracks in the presentation and carries its own temporal and spatial 
information. Each track will contain its associated Media Box.*/ 
type Trak struct { 
	//extends Box
	size 		uint32 //derived from the sum of all box sizes contained in this box + 4 + 4
	boxtype		uint32
}