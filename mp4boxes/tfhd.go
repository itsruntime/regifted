package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
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

func NewTfhd(s uint32, ver uint32, flag [3]byte, trackID uint32) *Tfhd {
	newTfhd := &Tfhd{&FullBoxFields{Size: s,
		BoxType: 0x74666864,
		Version: ver,
		Flags:   flag},
		TrackID: trackID}
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
	if m.baseDataOffset != 0 {
		err = binary.Write(buf, binary.BigEndian, t.baseDataOffset)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if m.sampleDescriptionIndex != 0 {
		err = binary.Write(buf, binary.BigEndian, t.sampleDescriptionIndex)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if m.defaultSampleDuration != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleDuration)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if m.defaultSampleSize != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleSize)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	if m.defaultSampleFlags != 0 {
		err = binary.Write(buf, binary.BigEndian, t.defaultSampleFlags)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	return buf.Bytes()
}
