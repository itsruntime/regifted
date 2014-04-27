// tfhd_test
package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestTfhdConstructor(t *testing.T) {
	a := []byte{1, 2, 3}
	t1 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	if t1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Tfhd.Size in the Constructor test.")
	}
	if t1.Version != 1 {
		t.Fail()
		fmt.Println("Error checking Tfhd.Version in the Constructor test.")
	}
	if t1.TrackID != 2 {
		t.Fail()
		fmt.Println("Error checking Tfhd.TrackID in the Constructor test.")
	}
	c := []byte{1, 2, 3}
	d := t1.Flags
	if len(c) != len(d) {
		t.Fail()
		fmt.Printf("Error checking Tfhd.Flags - wrong flags length, "+
			"len(Tfhd.Flags) is %d\n", len(t1.Flags))
	}
	for i := range c {
		if c[i] != d[i] {
			t.Fail()
			fmt.Printf("Slice data different in index %d.", i)
		}
	}
	if t1.BaseDataOffset != 5 {
		t.Fail()
		fmt.Println("Error checking Tfhd.BaseDataOffset in the Constructor test.")
	}
	if t1.SampleDescriptionIndex != 7 {
		t.Fail()
		fmt.Println("Error checking Tfhd.SampleDescriptionIndex in the Constructor test.")
	}
	if t1.BoxType != 1952868452 {
		t.Fail()
	}
	if t1.DefaultSampleDuration != 9 {
		t.Fail()
		fmt.Println("Error checking Tfhd.DefaultSampleDuration in the Constructor test.")
	}
	if t1.DefaultSampleSize != 11 {
		t.Fail()
		fmt.Println("Error checking Tfhd.DefaultSampleSize in the Constructor test.")
	}
	if t1.DefaultSampleFlags != 13 {
		t.Fail()
		fmt.Println("Error checking Tfhd.DefaultSampleFlags in the Constructor test.")
	}
}

func TestTfhdSetSize(t *testing.T) {
	a := []byte{1, 2, 3}
	t3 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	c := reflect.TypeOf(t3.Size).Kind()
	if c != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if t3.Size != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
	t3.SetSize(32)
	d := reflect.TypeOf(t3.Size).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if t3.Size != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestTfhdGetSize(t *testing.T) {
	a := []byte{1, 2, 3}
	t4 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	c := t4.GetSize()
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if c != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestTfhdGetBoxType(t *testing.T) {
	a := []byte{1, 2, 3}
	t5 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	c := t5.GetBoxType()
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if c != 1952868452 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestTfhdCalculateSize(t *testing.T) {
	// minimal size
	a := []byte{1, 2, 3}
	t6 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(0), uint32(0),
		uint32(0), uint32(0), uint32(0))
	t6.CalculateSize()
	c := t6.Size
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if c != 16 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size, expected 16")
	}
	// full size
	t7 := NewTfhd(uint32(16), uint8(1), a, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	t7.CalculateSize()
	g := t7.Size
	h := reflect.TypeOf(g).Kind()
	if h != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if g != 40 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size, expected 40")
	}
}

func TestTfhdWrite(t *testing.T) {
	z := []byte{1, 2, 3}
	t8 := NewTfhd(uint32(16), uint8(0), z, uint32(2), uint64(5), uint32(7),
		uint32(9), uint32(11), uint32(13))
	a := t8.Write()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Slice {
		t.Fail()
		fmt.Println(b)
		fmt.Println("Error checking write returns an Array.")
	}
	c := []byte{0, 0, 0, 16,
		116, 102, 104, 100,
		0, 1, 2, 3,
		0, 0, 0, 2,
		0, 0, 0, 0, 0, 0, 0, 5,
		0, 0, 0, 7,
		0, 0, 0, 9,
		0, 0, 0, 11,
		0, 0, 0, 13}
	if len(a) != len(c) {
		t.Fail()
		fmt.Printf("Slice lengths are different; should be identical. "+
			"len(a) = %d, len(c) = %d\n", len(a), len(c))
	}
	for i := range a {
		if a[i] != c[i] {
			t.Fail()
			fmt.Printf("Slice data different in index %d.", i)
		}
	}
}
