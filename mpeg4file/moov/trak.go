package moov

/*
Tracks are used for two purposes: (a) to contain media data (media tracks) 
and (b) to contain packetization information for streaming protocols 
(hint tracks). 
*/ 
type Trak struct { 
	//extends Box
	//derived from the sum of all box sizes contained in this box + 4 + 4
	size 		uint32 
	boxtype		uint32
}

func NewTrak (s uint32, box uint32) *Trak {
	newTrak := new(Trak)
	newTrak.size = s
	newTrak.boxtype = "trak"
	return newTrak
}

func (trak *Trak)SetSize (s uint32) {
	trak.size = s
}