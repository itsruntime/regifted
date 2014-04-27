package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestMfhdConstructor(t *testing.T) {
	a := []byte{1, 2, 3}
	m1 := NewMfhd(uint32(16), uint8(1), a, uint32(1))
	if m1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Mfhd.Size in the Constructor test.")
	}
	if m1.Version != 1 {
		t.Fail()
		fmt.Println("Error checking Mfhd.Version in the Constructor test.")
	}
	b := []byte{1, 2, 3}
	if len(b) != len(m1.Flags) {
		t.Fail()
		fmt.Printf("Error checking Mfhd.Flags - wrong flags length, "+
			"len(Mfhd.Flags) is %d\n", len(m1.Flags))
	}
	if m1.BoxType != 1835427940 {
		t.Fail()
		fmt.Println("BoxType from Mfhd.BoxType is incorrect.")
	}
	if m1.SequenceNumber != 1 {
		t.Fail()
		fmt.Println("Error checking Mfhd.SequenceNumber in the Constructor test.")
	}
}

func TestMfhdSetSize(t *testing.T) {
	a := []byte{1, 2, 3}
	m3 := NewMfhd(uint32(16), uint8(1), a, uint32(1))
	b := reflect.TypeOf(m3.Size).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
	m3.SetSize(32)
	c := reflect.TypeOf(m3.Size).Kind()
	if c != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestMfhdGetSize(t *testing.T) {
	a := []byte{1, 2, 3}
	m4 := NewMfhd(uint32(32), uint8(1), a, uint32(1))
	b := m4.GetSize()
	c := reflect.TypeOf(b).Kind()
	if c != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if b != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestMfhdGetBoxType(t *testing.T) {
	a := []byte{1, 2, 3}
	m5 := NewMfhd(uint32(32), uint8(1), a, uint32(1))
	b := m5.GetBoxType()
	c := reflect.TypeOf(b).Kind()
	if c != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if b != 1835427940 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestMfhdCalculateSize(t *testing.T) {
	a := []byte{1, 2, 3}
	m6 := NewMfhd(uint32(8), uint8(1), a, uint32(1))
	c := m6.Size
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if m6.Size != 8 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size.")
	}
	m6.CalculateSize()
	c = m6.Size
	d = reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if c != 16 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size.")
	}

}

func TestMfhdWrite(t *testing.T) {
	a := []byte{1, 2, 3}
	m7 := NewMfhd(uint32(8), uint8(1), a, uint32(1))
	b := m7.Write()
	c := reflect.TypeOf(a).Kind()
	if c != reflect.Slice {
		t.Fail()
		fmt.Println(b)
		fmt.Println("Error checking write returns an Array.")
	}
	d := []byte{0, 0, 0, 8, 109, 102, 104, 100, 1, 1, 2, 3, 0, 0, 0, 1}
	if len(b) != len(d) {
		t.Fail()
		fmt.Println("Slice lengths are different; should be identical.")
	}
	for i := range a {
		if b[i] != d[i] {
			t.Fail()
			fmt.Printf("Slice data different in index %d.", i)
		}
	}
}
