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
//BOX interface
type BoxInterface interface{
	Read(data *data.Reader) BoxInterface
	calcSize() int
	String() string
}

// Start mfhd
type Mfhd struct {
	version  	uint
	flags    	uint
	sequence 	uint
	size 		uint
	boxtype		uint
}

func (m *Mfhd) Read (data *data.Reader) BoxInterface{
	m.size = data.Read(BYTESINSIZE)
	m.boxtype = data.Read(BYTESINBOXTYPE)
	m.version = data.Read(BYTESINVERSION)
	m.flags = data.Read(BYTESINFLAGS)
	m.sequence = data.Read(BYTESINSEQ)
	return m
}

func (m *Mfhd) calcSize() int{
	return BYTESINSIZE+BYTESINBOXTYPE+BYTESINVERSION+BYTESINFLAGS+BYTESINSEQ
}

func (m *Mfhd) String () string {
	return fmt.Sprintf("mfhd [%d] flags=%x sequence=%d", m.calcSize(), m.flags, m.sequence)
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

func (t *Tfhd) String () string {
	return fmt.Sprintf("tfhd [%d] trackId=%d baseDataOffset=%d sampleDescriptionIndex=%d defaultSampleDuration=%d defaultSampleSize=%d defaultSampleFlags=%08x" ,t.calcSize(), t.trackId, t.baseDataOffset, t.sampleDescriptionIndex, t.defaultSampleDuration, t.defaultSampleSize, t.defaultSampleFlags)
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

func (si *SampleInformation) StringSampleInforamtion() string{
	return fmt.Sprintf("s duration=%d size=%d flags=%08x offset=%08x" ,si.duration, si.size, si.flags, si.offset)
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
	samples                 []SampleInformation 

	size uint
	name uint
}

func (trun *Trun) Read (data *data.Reader) BoxInterface{
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
    return trun
}

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

func (trun *Trun) String () string {
	out := fmt.Sprintf("trun [%d] version=%d dataOffsetPresent=%d firstSampleFlagsPresent=%d sampleDurationPresent=%d sampleSizePresent=%d sampleFlagsPresent=%d sampleOffsetPresent=%d sampleCount=%d dataOffset=%d firstSampleFlags=%08x" ,trun.calcSize(), trun.version, trun.dataOffsetPresent, trun.firstSampleFlagsPresent, trun.sampleDurationPresent, trun.sampleSizePresent, trun.sampleFlagsPresent, trun.sampleOffsetPresent, len(trun.samples), trun.dataOffset, trun.firstSampleFlags)
	count :=0
	for _,element := range trun.samples{
		out = out+ fmt.Sprintf("  %d %s" ,count, element.StringSampleInforamtion())
		count++
	} 
	return out
}
// End trun

// Start traf
type Traf struct {
	boxes []BoxInterface
	size uint
	name uint
}

func (traf *Traf) Read(data *data.Reader) BoxInterface{
	cursor:=data.Cursor
	traf.size = data.Read(4)
	traf.name = data.Read(4)
	//Test name to BOXTYPE
	for (data.Cursor-cursor)<uint64(traf.size) {
		box := new(Box)
		traf.boxes = append(traf.boxes, box.ReadBox(data))
	}
	return traf
}

func (traf *Traf) calcSize () int{
	sum := BYTESINSIZE + BYTESINBOXTYPE
	for _,box := range traf.boxes{
		sum += box.calcSize()
	}
	return sum
}

func (traf *Traf) String () string{
	out := fmt.Sprintf("traf [%d] [%d]" ,traf.calcSize(), len(traf.boxes))
	count := 0
	for _,box := range traf.boxes{
		out = out + box.String() + fmt.Sprintf(" %d ",count)
		count++
	}
	return "\n"+fmt.Sprintf(out)
}

// End traf

// Start box
type Box struct {
	size uint
	name uint
}

func (box *Box) ReadBox (data *data.Reader) BoxInterface{
	box.size = data.Read(4)
	box.name = data.Read(4)
	data.Cursor -= 8
	if box.name == TRAF_BOX {
		traf:=new(Traf)
		return traf.Read(data)
	} else if box.name == TRUN_BOX{
		trun := new(Trun)
		return trun.Read(data)
	} else if box.name == MFHD_BOX{
		mfhd := new(Mfhd)
		return mfhd.Read(data)
	} else if box.name == TFHD_BOX{
		tfhd := new(Tfhd)
		return tfhd.Read(data)
	}else{
		//Error error
		data.Cursor += uint64(box.size)
	}
	return nil
}

//func (box *Box) String(addition string) string{
//	return fmt.Sprintf(addition + "box %d %d", box.name, box.size)
//}

// End box

// Start moof
type Moof struct {
	boxes []BoxInterface
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
		moof.boxes = append(moof.boxes, box.ReadBox(data))
	}
	return *moof
}

func (moof *Moof) String() string{
	return fmt.Sprintf("moof [%d] [%d]",moof.size,len(moof.boxes))
}

func (moof *Moof) calcSize () int{
	sum := BYTESINSIZE + BYTESINBOXTYPE
	for _,box:= range moof.boxes {
		sum += box.calcSize()
	}
	return sum
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

func (mdat *Mdat) calcSize () int{
	return 16 + len(mdat.bytes)
}

func (mdat *Mdat) String() string{
	out:= fmt.Sprintf("mdat [%d] [%d]", mdat.calcSize(), len(mdat.bytes)) 
	//add more to print out
	return out
}

// End mdat
