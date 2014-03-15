package moov

type Tkhd struct {
	//extends FullBox
	size				uint32
	boxtype				uint32
	version				uint8 //version = 0 or 1
	//flag#1 Track_enabled 0x000001. Interpret as true = track present
	//flag#2 Track_in_movie 0x000002. True = track used in presentation
	//flag#3 Track_in_preview 0x000004. True = track used to preview presentation
	//do we actually need a flags variable or just the individual flags for constructing an mp4?
	flags				uint  //bit(24)
	//creation time of the track; jan 1, 1904 utc
	creation_time		uint32 //uint64 when version = 1
	//last modification time of the track; jan 1 1904 utc
	modification_time	uint32 //uint54 when version = 1
	track_ID			uint32
	const reserved uint32 = 0
	//duration of the track in the timescale of mvhd
	//equals sum of the duration of all track's edits OR
	//sum of the sample durations converted to timescale of mvhd
	//all 1s if undetermined
	duration			uint32 //uint64 when version = 1
	const reserved [2]uint32 = 0
	const reserved uint16 = 0
	//layer int16 = 0 //template, front to back ordering of video tracks (not sure what this means)
	//alternative_group int16 = 0//a value that can refer to a collection of tracks
	//volume int16 = {if track_is_audio 0x0100 else 0 } //fixed 8.8 value specifying track's relative audio volume, 1.0 is normal
	//matrix []uint32 = {0x00010000,0,0,0,0x00010000,0,0,0,0x40000000} //unity matrix - tranformation matrix for the video
	//track's width as fixed 16.16 value; pixel dimensions of images are the default
	width 				uint32
	//track's height as fixed 16.16 value; pixel dimensions of images are the default
	height				uint32
}
	