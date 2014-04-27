package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Mdat struct {
	*BoxFields
	data []byte
}

// Need to incorporate the uint64 handling later, unimplemented 4/13/14
func NewMdat(s uint32, payload []byte) *Mdat {
	newMdat := &Mdat{&BoxFields{Size: s, BoxType: 0x6d646174}, payload}
	return newMdat
}

func (m *Mdat) SetSize(s uint64) {
	if s == 0 {
		m.Size = 0
	} else {
		if s < 4294967295 {
			m.Size = uint32(s)
		} else {
			m.Size = 1
			//m.LargeSize = s
		}
	}
}

func (m *Mdat) GetSize() uint32 {
	return m.Size
}

func (m *Mdat) GetBoxType() uint32 {
	return m.BoxType
}

func (m *Mdat) Write() []byte {
	buf := new(bytes.Buffer)
	var err error
	// Size
	err = binary.Write(buf, binary.BigEndian, m.Size)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf, binary.BigEndian, m.BoxType)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	buf.Write(m.data)
	return buf.Bytes()
}
