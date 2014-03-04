/*
Test of the function and methods in data.go
*/
package data

import (
	"testing"
	"fmt"
)

func TestRead(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	reader := NewReader(testData)
	one := reader.Read(1)
	if one != 1 {
		t.Fail()
	}
	two := reader.Read(2)
	fmt.Printf("Two equals %d", two)
	if two != 515 {
		t.Fail()
	}
}
// struct for Reader


// struct for Writer


// def read(self, size)

// def readBytes(self, size)

// def readString(self)

// def readAll(self)

// def hasBytes(self)

// def __str__(self)

// def write(self, size, value

// def writeBytes(self, bytes)

// def writeString(self, s)

// def getBytes(self)

// def getSize(self)

// def __str__(self)
