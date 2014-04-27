package mp4box

import (
	"fmt"
	"reflect"
	//"strings"
	"testing"
)

func TestTrafConstructor(t *testing.T) {
	t1 := NewTraf(16)
	if t1.Size != 16 {
		t.Fail()
		fmt.Println("Error checking Traf.Size in the Constructor test.")
	}
	if t1.BoxType != 1953653094 {
		t.Fail()
	}
}

func TestTrafSetSize(t *testing.T) {
	t3 := NewTraf(16)
	a := reflect.TypeOf(t3.Size).Kind()
	if a != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if t3.Size != 16 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
	t3.SetSize(32)
	b := reflect.TypeOf(t3.Size).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of Size.")
	}
	if t3.Size != 32 {
		t.Fail()
		fmt.Println("Error checking value of Size.")
	}
}

func TestTrafGetSize(t *testing.T) {
	t4 := NewTraf(32)
	a := t4.GetSize()
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

func TestTrafGetBoxType(t *testing.T) {
	t5 := NewTraf(0)
	a := t5.GetBoxType()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of BoxType.")
	}
	if a != 1953653094 {
		t.Fail()
		fmt.Println("Error checking value of BoxType.")
	}
}

func TestTrafCalculateSize(t *testing.T) {
	t6 := NewTraf(16)
	a, b := 32, 64
	t6.CalculateSize(uint32(a), uint32(b))
	c := t6.Size
	d := reflect.TypeOf(c).Kind()
	if d != reflect.Uint32 {
		t.Fail()
		fmt.Println("Error checking data type of calculated Size.")
	}
	e := a + b
	if t6.Size-uint32(e) != 8 {
		t.Fail()
		fmt.Println("Error checking value of calculated Size.")
	}

}

func TestTrafWrite(t *testing.T) {
	t7 := NewTraf(8)
	a := t7.Write()
	b := reflect.TypeOf(a).Kind()
	if b != reflect.Slice {
		t.Fail()
		fmt.Println(b)
		fmt.Println("Error checking write returns an Array.")
	}
	c := []byte{0, 0, 0, 8, 116, 114, 97, 102}
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
