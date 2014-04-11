package mp4box

import (
	//"fmt"
	//"reflect"
	//"strings"
	"testing"
)

func TestMoofConstructor(t *testing.T) {
	m1 := NewMoof(16)
	if m1.Size != 16 {
		t.Fail()
	}
	if m1.BoxType != 1836019558 {
		t.Fail()
	}
}

func TestMoofSetSize(t *testing.T) {
	m3 := NewMoof(16)
	if m3.Size != 16 {
		t.Fail()
	}
	m3.SetSize(32)
	if m3.Size != 32 {
		t.Fail()
	}
}

func TestMoofWrite(t *testing.T) {
	//Not sure of what the goal is of the method
}
