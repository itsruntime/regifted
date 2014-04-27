package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Contains all the variables associated with the 
// Media Data Box
type Mdat struct {
	*BoxFields
	data []byte
}

// The Media Data Box (mdat) contains all the media data 
// (video and audio tracks) for a fragment of the presentation.
func NewMdat(s uint32, payload []byte) *Mdat {
	// Need to incorporate the uint64 handling later, unimplemented 4/13/14
	newMdat := &Mdat{&BoxFields{Size: s, BoxType: 0x6d646174}, payload}
	return newMdat
}

// DEPRECATED. 
// Sets the size variable in the Media Data Box (mdat). 
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

// Returns the size (bytes) of the Media Data Box (mdat).
func (m *Mdat) GetSize() uint32 {
	return m.Size
}

// Returns the integer identifier of the Media Data Box (mdat).
func (m *Mdat) GetBoxType() uint32 {
	return m.BoxType
}

// Returns a array of bytes representing the Media Data Box (mdat). 
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
