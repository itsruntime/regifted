package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

type Trun struct {
	*FullBoxFields
	SampleCount uint32
	//optional fields
	dataOffset      int32
	firstSampleFlag uint32
	samples         []sample
}

type Sample struct {
	SampleDuration                    uint32
	SampleSize                        uint32
	SampleFlags                       uint32
	SampleCompositionTimeOffset       uint32 //when version is 0
	SampleCompositionTimeOffsetNormal int32  // when version is 1
}

func NewTrun(s uint32, ver uint32, flag [3]byte, count uint32) *Trun {
	newTrun := &Tfhd{&FullBoxFields{Size: s,
		BoxType: 0x7472756E,
		Version: ver,
		Flags:   flag},
		SampleCount: count}
	return newTrun
}

func (t *Trun) SetSize(s uint32) {
	t.Size = s
}

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
	if m.dataOffset != 0 {
		err = binary.Write(buf, binary.BigEndian, t.dataOffset)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	// first sample flags
	if m.firstSampleFlag != 0 {
		err = binary.Write(buf, binary.BigEndian, t.firstSampleFlag)
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	for i := 0; i < len(m.samples)-1; i++ {
		if m.samples[i].sampleDuration != 0 {
			err = binary.Write(buf, binary.BigEndian, t.samples[i].sampleDuration)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if m.samples[i].sampleSize != 0 {
			err = binary.Write(buf, binary.BigEndian, m.samples[i].sampleSize)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if m.samples[i].sampleFlags != 0 {
			err = binary.Write(buf, binary.BigEndian, m.samples[i].sampleFlags)
			if err != nil {
				fmt.Println("binary.Write failed:", err)
			}
		}
		if m.version == 0 {
			if m.samples[i].sampleDuration != 0 {
				err = binary.Write(buf, binary.BigEndian, m.samples[i].sampleCompositionTimeOffset)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
			}
		} else {
			if m.samples[i].sampleDuration != 0 {
				err = binary.Write(buf, binary.BigEndian, m.samples[i].sampleCompositionTimeOffsetNormal)
				if err != nil {
					fmt.Println("binary.Write failed:", err)
				}
			}
		}
	}
	return buf.Bytes()
}
