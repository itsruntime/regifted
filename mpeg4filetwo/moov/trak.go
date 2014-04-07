package moov

import (
	"mpeg4file/box"
)

//Trak boxes contain a single track of the presentation.
//Traks have two functions:
//1) contain the metadata and data of a track in a presentation
//2) contain packetization information for streaming protocols
type Trak struct { 
	*box.Box
}

//Trak constructor.
func NewTrak (s uint32, box uint32) Trak {
	Trak := {&box.Box{s,0,box}}
	return Trak
}

//SetSize is required to calculate the size from the Trak
//sub-boxes, backtrack, and set the final size of the box.
func (trak *Trak)SetSize (s uint32) {
	trak.size = s
}