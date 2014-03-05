package data

import (
	//"fmt"
	"bytes"
	"testing"
)

func TestReader(t *testing.T) {
	// todo( mathew guest ) I'm unsure if go has a way to do death tests. Slices
	// still allow you to runtime panic if you access out of bounds. We need a
	// way to recover from failure aka death tests. I don't see anything obvious
	// in the go testing library but maybe with panic or something unknown.
	var s []byte
	baseline := []byte{'b', 'a', 't', 'm', 'a', 'n'}

	// null slice - the best I can currently do is that they don't segfault
	s = make([]byte, 6, 6)
	copy(s, baseline)
	reader := NewReader(s)
	_ = reader
	reader.Read(100)
	reader.ReadBytes(100)

	s = make([]byte, 6, 6)
	copy(s, baseline)
	reader = NewReader(s)
	reader.ReadBytes(100)
	reader.Read(100)

	// read past end
	s = make([]byte, 6, 6)
	copy(s, baseline)
	reader = NewReader(s)
	reader.Read(100000000)

	// read after cursor is too far
	_ = reader.Read(0)
	_ = reader.Read(1)
	_ = reader.Read(10)

	s = []byte{'b', 'a', 't', 'm', 'a', 'n'}
	reader = NewReader(s)
	reader.Read(0)
	reader.Read(0)

	s = make([]byte, 5, 5)
	reader = NewReader(s)
	reader.Read(5)
}

func TestReaderConstructor(t *testing.T) {
	empty := []byte{}
	batman := []byte{'b', 'a', 't', 'm', 'a', 'n'}
	var reader *Reader

	reader = NewReader(empty)
	if !bytes.Equal(empty, reader.data) {
		t.Error("constructor failed")
	}

	reader = NewReader(batman)
	if !bytes.Equal(batman, reader.data) {
		t.Error("constructor failed")
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
