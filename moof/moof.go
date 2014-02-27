package moof

import (
	"moof/data"
)

const (
	BYTESINSIZE = 4
	BYTESINBOXTYPE = 4
	BYTESINVERSION = 1
	BYTESINSEQ = 4
	BYTESINFLAGS = 3
	MFHD_BOX = 0x6d666864
	TFHD_BOX = 0x74666864
	TRUN_BOX = 0x7472756e
	TRAF_BOX = 0x74726166
	MOOF_BOX = 0x6d6f6f66
	MDAT_BOX = 0x6d646174
)
// Start mfhd
type Mfhd struct {
	version  	uint
	flags    	uint
	sequence 	uint
	size 		uint
	boxtype		uint
}

func (m *Mfhd) Read (data *data.Reader){
	m.size = data.Read(BYTESINSIZE)
	m.boxtype = data.Read(BYTESINBOXTYPE)
	m.version = data.Read(BYTESINVERSION)
	m.flags = data.Read(BYTESINFLAGS)
	m.sequence = data.Read(BYTESINSEQ)
}

// End mfhd

// Start tfhd
type Tfhd struct {
	version                      int
	trackId                      int
	baseDataOffsetPresent        bool
	sampleDescriptionPresent     bool
	defaultSampleDurationPresent bool
	defaultSampleSizePresent     bool
	defaultSampleFlagsPresent    bool
	durationIsEmpty              bool
	baseDataOffset               int
	sampleDescriptionIndex       int
	defaultSampleDuration        int
	defaultSampleSize            int
	defaultSampleFlags           int
}

// End tfhd

// Start SampleInformation
type SampleInformation struct {
	duration int
	size     int
	flags    int
	offset   int
}

// End SampleInformation

// Start trun
type Trun struct {
	version                 int
	dataOffsetPresent       bool
	firstSampleFlagsPresent bool
	sampleDurationPresent   bool
	sampleSizePresent       bool
	sampleFlagsPresent      bool
	sampleOffsetPresent     bool
	dataOffset              int
	firstSampleFlags        int
	samples                 []SampleInformation //I think
}

// End trun

// Start traf
type Traf struct {
	boxes []Box
}

// End traf

// Start box
type Box struct {
	// No variables
}

// End box

// Start moof
type Moof struct {
	boxes []Box
}

// End moof

// Start mdat
type Mdat struct {
	bytes []byte //I think
}

// End mdat
