package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestMdatConstructor(t *testing.T) {
	z := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	m1 := NewMdat(16, z)
	if m1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Mdat.Size in the Constructor test.")
	}
	if m1.BoxType != 1835295092 {
		t.Fail()
	}
}

func TestMdatSetSize(t *testing.T) {
	z := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	m3 := NewMdat(16, z)
	a := reflect.TypeOf(m3.Size).Kind()
	if a != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size. -1")
	}
	m3.SetSize(uint64(32))
	b := reflect.TypeOf(m3.Size).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 32 {
		t.Fail()
		fmt.Println(m3.Size)
		fmt.Println("Error checking value of Size. -2")
	}
}

func TestMdatGetSize(t *testing.T) {
	z := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	m4 := NewMdat(32, z)
	a := m4.GetSize()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if a != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestMdatGetBoxType(t *testing.T) {
	z := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	m5 := NewMdat(0, z)
	a := m5.GetBoxType()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if a != 1835295092 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestMdatWrite(t *testing.T) {
	z := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	m6 := NewMdat(8, z)
	a := m6.Write()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Slice {
		t.Fail()
		fmt.Println("Error checking write returns an Array.")
	}
	c := []byte{0, 0, 0, 8, 109, 100, 97, 116, 1, 2, 3, 4, 5, 6, 7, 8}
	if len(a) != len(c) {
		t.Fail()
		fmt.Println("Slice lengths are different; should be identical.")
	}
	for i := range a {
		if a[i] != c[i] {
			t.Fail()
			fmt.Printf("Slice data different in index %d.", i)
		}
	}
}
