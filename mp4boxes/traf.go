package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"strconv"
)

type Traf struct {
	*BoxFields
}

func NewTraf(s uint32) *Traf {
	newTraf := &Traf{&BoxFields{Size: s, BoxType: 0x74726166}}
	return newTraf
}

func (t *Traf) SetSize(s uint32) {
	t.Size = s
}

func (t *Traf) GetSize() uint32 {
	return t.Size
}

func (t *Traf) GetBoxType() uint32 {
	return t.BoxType
}

func (t *Traf) CalculateSize(totalTrunSize uint32, tfhdSize uint32) {
	// need to modify to handle multiple truns
	t.Size = 8 + totalTrunSize + tfhdSize
}

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
