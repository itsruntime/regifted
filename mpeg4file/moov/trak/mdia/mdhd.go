package moov

type Mdhd struct {
	//extends FullBox
	size				uint32
	largeSize			uint64	//size if size==1
	boxtype				uint32 	// 'mdhd'
	flags				uint32	//bit(24), no notes in spec on flags
	version				uint8	//either 0 or 1
	//creation time since jan 14, 1904 UTC
	creation_time		uint32  //uint64 if version == 1
	//last modified since jan 14, 1904 UTC
	modification_time	uint32  //uint64 if version == 1
	//number of time units to pass in a second; e.g. sixtieths of a second = 60
	timescale			uint32
	//duration of media in timescale units or all 1s if undetermined
	duration			uint32  //uint64 if version == 1
	pad					uint8   //bit(1) in spec
	// ISO-639-2/T language code; sets of 3; each is difference between ASCII code and 0x60
	language			[3]uint8 //int5 in spec
	pre_defined			uint16	//set to zero in the spec
}