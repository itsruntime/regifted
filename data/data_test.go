package data

import (
	// "fmt"
	"bytes"
	"encoding/binary"
	// "encoding/hex"
	// "os"
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
	batman := []byte{'a', 'a', 't', 'm', 'a', 'n'}
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

func TestReaderRead(t *testing.T) {
	// arr := []byte{0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x04}
	arr := []byte{0x00, 0x00, 0x00, 0x0a}
	arr_int := binary.BigEndian.Uint32(arr)
	// binary.UVarint only reads one byte for me unless the first bytes are 0xff
	// not sure.... The other function works.

	// big endian - we are at processor level and the registers have already
	// been converted for us

	var reader *Reader
	var uint_ uint

	// one byte at a time
	reader = NewReader(arr)
	for _, v := range arr {
		x := reader.Read(1)
		if byte(x) != v {
			t.Error("Reader.Read() failed")
		}
	}

	// 4 bytes
	reader = NewReader(arr)
	uint_ = reader.Read(4)
	bytes_ := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes_, uint32(uint_))

	if !bytes.Equal(arr, bytes_) {
		t.Errorf("Read(4): got %s; want %s", bytes_, arr)
	}
	if uint32(uint_) != arr_int {
		t.Errorf("Read(4): got %d; want %d", uint_, arr_int)
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
