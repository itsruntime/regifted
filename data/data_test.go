package data

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
	// "encoding/hex"
	// "io/ioutil"
	"os"
)

const BYTE_FILE = "/tmp/genned_bytes"

func TestBufferedReader(t *testing.T) {
	filename := BYTE_FILE
	fh, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
	defer func() {
		if err := fh.Close(); err != nil {
			panic(err)
		}
	}()
	reader := NewBufferedReaderFromStream(fh)
	_ = reader
}

func TestReader(t *testing.T) {
	fmt.Println("")
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
	// been converted for us... I think...

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

func TestReaderReadBytes(t *testing.T) {
	batman := []byte{'a', 'a', 't', 'm', 'a', 'n'}
	var expect []byte
	var bytes_ []byte
	var reader *Reader

	reader = NewReader(batman)
	bytes_ = reader.ReadBytes(0)
	// empty read
	expect = batman[0:0]
	if !bytes.Equal(bytes_, expect) {
		t.Errorf("ReadBytes(0): got %s; want %s", bytes_, expect)
	}

	// one byte at a time
	n := len(batman)
	for idx := 0; idx < n; idx++ {
		bytes_ = reader.ReadBytes(1)
		expect = batman[idx : idx+1]
		if !bytes.Equal(bytes_, expect) {
			t.Errorf("ReadBytes(1): got %s; want %s", bytes_, expect)
		}
	}

	// read past end
	bytes_ = reader.ReadBytes(1)
	expect = []byte{}
	if !bytes.Equal(bytes_, expect) {
		t.Errorf("ReadBytes(1): got %s; want %s", bytes_, expect)
	}

	// read full slice at once
	reader = NewReader(batman)
	bytes_ = reader.ReadBytes(uint64(len(batman)))
	expect = batman
	if !bytes.Equal(bytes_, expect) {
		t.Errorf("ReadBytes(6): got %s; want %s", bytes_, expect)
	}

	// read more than full slice at once
	reader = NewReader(batman)
	bytes_ = reader.ReadBytes(uint64(len(batman)) + 1)
	expect = batman
	if !bytes.Equal(bytes_, expect) {
		t.Errorf("ReadBytes(6): got %s; want %s", bytes_, expect)
	}
}
