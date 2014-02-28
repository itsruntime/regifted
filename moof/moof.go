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
	BYTESINTRACKID = 4
	BYTESINBASEDATAOFFSET = 8
	BYTESINDESCRIPTIONINDEX = 4
	BYTESSAMPLEDURATION = 4
	BYTESSAMPLESIZE = 4
	BYTESSAMPLEFLAGS = 4
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
	version                      	uint
	trackId                      	uint
	baseDataOffsetPresent        	bool
	sampleDescriptionPresent     	bool
	defaultSampleDurationPresent 	bool
	defaultSampleSizePresent     	bool
	defaultSampleFlagsPresent    	bool
	durationIsEmpty              	bool
	baseDataOffset               	uint
	sampleDescriptionIndex       	uint
	defaultSampleDuration        	uint
	defaultSampleSize            	uint
	defaultSampleFlags           	uint
	size 							uint
	boxtype							uint
	flags							uint
}

func (t *Tfhd) Read (data *data.Reader){
	t.size = data.Read(BYTESINSIZE)
	t.boxtype = data.Read(BYTESINBOXTYPE)
	t.version = data.Read(BYTESINVERSION)
	t.flags = data.Read(BYTESINFLAGS)
	t.baseDataOffsetPresent = flags & 0x000001
	t.sampleDescriptionPresent = flags & 0x000002
	t.defaultSampleDurationPresent = flags & 0x000008
	t.defaultSampleSizePresent = flags & 0x000010
	t.defaultSampleFlagsPresent = flags & 0x000020
	t.durationIsEmpty = flags & 0x010000
	t.trackId = data.Read(BYTESINTRACKID)
	if t.baseDataOffsetPresent {
		t.baseDataOffset = data.Read(BYTESINBASEDATAOFFSET)
	} else {
		t.baseDataOffset = 0;
	}
	if t.sampleDescriptionPresent {
		t.sampleDescriptionIndex = data.Read(BYTESINDESCRIPTIONINDEX)
	} else {
		t.sampleDescriptionIndex = 0;
	}
	if t.defaultSampleDurationPresent {
		t.defaultSampleDuration = data.Read(BYTESSAMPLEDURATION)
	} else {
		t.defaultSampleDuration = 0;
	}
	if t.defaultSampleSizePresent {
		t.defaultSampleSize = data.Read(BYTESSAMPLESIZE)
	} else {
		t.defaultSampleSize = 0
	}
	if t.defaultSampleFlagsPresent {
		t.defaultSampleFlags = data.Read(BYTESSAMPLEFLAGS)
	} else {
		t.defaultSampleFlags = 0
	}
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
