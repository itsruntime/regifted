package moof

// Box types
const(
)


// Start mfhd
type Mfhd struct {
	version int
	flags int
	sequence int
}
// End mfhd

// Start tfhd
type Tfhd struct{
	version int
	trackId int
	baseDataOffsetPresent bool
	sampleDescriptionPresent bool
	defaultSampleDurationPresent bool
	defaultSampleSizePresent bool
	defaultSampleFlagsPresent bool
	durationIsEmpty bool
	baseDataOffset int
	sampleDescriptionIndex int
	defaultSampleDuration int
	defaultSampleSize int
	defaultSampleFlags int
}
// End tfhd

// Start SampleInformation
type SampleInformation struct{
	duration int
	size int
	flags int
	offset int
}
// End SampleInformation

// Start trun
type Trun struct{
	version int
	dataOffsetPresent bool
	firstSampleFlagsPresent bool
    sampleDurationPresent bool
    sampleSizePresent bool
    sampleFlagsPresent bool
    sampleOffsetPresent bool
    dataOffset int
    firstSampleFlags int
    samples = []SampleInformation//I think
}
// End trun

// Start traf
type Traf struct{
	boxes []Box
}
// End traf

// Start box
type Box struct{
	// No variables 
}
// End box

// Start moof
type Moof struct{
	boxes []Box
}
// End moof

// Start mdat
type Mdat struct{
	bytes []byte //I think
}
// End mdat
