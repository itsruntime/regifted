package moov

//SyncSampleBox identifies the sync samples within the stream. The entries
//must be sorted increasing with sample number. When this box is not present
//then every sample is a sync sample.
type Stss struct {
	//extends FullBox('stss', version = 0, 0)
	
	//entryCount is the number of sync samples; loop control variable
	entryCount			uint32
	for i := 0; i < entryCount; i++ {
		//sampleNumber is the numbers of the sync samples
		sampleNumber		uint32
	}
}
