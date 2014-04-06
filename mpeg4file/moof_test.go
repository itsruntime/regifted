package mpeg4file

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestMoofConstructor(t *testing.T) {
	// moof constructed with uint32 size
	m1 := NewMoof(16, 0x6d6f6f66)
	if m1.size != 16 {
		t.Fail()
	}
	if m1.boxType != 1836019558 {
		t.Fail()
	}
	if m1.largeSize != 0 {
		t.Fail()
	}
	// moof constructed with uint64 size
	m2 := NewMoof(4294967296, 0x6d6f6f66)
	if m2.size != 1 {
		t.Fail()
	}
	if m2.boxType != 1836019558 {
		t.Fail()
	}
	if m2.largeSize != 4294967296 {
		t.Fail()
	}
}

func TestMoofSetSize(t *testing.T) {
	// moof constructed with uint32 size
	m3 := NewMoof(16, 0x6d6f6f66)
	if m3.size != 16 {
		t.Fail()
	}
	if m3.largeSize != 0 {
		t.Fail()
	}
	m3.SetSize(32)
	if m3.size != 32 {

	}
	if m3.largeSize != 0 {
		t.Fail()
	}

	// moof constructed with uint64 size
	m4 := NewMoof(4294967296, 0x6d6f6f66)
	fmt.Printf("Moof Size = %d\n", m4.size)
	if m4.size != 1 {
		t.Fail()
	}
	sizeM4 := reflect.TypeOf(m4.size).Name()
	fmt.Println(sizeM4)
	fmt.Println(strings.EqualFold(sizeM4, "uint64"))
	if strings.EqualFold(sizeM4, "uint64") {
		t.Fail()
	}
	if !(strings.EqualFold(sizeM4, "uint32")) {
		t.Fail()
	}
}

func TestMoofString(t *testing.T) {
	//Not sure what the goal is of the method
}

func TestMoofWrite(t *testing.T) {
	//Not sure of what the goal is of the method
}
