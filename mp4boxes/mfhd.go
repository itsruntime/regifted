package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

// Contains all of the variables associated with 
// the Movie Fragmant Header Box (moof). 
type Mfhd struct {
	*FullBoxFields
	SequenceNumber uint32
}

// The Movie Fragmant Header Box (mfhd) is used for soring the
// sequence number (as a safty check) of the movie fragment. The
// movie fragments must be in the correct order so the the MPEG-4
// file can be correctly parsed. 
func NewMfhd(s uint32, ver uint8, flag []byte, sqn uint32) *Mfhd {
	newMfhd := &Mfhd{&FullBoxFields{
		&BoxFields{s, 0x6d666864},
		ver,
		flag},
		sqn}
	return newMfhd
}

// DEPRECATED. 
// Sets the size variable in the Movie Fragmant Header Box (mfhd). 
func (m *Mfhd) SetSize(s uint32) {
	m.Size = s
}

// Returns the Size of the Movie Fragment Header Box (mfhd).
func (m *Mfhd) GetSize() uint32 {
	return m.Size
}

// Returns the integer identifier of the Movie Fragmant 
// Header Box (mfhd) box.
func (m *Mfhd) GetBoxType() uint32 {
	return m.BoxType
}

// Calculates the size of the Movie Fragment Header Box (mfhd).
// Currently the size of the Movie Fragment Header Box is a constant
// 16 bytes. 
func (m *Mfhd) CalculateSize() {
	m.Size = 16 // size = 4, boxtype = 4, version = 1, flags = 3, Sequence # = 4
}

// Returns a array of bytes representing the Movie Fragmant Header Box (mfhd) box. 
func (m *Mfhd) Write() []byte {
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
	//version
	err = binary.Write(buf, binary.BigEndian, m.Version)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//flags
	err = binary.Write(buf, binary.BigEndian, m.Flags)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	//sequence number
	err = binary.Write(buf, binary.BigEndian, m.SequenceNumber)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
