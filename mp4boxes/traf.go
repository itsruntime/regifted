package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

// Contains all of the variables associated with 
// the Track Fragment Box (traf). 
type Traf struct {
	*BoxFields
}

// Within the Movie Fragment Box there is a set of track fragments, 
// these tracks are sub tracks of the Movie Fragment.The track 
// fragments contains zero or more track runs, 
// each of which document a contiguous run of samples for that track.
func NewTraf(s uint32) *Traf {
	newTraf := &Traf{&BoxFields{Size: s, BoxType: 0x74726166}}
	return newTraf
}

// DEPRECATED. 
// Sets the size variable in the Track Fragment Box (traf)
func (t *Traf) SetSize(s uint32) {
	t.Size = s
}

// Returns the size of the Track Fragment Box (traf) and 
// any boxes contained within the Track Fragment Box (traf). 
func (t *Traf) GetSize() uint32 {
	return t.Size
}

// Returns the integer identifier of the Track Fragment Box (traf). 
func (t *Traf) GetBoxType() uint32 {
	return t.BoxType
}

// Calculates the size (bytes) of the Track Fragment Box (traf). 
func (t *Traf) CalculateSize(totalTrunSize uint32, tfhdSize uint32) {
	// need to modify to handle multiple truns
	t.Size = 8 + totalTrunSize + tfhdSize
}

// Returns a array of bytes representing the Track Fragment Box (traf). 
func (t *Traf) Write() []byte {
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
	return buf.Bytes()
}
