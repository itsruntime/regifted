package moov

//There is a different media information header for each //track type (corresponding to the media handler-type); 
//the matching header shall be present, which may be one //of those defined here, or one defined in a derived 
//specification.

//One type of mediaheader must be present
//'vmhd', 'smhd', 'hmhd', 'nmhd'

//The hint media header contains general information,
//independent of the protocol, for hint tracks. (A PDU
//is a Protocol Data Unit.) 

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