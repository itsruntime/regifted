package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

// Contains all of the variables associated with 
// the Track Fragment Header Box (tfhd). 
type Tfhd struct {
	*FullBoxFields
	TrackID uint32
	//optional fields
	BaseDataOffset         uint64
	SampleDescriptionIndex uint32
	DefaultSampleDuration  uint32
	DefaultSampleSize      uint32
	DefaultSampleFlags     uint32
}

// The Track Fragment Header Box (tfhd) contains all the metadata
// about which elements to use when parsing the sample information.
func NewTfhd(s uint32, ver uint8, flag []byte, trackID uint32,
	bdoff uint64, sdind uint32, defSD uint32, defSS uint32,
	defSF uint32) *Tfhd {
	newTfhd := &Tfhd{&FullBoxFields{
		&BoxFields{s, 0x74666864},
		ver,
		flag},
		trackID,
		bdoff,
		sdind,
		defSD,
		defSS,
		defSF}
	return newTfhd
}

// DEPRECATED. 
// Sets the size variable in the Track Fragment Header Box(tfhd) 
func (t *Tfhd) SetSize(s uint32) {
	t.Size = s
}

// Returns the size variable of the Track Fragment Header Box (tfhd)
func (t *Tfhd) GetSize() uint32 {
	return t.Size
}

// Returns the integer identifier of the Track Fragement Header Box (tfhd)
func (t *Tfhd) GetBoxType() uint32 {
	return t.BoxType
}

// Calculates the size (bytes) of the Track Fragment Header Box (traf). 
// The size of this box varies depending on what flags are set during
// construction of the box. 
func (t *Tfhd) CalculateSize() {
	t.Size = 16
	if t.BaseDataOffset != 0 {
		t.Size += 8
	}
	if t.SampleDescriptionIndex != 0 {
		t.Size += 4
	}
	if t.DefaultSampleDuration != 0 {
		t.Size += 4
	}
	if t.DefaultSampleSize != 0 {
		t.Size += 4
	}
	if t.DefaultSampleFlags != 0 {
		t.Size += 4
	}
}

// Returns a array of bytes representing the Track Fragment Header Box (tfhd). 
func (t *Tfhd) Write() []byte {
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
	//trackID
	err = binary.Write(buf, binary.BigEndian, t.TrackID)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	if t.BaseDataOffset != 0 {
		err = binary.Write(buf, binary.BigEndian, t.BaseDataOffset)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.SampleDescriptionIndex != 0 {
		err = binary.Write(buf, binary.BigEndian, t.SampleDescriptionIndex)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.DefaultSampleDuration != 0 {
		err = binary.Write(buf, binary.BigEndian, t.DefaultSampleDuration)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.DefaultSampleSize != 0 {
		err = binary.Write(buf, binary.BigEndian, t.DefaultSampleSize)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.DefaultSampleFlags != 0 {
		err = binary.Write(buf, binary.BigEndian, t.DefaultSampleFlags)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	return buf.Bytes()
}
