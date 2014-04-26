package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

// Contains all of the variables associated with 
// the Movie Fragmant Box (moof). 
type Moof struct {
	*BoxFields
}

// The Movie Fragmant Box (moof) is used for holding the 
// metadata for a presentation. t contains a Movie 
// Fragment Header Box (mfhd), and then one or more Track 
// Fragment Boxes (traf).
func NewMoof(s uint32) *Moof {
	//newMoof := &moof{&box.Box{}} // creates and empty struct
	// to keep this the same, it must get the box data from Box
	// and assign it to the moof variables struct; then the handling
	// of common Box fields is outside of the moof box's control
	newMoof := &Moof{&BoxFields{Size: s, BoxType: 0x6d6f6f66}}
	return newMoof
}

// DEPRECATED. 
// Sets the size variable in the Movie Fragmant Box (moof). 
func (m *Moof) SetSize(s uint32) {
	m.Size = s
}

// Returns the size (bytes) of the Movie Fragmant Box (moof).
func (m *Moof) GetSize() uint32 {
	return m.Size
}

// Returns the integer identifier of the Movie Fragmant Box (moof).
func (m *Moof) GetBoxType() uint32 {
	return m.BoxType
}

// Calcualates the size (bytes) of the Movie Fragmant Box (moof).
func (m *Moof) CalculateSize(totalTrafSize uint32, mfhdSize uint32) {
	m.Size = 8 + totalTrafSize + mfhdSize
}

// Returns a array of bytes representing the Movie Fragmant Box (moof). 
func (m *Moof) Write() []byte {
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
	return buf.Bytes()
}
