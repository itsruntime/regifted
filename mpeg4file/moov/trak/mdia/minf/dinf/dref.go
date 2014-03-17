package moov
 
//this file contains dref, urn, and url boxes
 
 //Contains a table of DataEntryBox(es) that specify
//where the presentation data is located. There is a
//data reference index in sample descriptions that tie
//this table to samples in a track; allows splitting
//several sources.

//how flags are handled:
//only one flag is defined: 0x000001; if set then the
//media data is in the same file as the moov.

type dref struct {
	//extends FullBox
	size 			uint32
	largeSize		uint64 	//if used size = 1
	boxtype			uint32
	version			uint8 	//set to 0 in spec
	flags			[3]byte
	entry_count		uint32
	data_entry		[]DataEntryBox //not sure about this yet
}

type DataEntryBox interface {
	Location() string
}

type DataEntryUrlBox struct {
	//extends FullBox
	size			uint32
	boxtype			uint32 	//'url '
	version			uint8  	//set to 0
	flags			[3]byte	//0x000001 if sharing same file as moov
	location		string  //if relative URL then from moov location
}

type DataEntryUrnBox struct {
	//extends FullBox
	size			uint32
	boxtype			uint32 	//'url '
	version			uint8  	//set to 0
	flags			[3]byte	//0x000001 if sharing same file as moov
	location		string  
	name			string  //name of the URN
}

func (url *DataEntryUrlBox) Location string {
	return url.location
}

func (urn *DataEntryUrnBox) Location string {
	return urn.location
}