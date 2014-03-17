package moov

//There are four track types: video, sound, hint, and null. 
//Each type has its own media header. Each header corresponds
//to one media handler type.

//The hint media header contains general information
//for hint tracks.

//PDU = protocol data unit

type Smhd struct {
	//extends FullBox
	size		uint32
	largeSize	uint64 //size==1 if used
	boxtype		uint32
	flags		uint32 //no note on how to handle flags
	version		uint8 //seems to be set to 0 in the spec
	// gives the size in bytes of the largest PDU in this (hint) stream
	maxPDUsize  uint16
	// gives the average size of a PDU over the entire presentation
	avgPDUsize  uint16
	// gives the maximum rate in bits/second over any window of one second 
	maxbitrate	uint32
	// gives the average rate in bits/second over the entire presentation
	avgbitrate  uint32
	reserved 	uint32 //set at 0 in spec
}