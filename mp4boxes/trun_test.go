package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestTrunConstructor(t *testing.T) {
	a := []Sample{Sample{1, 2, 3, 4}, Sample{5, 6, 7, 8}}
	b := []byte{1, 2, 3}
	t1 := NewTrun(uint32(16), uint8(1), b, uint32(2), int32(5), uint32(7), a)
	if t1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Trun.Size in the Constructor test.")
	}
	if t1.Version != 1 {
		t.Fail()
		fmt.Println("Error checking Trun.Version in the Constructor test.")
	}
	if t1.SampleCount != 2 {
		t.Fail()
		fmt.Println("Error checking Trun.SampleCount in the Constructor test.")
	}
	c := []byte{1, 2, 3}
	d := t1.Flags
	if len(c) != len(d) {
		t.Fail()
		fmt.Printf("Error checking Trun.Flags - wrong flags length, "+
			"len(Trun.Flags) is %d\n", len(t1.Flags))
	}
	for i := range c {
		if c[i] != d[i] {
			t.Fail()
			fmt.Printf("Slice data different in index %d.", i)
		}
	}
	if t1.DataOffset != 5 {
		t.Fail()
		fmt.Println("Error checking Trun.DataOffset in the Constructor test.")
	}
	if t1.FirstSampleFlag != 7 {
		t.Fail()
		fmt.Println("Error checking Trun.FirstSampleFlag in the Constructor test.")
	}
	if t1.BoxType != 1953658222 {
		t.Fail()
	}
	e := []Sample{Sample{1, 2, 3, 4}, Sample{5, 6, 7, 8}}
	f := t1.Samples
	if len(e) != len(f) {
		t.Fail()
		fmt.Printf("Error checking Trun.Samples - wrong flags length, "+
			"len(Trun.Flags) is %d\n", len(t1.Samples))
	}
	for i := range e {
		if e[i].SampleDuration != f[i].SampleDuration {
			t.Fail()
			fmt.Printf("Slice data different in index %d. Sample.SampleDuration", i)
		}
		if e[i].SampleCompositionTimeOffset != f[i].SampleCompositionTimeOffset {
			t.Fail()
			fmt.Printf("Slice data different in index %d. Sample.CompositionTimeOffset", i)
		}
		if e[i].SampleFlags != f[i].SampleFlags {
			t.Fail()
			fmt.Printf("Slice data different in index %d. Sample.Flags", i)
		}
		if e[i].SampleSize != f[i].SampleSize {
			t.Fail()
			fmt.Printf("Slice data different in index %d. Sample.SampleSize", i)
		}
	}
}

func TestTrunSetSize(t *testing.T) {
	a := []Sample{Sample{1, 2, 3, 4}, Sample{5, 6, 7, 8}}
	b := []byte{1, 2, 3}
	t3 := NewTrun(uint32(16), uint8(1), b, uint32(2), int32(5), uint32(7), a)
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

func TestTrunGetSize(t *testing.T) {
	a := []Sample{Sample{1, 2, 3, 4}, Sample{5, 6, 7, 8}}
	b := []byte{1, 2, 3}
	t4 := NewTrun(uint32(16), uint8(1), b, uint32(2), int32(5), uint32(7), a)
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

func TestTrunGetBoxType(t *testing.T) {
	a := []Sample{Sample{1, 2, 3, 4}, Sample{5, 6, 7, 8}}
	b := []byte{1, 2, 3}
	t5 := NewTrun(uint32(16), uint8(1), b, uint32(2), int32(5), uint32(7), a)
	c := t5.GetBoxType()
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if c != 1953658222 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestTrunCalculateSize(t *testing.T) {
	// minimal size
	a := []Sample{Sample{0, 0, 0, 0}, Sample{0, 0, 0, 0}}
	b := []byte{1, 2, 3}
	t6 := NewTrun(uint32(16), uint8(1), b, uint32(2), int32(0), uint32(0), a)
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
	// no sample sizes
	e := []Sample{Sample{0, 0, 0, 0}, Sample{0, 0, 0, 0}}
	f := []byte{1, 2, 3}
	t7 := NewTrun(uint32(16), uint8(1), f, uint32(2), int32(5), uint32(7), e)
	t7.CalculateSize()
	g := t7.Size
	h := reflect.TypeOf(g).Kind()
	if h != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if g != 24 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size, expected 24")
	}
	// full size for two samples
	i := []Sample{Sample{1, 1, 1, 1}, Sample{1, 1, 1, 1}}
	j := []byte{1, 2, 3}
	t8 := NewTrun(uint32(16), uint8(1), j, uint32(2), int32(5), uint32(7), i)
	t8.CalculateSize()
	k := t8.Size
	l := reflect.TypeOf(k).Kind()
	if l != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	if k != 56 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size, expected 56")
	}
}

func TestTrunWrite(t *testing.T) {
	i := []Sample{Sample{1, 1, 1, 1}, Sample{1, 1, 1, 1}}
	j := []byte{1, 2, 3}
	t8 := NewTrun(uint32(8), uint8(0), j, uint32(2), int32(5), uint32(7), i)
	a := t8.Write()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Slice {
		t.Fail()
		fmt.Println(b)
		fmt.Println("Error checking write returns an Array.")
	}
	c := []byte{0, 0, 0, 8,
		116, 114, 117, 110,
		0, 1, 2, 3,
		0, 0, 0, 2,
		0, 0, 0, 5,
		0, 0, 0, 7,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1,
		0, 0, 0, 1}
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
