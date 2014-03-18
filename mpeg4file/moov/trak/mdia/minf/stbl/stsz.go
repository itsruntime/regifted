package moov


//SampleSizeBox contains a count of the samples and their sizes.
//This is the version with set sample sizes.
type Stsz struct {
	//extends FullBox('stsz', version = 0, 0)
	size				uint32
	largeSize			uint64
	boxtype				uint32
	version				uint8
	flags				[3]byte
	//sampleSize indicates the uniform size of the samples or, if set to
	//0 it specifies that there are different sizes and they are stored
	//in the sample size table.
	sampleSize			uint32
	//sampleCount is the number of samples in the track; loop control variable
	sampleCount			uint32
	if(sampleSize == 0) {
		for i := 1; i <= sampleCount; i++ {
			//entrySize is the varying sample sizes indexed by sample number
			entrySize		uint32
		}
	}
}
