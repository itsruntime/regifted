package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

type Tfhd struct {
	*FullBoxFields
	TrackID uint32
	//optional fields
	baseDataOffset         uint64
	sampleDescriptionIndex uint32
	defaultSampleDuration  uint32
	defaultSampleSize      uint32
	defaultSampleFlags     uint32
}

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

func (t *Tfhd) SetSize(s uint32) {
	t.Size = s
}

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
	if t.baseDataOffset != 0 {
		err = binary.Write(buf, binary.BigEndian, t.baseDataOffset)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.sampleDescriptionIndex != 0 {
		err = binary.Write(buf, binary.BigEndian, t.sampleDescriptionIndex)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.defaultSampleDuration != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleDuration)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.defaultSampleSize != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleSize)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if t.defaultSampleFlags != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleFlags)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	return buf.Bytes()
}
