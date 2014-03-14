package moov

type Mvhd struct {
	// extends FullBox
	size				uint32 //largesize uint64 if version==1
	boxtype				uint32
	version				uint8
	flags				uint //bit(24)
	//uint64 when version==1; creation time of the presentation from jan1,1904 UTC
	creation_time		uint32
	//uint64 when version==1; most recent modification saved, from jan1,1904 UTC
	modification_time	uint32 
	//the time scale for the entire presentation; number of time units to pass in a second
	timescale			uint32 
	//uint64 when version==1; the length of the entire presentation (in the timescale units)
	//duration is derived from the presentation's tracks: the value 
	//is the duration of the longest track in the presentation
	duration			uint32 
	//template, typically 1.0, a fixed point 16.16 number, referring to rate of 
	//play, preferred 1.0
	//rate int32 = 0x00010000 
	//template, typically full volume, a fixed point 8.8 number, volume, preferred 
	//at 1.0 (full)
	//volume int16 = 0x0100 
	const (
		reserved uint16 = 0    // specified as 'bit16'
		reserved [2]uint32 = 0 
	)
	//template, Unity matrix, tranformation matrix for the video
	//matrix int32 = {0x00010000,0,0,0,0x00010000,0,0,0,0x40000000} 
	pre_defined			[6]uint32
	//a nonzero integer indicating a value ot use for the next trackID added to the presentation
	//zero is not valid for next_track_ID, this value must be larger than the 
	//largest in use; if all 1's, search for unused ID
	next_track_ID		uint32 	
}