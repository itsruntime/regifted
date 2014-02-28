package moof

import (
	"moof/data"
	"fmt"
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
	baseDataOffsetPresent        	uint
	sampleDescriptionPresent     	uint
	defaultSampleDurationPresent 	uint
	defaultSampleSizePresent     	uint
	defaultSampleFlagsPresent    	uint
	durationIsEmpty              	uint
	baseDataOffset               	uint64
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
	t.baseDataOffsetPresent = t.flags & 0x000001
	t.sampleDescriptionPresent = t.flags & 0x000002
	t.defaultSampleDurationPresent = t.flags & 0x000008
	t.defaultSampleSizePresent = t.flags & 0x000010
	t.defaultSampleFlagsPresent = t.flags & 0x000020
	t.durationIsEmpty = t.flags & 0x010000
	t.trackId = data.Read(BYTESINTRACKID)
	if t.baseDataOffsetPresent != 0 {
		t.baseDataOffset = uint64(data.Read(BYTESINBASEDATAOFFSET))
	} else {
		t.baseDataOffset = 0;
	}
	if t.sampleDescriptionPresent != 0{
		t.sampleDescriptionIndex = data.Read(BYTESINDESCRIPTIONINDEX)
	} else {
		t.sampleDescriptionIndex = 0;
	}
	if t.defaultSampleDurationPresent != 0{
		t.defaultSampleDuration = data.Read(BYTESSAMPLEDURATION)
	} else {
		t.defaultSampleDuration = 0;
	}
	if t.defaultSampleSizePresent != 0{
		t.defaultSampleSize = data.Read(BYTESSAMPLESIZE)
	} else {
		t.defaultSampleSize = 0
	}
	if t.defaultSampleFlagsPresent != 0{
		t.defaultSampleFlags = data.Read(BYTESSAMPLEFLAGS)
	} else {
		t.defaultSampleFlags = 0
	}
}
		
// End tfhd

// Start SampleInformation
type SampleInformation struct {
	duration uint
	size     uint
	flags    uint
	offset   uint
}

func (si *SampleInformation) Read (data *data.Reader, trun *Trun) SampleInformation{
	if trun.sampleDurationPresent != 0 {
		si.duration=data.Read(4)
	}else {
			si.duration =0 
	}
	if trun.sampleSizePresent != 0 {
		si.size = data.Read(4)
	}else{
		si.size = 0 
	}
	if trun.sampleFlagsPresent != 0 {
		si.flags = data.Read(4)
	}else {
		si.flags = 0
	}
	if trun.sampleOffsetPresent != 0 {
		si.offset = data.Read(4)
	}else{
		si.offset = 0
	}
	return *si
}



// End SampleInformation

// Start trun
type Trun struct {
	version                 uint
	dataOffsetPresent       uint
	firstSampleFlagsPresent uint
	sampleDurationPresent   uint
	sampleSizePresent       uint
	sampleFlagsPresent      uint
	sampleOffsetPresent     uint
	dataOffset              int64
	firstSampleFlags        uint
	samples                 []SampleInformation //I think

	size uint
	name uint
}

func (trun *Trun) Read (data *data.Reader){
	trun.size = data.Read(4)
	trun.name = data.Read(4)
	//Test for Error
	trun.version = data.Read(1)
	flags := data.Read(3)
	trun.dataOffsetPresent = flags & 0x000001
	trun.firstSampleFlagsPresent =             flags & 0x000004
    trun.sampleDurationPresent =               flags & 0x000100
    trun.sampleSizePresent =                   flags & 0x000200
    trun.sampleFlagsPresent =                  flags & 0x000400
    trun.sampleOffsetPresent = flags & 0x000800
    count := data.Read(4)
    if trun.dataOffsetPresent != 0 {
    	trun.dataOffset = int64(data.Read(4))
    }else {
    	trun.dataOffset = 0
    }
    if ((trun.dataOffset & 0x80000000)!=0){
    	trun.dataOffset = -0x100000000 + trun.dataOffset
    }
    if trun.firstSampleFlagsPresent !=0 {
    	trun.firstSampleFlags = data.Read(4)
    }else{
    	trun.firstSampleFlags = 0
    }
    
    for count>0 {
    	si := SampleInformation{}
    	trun.samples = append(trun.samples, si.Read(data, trun))
    }
}

// End trun

// Start traf
type Traf struct {
	boxes []Box
	size uint
	name uint
}

func (traf *Traf) Read(data *data.Reader){
	cursor:=data.Cursor
	traf.size = data.Read(4)
	traf.name = data.Read(4)
	//Test name to BOXTYPE
	for (data.Cursor-cursor)<uint64(traf.size) {
		box := Box{}
		traf.boxes = append(traf.boxes, box.Read(data))
	}

}

// End traf

// Start box
type Box struct {
	size uint
	name uint
}

func (box *Box) Read (data *data.Reader) Box{
	box.size = data.Read(4)
	box.name = data.Read(4)
	data.Cursor -= 8
	if box.name == TRAF_BOX {
		traf:=new(Traf)
		traf.Read(data)
	} else if box.name == TRUN_BOX{
		trun := new(Trun)
		trun.Read(data)
	} else if box.name == MFHD_BOX{
		mfhd := new(Mfhd)
		mfhd.Read(data)
	} else if box.name == TFHD_BOX{
		tfhd := new(Tfhd)
		tfhd.Read(data)
	}else{
		//Error error
		data.Cursor += uint64(box.size)
	}
	return *box
}

// End box

// Start moof
type Moof struct {
	boxes []Box
	size uint
	name uint
}

func (moof *Moof) Read(data *data.Reader) Moof{
	cursor:=data.Cursor
	moof.size = data.Read(4)
	moof.name = data.Read(4)
	//Test name to BOXTYPE
	for (data.Cursor-cursor)<uint64(moof.size) {
		box := Box{}
		moof.boxes = append(moof.boxes, box.Read(data))
	}
	return *moof
}

func (moof *Moof) String() string{
	return fmt.Sprintf("moof [%d] [%d]",moof.size,len(moof.boxes))
}
// End moof

// Start mdat
type Mdat struct {
	bytes []byte //I think
	size uint64
	name uint
}

func (mdat *Mdat) Read (data *data.Reader) {
	mdat.size = uint64(data.Read(4))
	mdat.name = data.Read(4)
	//Compare Name to MDAT_BOX
	if mdat.size==1 {
		mdat.size = uint64(data.Read(8))
	}
	mdat.bytes = data.ReadBytes(mdat.size)
}

// End mdat
