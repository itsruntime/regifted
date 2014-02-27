
package moof

import (
	"moof/data"
	"testing"
	"fmt"
)

func TestRead(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	reader := data.NewReader(testData)
	m := new (Mfhd)
	m.Read(reader)
	fmt.Printf("Size = %d\n", m.size)
	if m.size != 16909060 {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", m.boxtype)
	if m.boxtype != 84281096 {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", m.version)
	if m.version != 9 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", m.flags)
	if m.flags != 658188 {
		t.Fail()
	}
	fmt.Printf("Sequence = %d\n", m.sequence)
	if m.sequence != 219025168{
		t.Fail()
	}
}