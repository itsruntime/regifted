package moov


//TimeToSampleBox ('stts') contains a compact table that indexes from decoding time
//to sample number (matching the times to the samples that should be displayed at
//the time). The table has two columns. The first column gives the number of
//consecutive samples with the same time delta. The second column gives the delta
//of the sample. The whole table provides a complete time-to-sample-map.
//This is one version of a time to sample box, and a version of the decoding time
//to sample box.
type Stts struct {
	//extends FullBox('stts', version = 0, 0)
	size				uint32
	largeSize			uint64
	boxtype				uint32
	//entryCount is the number of table entries, loop control variable
	entryCount			uint32
	for i:=0; i<entryCount; i++ {
		//sampleCount is the number of consecutive samples with the same time delta (i.e. in the
		//same duration).
		sampleCount		uint32
		//sampleDelta is the time delta between successive decoding times; represented in the
		//timescale of the media
		sampleDelta		uint32
	}
}
