package main

import (
	"regifted/data"
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

// A interface that defines all the methods that have to be
// implemented in order to be considered a box.
type BoxInterface interface{
	// Reads the information about a box from the reader and
	// saves it to the correct variables. It then returns
	// a pointer to the struct that was being edited.
	Read(data *data.Reader) BoxInterface
	// Calculates the size of the variables that have been
	// initialized after calling Read. This is the size of
	// variable not what it contains. The value returned is
	// the sum of all the variables in the box.
	calcSize() int
	// Returns 	a string representation of the box.
	String() string
}

// Contains all the information about
// Movie Fragment Header box (Mfhd)
type Mfhd struct {
	version  	uint
	flags    	uint
	sequence 	uint
	size 		uint
	boxtype		uint
}

// Reads the information about the Mfhd from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (m *Mfhd) Read (data *data.Reader) BoxInterface{
	m.size = data.Read(BYTESINSIZE)
	m.boxtype = data.Read(BYTESINBOXTYPE)
	m.version = data.Read(BYTESINVERSION)
	m.flags = data.Read(BYTESINFLAGS)
	m.sequence = data.Read(BYTESINSEQ)
	return m
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (m *Mfhd) calcSize() int{
	return BYTESINSIZE+BYTESINBOXTYPE+BYTESINVERSION+BYTESINFLAGS+BYTESINSEQ
}

// Returns 	a string representation of the Mfhd struct.
func (m *Mfhd) String () string {
	return fmt.Sprintf("mfhd [%d] flags=%x sequence=%d\n", m.calcSize(), m.flags, m.sequence)
}

// Contains all the information about the Track Fragment Header (Tfhd)
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

// Reads the information about the Tfhd from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (t *Tfhd) Read (data *data.Reader) BoxInterface{
	t.size = data.Read(BYTESINSIZE)
	t.boxtype = data.Read(BYTESINBOXTYPE)
	t.version = data.Read(BYTESINVERSION)
	t.flags = data.Read(BYTESINFLAGS)
	t.trackId = data.Read(BYTESINTRACKID)
	t.baseDataOffsetPresent = t.flags & 0x000001
	t.sampleDescriptionPresent = t.flags & 0x000002
	t.defaultSampleDurationPresent = t.flags & 0x000008
	t.defaultSampleSizePresent = t.flags & 0x000010
	t.defaultSampleFlagsPresent = t.flags & 0x000020
	t.durationIsEmpty = t.flags & 0x010000

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
	return t
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (t *Tfhd) calcSize() int{
	sum := BYTESINSIZE + BYTESINBOXTYPE + BYTESINVERSION + BYTESINFLAGS + BYTESINTRACKID
	if t.baseDataOffsetPresent !=0 {
		sum += BYTESINBASEDATAOFFSET
	}
	if t.sampleDescriptionPresent != 0 {
		sum += BYTESINDESCRIPTIONINDEX
	}
	if t.defaultSampleDurationPresent !=0 {
		sum += BYTESSAMPLEDURATION
	}
	if t.defaultSampleSizePresent != 0 {
		sum += BYTESSAMPLESIZE
	}
	if t.defaultSampleFlagsPresent != 0 {
		sum +=BYTESSAMPLEFLAGS
	}
	return sum
}

// Returns 	a string representation of the Tfhd struct.
func (t *Tfhd) String () string {
	return fmt.Sprintf("tfhd [%d] trackId=%d baseDataOffset=%d sampleDescriptionIndex=%d defaultSampleDuration=%d defaultSampleSize=%d defaultSampleFlags=%08x\n" ,t.calcSize(), t.trackId, t.baseDataOffset, t.sampleDescriptionIndex, t.defaultSampleDuration, t.defaultSampleSize, t.defaultSampleFlags)
}

// Contains
type SampleInformation struct {
	duration uint
	size     uint
	flags    uint
	offset   uint
}

// Reads the information about the smaple information from
// the reader and saves it to the correct variables. It
// then returns a pointer to the struct that was being edited.
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

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (si *SampleInformation) calcSize (trun *Trun) int{
	sum := 0
	if trun.sampleDurationPresent!=0 {
		sum+=4
	}
	if trun.sampleSizePresent != 0 {
		sum+=4
	}
	if trun.sampleFlagsPresent !=0 {
		sum+=4
	}
	if trun.sampleOffsetPresent !=0{
		sum+=4
	}
	return sum
}

// Returns 	a string representation of the variables in
// the sample information struct.
func (si *SampleInformation) StringSampleInforamtion() string{
	return fmt.Sprintf("s duration=%d size=%d flags=%08x offset=%08x" ,si.duration, si.size, si.flags, si.offset)
}

// Contains the information about the Track Fragment Run (Trun)
// box
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
	samples                 []SampleInformation

	size uint
	boxtype uint
	flags uint
	count uint
}

// Reads the information about the Trun from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (trun *Trun) Read (data *data.Reader) BoxInterface{
	trun.size = data.Read(4)
	trun.boxtype = data.Read(4)
	//Test for Error
	trun.version = data.Read(1)
	trun.flags = data.Read(3)
	trun.dataOffsetPresent = trun.flags & 0x000001
	trun.firstSampleFlagsPresent =             trun.flags & 0x000004
    trun.sampleDurationPresent =               trun.flags & 0x000100
    trun.sampleSizePresent =                   trun.flags & 0x000200
    trun.sampleFlagsPresent =                  trun.flags & 0x000400
    trun.sampleOffsetPresent = trun.flags & 0x000800
    trun.count = data.Read(4)
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

    for trun.count>0 {
    	si := SampleInformation{}
    	trun.samples = append(trun.samples, si.Read(data, trun))
    	trun.count--
    }
    return trun
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (trun *Trun) calcSize () int {
	sum := BYTESINSIZE + BYTESINBOXTYPE + BYTESINVERSION + BYTESINFLAGS + 4 //count
	if trun.dataOffsetPresent != 0 {
		sum +=4
	}
	if trun.firstSampleFlagsPresent != 0 {
		sum += 4
	}
	for _,element := range trun.samples{
		sum += element.calcSize(trun)
	}
	return sum
}

// Returns 	a string representation of the Trun struct.
func (trun *Trun) String () string {
	out := fmt.Sprintf("trun [%d] version=%d dataOffsetPresent=%d firstSampleFlagsPresent=%d sampleDurationPresent=%d sampleSizePresent=%d sampleFlagsPresent=%d sampleOffsetPresent=%d sampleCount=%d dataOffset=%d firstSampleFlags=%08x \n" ,trun.calcSize(), trun.version, trun.dataOffsetPresent, trun.firstSampleFlagsPresent, trun.sampleDurationPresent, trun.sampleSizePresent, trun.sampleFlagsPresent, trun.sampleOffsetPresent, len(trun.samples), trun.dataOffset, trun.firstSampleFlags)
	count :=0
	for _,element := range trun.samples{
		out = out+ fmt.Sprintf("   %d %s\n" ,count, element.StringSampleInforamtion())
		count++
	}
	return out
}

// Contains the information about the Track Fragment box (Traf)
type Traf struct {
	boxes []BoxInterface
	size uint
	boxtype uint
}

// Reads the information about the Traf from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (traf *Traf) Read(data *data.Reader) BoxInterface{
	cursor:=data.Cursor
	traf.size = data.Read(4)
	traf.boxtype = data.Read(4)
	//Test name to BOXTYPE
	for (data.Cursor-cursor)<uint64(traf.size) {
		traf.boxes = append(traf.boxes, ReadBox(data))
	}
	return traf
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (traf *Traf) calcSize () int{
	sum := BYTESINSIZE + BYTESINBOXTYPE
	for _,box := range traf.boxes{
		sum += box.calcSize()
	}
	return sum
}

// Returns 	a string representation of the Traf struct.
func (traf *Traf) String () string{
	out := fmt.Sprintf("traf [%d] [%d]\n" ,traf.calcSize(), len(traf.boxes))
	count := 0
	for _,box := range traf.boxes{
		out = out + fmt.Sprintf(" %d ",count) + box.String()
		count++
	}
	return fmt.Sprintf(out)
}

// Reads the information about a box from the reader and
// resets the cursor before calling the box founds Read
// function. It then returns a pointer to the struct
// found.
func ReadBox (data *data.Reader) BoxInterface{
	size := data.Read(4)
	boxtype := data.Read(4)
	data.Cursor -= 8
	if boxtype == TRAF_BOX {
		traf:=new(Traf)
		return traf.Read(data)
	} else if boxtype == TRUN_BOX{
		trun := new(Trun)
		return trun.Read(data)
	} else if boxtype == MFHD_BOX{
		mfhd := new(Mfhd)
		return mfhd.Read(data)
	} else if boxtype == TFHD_BOX{
		tfhd := new(Tfhd)
		return tfhd.Read(data)
	}else{
		//Error error
		data.Cursor += uint64(size)
	}
	return nil
}

// Contains the information about the Movie Fragment box (Moof)
type Moof struct {
	boxes []BoxInterface
	size uint
	boxtype uint
}

// Reads the information about the Moof from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (moof *Moof) Read(data *data.Reader) Moof{
	cursor:=data.Cursor
	moof.size = data.Read(4)
	moof.boxtype = data.Read(4)
	//Test name to BOXTYPE
	for (data.Cursor-cursor)<uint64(moof.size) {
		moof.boxes = append(moof.boxes, ReadBox(data))
	}
	return *moof
}

// Returns 	a string representation of the Moof struct.
func (moof *Moof) String() string{
	out := fmt.Sprintf("moof [%d] [%d]\n",moof.size,len(moof.boxes))
	for _, box := range moof.boxes{
		out = out + fmt.Sprintf(box.String())
	}
	return out
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (moof *Moof) calcSize () int{
	sum := BYTESINSIZE + BYTESINBOXTYPE
	for _,box:= range moof.boxes {
		sum += box.calcSize()
	}
	return sum
}

// Contains the information about the Media Data Container (Mdat)
// box.
type Mdat struct {
	bytes []byte //I think
	size uint64
	boxtype uint
}

// Reads the information about the Mdat from the reader and
// saves it to the correct variables. It then returns
// a pointer to the struct that was being edited.
func (mdat *Mdat) Read (data *data.Reader) {
	mdat.size = uint64(data.Read(4))
	mdat.boxtype = data.Read(4)
	//Compare Name to MDAT_BOX
	if mdat.size==1 {
		mdat.size = uint64(data.Read(8))
	}
	mdat.bytes = data.ReadBytes(mdat.size)
}

// Calculates the size of the variables that have been
// initialized after calling Read. This is the size of
// variable not what it contains. The value returned is
// the sum of all the variables in the box.
func (mdat *Mdat) calcSize () int{
	return 16 + len(mdat.bytes)
}

// Returns 	a string representation of the Mdat struct.
func (mdat *Mdat) String() string{
	out:= fmt.Sprintf("mdat [%d] [%d]", mdat.calcSize(), len(mdat.bytes))
	//add more to print out
	return out
}
