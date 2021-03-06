package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestMoofConstructor(t *testing.T) {
	m1 := NewMoof(16)
	if m1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Moof.Size in the Constructor test.")
	}
	if m1.BoxType != 1836019558 {
		t.Fail()
	}
}

func TestMoofSetSize(t *testing.T) {
	m3 := NewMoof(16)
	a := reflect.TypeOf(m3.Size).Kind()
	if a != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
	m3.SetSize(32)
	b := reflect.TypeOf(m3.Size).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if m3.Size != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestMoofGetSize(t *testing.T) {
	m4 := NewMoof(32)
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

func TestMoofGetBoxType(t *testing.T) {
	m5 := NewMoof(0)
	a := m5.GetBoxType()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if a != 1836019558 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestMoofCalculateSize(t *testing.T) {
	m6 := NewMoof(16)
	a, b := 32, 64
	m6.CalculateSize(uint32(a), uint32(b))
	c := m6.Size
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	e := a + b
	if m6.Size-uint32(e) != 8 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size.")
	}

}

func TestMoofWrite(t *testing.T) {
	m7 := NewMoof(8)
	a := m7.Write()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Slice {
		t.Fail()
		fmt.Println(b)
		fmt.Println("Error checking write returns an Array.")
	}
	c := []byte{0, 0, 0, 8, 109, 111, 111, 102}
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
