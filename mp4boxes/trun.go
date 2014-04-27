package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

// Contains all of the variables associated with 
// the Track Fragment Run Box.
type Trun struct {
	*FullBoxFields
	SampleCount uint32
	//optional fields
	DataOffset      int32
	FirstSampleFlag uint32
	Samples         []Sample
}

// Contains all of the variables associated with 
// sample information. 
type Sample struct {
	SampleDuration              uint32
	SampleSize                  uint32
	SampleFlags                 uint32
	SampleCompositionTimeOffset uint32 //when version is 0
	//SampleCompositionTimeOffsetNormal int32  // when version is 1
}

// The Track Fragment Run Box (trun) contatins all the sample information
// for a single Track Fragment. 
func NewTrun(s uint32, ver uint8, flag []byte, count uint32, doff int32,
	first uint32, samples []Sample) *Trun {
	newTrun := &Trun{&FullBoxFields{
		&BoxFields{s, 0x7472756E},
		ver,
		flag},
		count,
		doff,
		first,
		samples}
	return newTrun
}

// A sample is all of the data contained in one 
// time stamp of the presentation. Only a samples metadata
// is stored in the Track Fragment Run Box (trun). There are usually several
// sets of sample information in one Track Fragment Run Box (trun). 
func NewSample(d uint32, s uint32, f uint32, scto uint32) *Sample {
	newSample := &Sample{SampleDuration: d,
		SampleSize:                  s,
		SampleFlags:                 f,
		SampleCompositionTimeOffset: scto}
	return newSample
}

// DEPRECATED. 
// Sets the size variable in the Track Fragment Run Box (trun). 
func (t *Trun) SetSize(s uint32) {
	t.Size = s
}

// Returns the size of the Track Fragment Run Box (trun). 
func (t *Trun) GetSize() uint32 {
	return t.Size
}

// Returns the integer identifier of the Track Fragment Run Box (trun).
func (t *Trun) GetBoxType() uint32 {
	return t.BoxType
}

// Calcualtes the size (bytes) of the Track Fragment Run Box (trun). 
func (t *Trun) CalculateSize() {
	t.Size = 16
	if t.DataOffset != 0 {
		t.Size += 4
	}
	if t.FirstSampleFlag != 0 {
		t.Size += 4
	}
	for _, Sample := range t.Samples {
		if Sample.SampleDuration != 0 {
			t.Size += 4
		}
		if Sample.SampleSize != 0 {
			t.Size += 4
		}
		if Sample.SampleFlags != 0 {
			t.Size += 4
		}
		if Sample.SampleCompositionTimeOffset != 0 {
			t.Size += 4
		}
	}
}

// Returns a array of bytes representing the Track Fragmant Run Box (trun).
func (t *Trun) Write() []byte {
	buf := new(bytes.Buffer)
	var err error
	// Size
	err = binary.Write(buf, binary.BigEndian, t.Size)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf, binary.BigEndian, t.BoxType)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//version
	err = binary.Write(buf, binary.BigEndian, t.Version)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//flags
	err = binary.Write(buf, binary.BigEndian, t.Flags)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//sample count
	err = binary.Write(buf, binary.BigEndian, t.SampleCount)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// data offset
	if t.DataOffset != 0 {
		err = binary.Write(buf, binary.BigEndian, t.DataOffset)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	// first sample flags
	if t.FirstSampleFlag != 0 {
		err = binary.Write(buf, binary.BigEndian, t.FirstSampleFlag)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	for i := 0; i < len(t.Samples); i++ {
		if t.Samples[i].SampleDuration != 0 {
			err = binary.Write(buf, binary.BigEndian, t.Samples[i].SampleDuration)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if t.Samples[i].SampleSize != 0 {
			err = binary.Write(buf, binary.BigEndian, t.Samples[i].SampleSize)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if t.Samples[i].SampleFlags != 0 {
			err = binary.Write(buf, binary.BigEndian, t.Samples[i].SampleFlags)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if t.Version == 0 {
			if t.Samples[i].SampleDuration != 0 {
				err = binary.Write(buf, binary.BigEndian,
					t.Samples[i].SampleCompositionTimeOffset)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
			}
		} else {
			fmt.Println("Error, version = 1 in TRUN")
			//if t.Samples[i].SampleDuration != 0 {
			//	err = binary.Write(buf, binary.BigEndian,
			//		t.Samples[i].SampleCompositionTimeOffsetNormal)
			//	if err != nil {
			//		fmt.Println("binary.Write failed:", err)
			//	}
			//}
		}
	}
	return buf.Bytes()
}
