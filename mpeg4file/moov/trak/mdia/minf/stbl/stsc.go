package moov

//SampleToChunkBox, first, media data is subdivided into chunks and chunks can be
//subdivided into samples. Chunks and samples can be variable and this table maps
//samples to their chunks, the sample position and description.
type Stsc struct {
	//extends FullBox('stsc', version = 0, 0)
	size			uint32
	largeSize		uint64
	boxtype			uint32
	version			uint8
	flags			[3]byte
	//entryCount is the number of sample-chunk entries; loop control variable
	entryCount		uint32
	for i := 1; i <= entryCount; i++ {
		//firstChunk is the index of the first chunk in this run of chunks that
		//share the same samples-per-chunk and sample-description-index. The index
		//of the first chunk in a track is 1.
		firstChunk					uint32
		//samplesPerChunk is the number of sample in these chunks
		samplesPerChunk				uint32
		//sampleDescriptionIndex is the index of the sample entry describing the
		//samples in the chunk; it ranges from 1 to Sample Description Box total
		//entries.
		sampleDescriptionIndex		uint32
	}
}