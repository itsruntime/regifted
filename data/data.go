package data

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unsafe"
)

// const DEBUG_SIZE int = 100

const BUFFER_SIZE int = 4096

type Reader struct {
	data   []byte
	Cursor uint64
	Size   uint64
}

type Writer struct {
	chunks []byte
	size   int
}

type BufferedReader struct {
	buffReader *bufio.Reader
	buff       []byte
	Cursor     int64
	Size       int64
}

// Creates a new Reader for reading the data from the byte array.
func NewReader(da []byte) *Reader {
	r := new(Reader)
	r.data = da
	r.Cursor = 0
	r.Size = uint64(len(da))
	return r
}

func NewReaderFromStream(fh *os.File) *Reader {
	// read the entire file at once
	stat, err := fh.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()
	bytes := make([]byte, size)
	// todo( mathew guest ) check allocation?
	n, err := fh.Read(bytes)
	_ = n
	if err != nil {
		log.Fatal(err)
	}
	return NewReader(bytes)
}

func NewBufferedReaderFromStream(fh *os.File) *BufferedReader {
	// read the entire file at once
	stat, err := fh.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()

	// buffer the stream
	buffReader := bufio.NewReader(fh)
	reader := new(BufferedReader)
	reader.buffReader = buffReader
	reader.buff = make([]byte, BUFFER_SIZE)
	reader.Size = size
	return reader
}

func (reader *BufferedReader) ReadBytes(size uint) []byte {
	// n, err := reader.buffReader.Read(reader.buff)
	// if err != nil && err != io.EOF { panic(err) }
	// if n == 0 { fmt.Println ("n==0") }

	returnBuffer := make([]byte, size)
	n, err := reader.buffReader.Read(returnBuffer)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == 0 {
		fmt.Println("n==0")
	}
	reader.Cursor += int64(size)

	fmt.Println("READ:")
	fmt.Println(returnBuffer)
	return returnBuffer
}

func (reader *BufferedReader) Read(size uint) uint {
	fmt.Printf("BufferedReader.Read(%u)", size)
	// buff := make([]byte, size)
	data := reader.ReadBytes(size)
	var idx uint = uint(reader.Cursor) + size
	var idx_max uint = uint(len(data))
	var idx_ uint
	if idx <= idx_max {
		idx_ = idx
	} else {
		log.Printf("attempted to read past end of buffer\n")
		idx_ = idx_max
	}
	idx_ -= uint(reader.Cursor)

	var value uint = 0
	var i uint = 0

	for ; i < idx_; i++ {
		value |= uint(uint(data[uint64(i)]) << ((idx_ - i - 1) * 8))
	}
	// reader.Cursor += int64(idx_)
	return value
}

// Reads the number of bytes passed in as size from the data byte array in
// the Reader struct. It then concatinates the bytes and returns them as a
// unsigned integer.
func (r *Reader) Read(size uint) uint {
	const N_MAX_BYTES = uint(unsafe.Sizeof(size)) // nothing to do with size
	// i'm just borrowing the var
	if size > N_MAX_BYTES {
		log.Printf("attempted to overflow register\n")
		return 0
	}
	if r.data == nil {
		log.Printf("attempted to read from null buffer in data.Read()\n")
		return 0
	}
	// if size < 0 {
	// 	log.Printf("attempted to read <0 bytes\n")
	// 	return 0
	// }

	var idx uint = uint(r.Cursor) + size
	var idx_max uint = uint(len(r.data))
	var idx_ uint
	if idx <= idx_max {
		idx_ = idx
	} else {
		log.Printf("attempted to read past end of buffer\n")
		idx_ = idx_max
	}
	idx_ -= uint(r.Cursor)

	var value uint = 0
	var i uint = 0

	for ; i < idx_; i++ {
		value |= uint(uint(r.data[r.Cursor+uint64(i)]) << ((idx_ - i - 1) * 8))
	}
	r.Cursor += uint64(idx_)
	return value
}

// Reads the number of bytes passed in as size from the data byte array in the Reader
// struct. It returns a byte array from the cursors current position to the cursor
// plus the size.
func (r *Reader) ReadBytes(size uint64) []byte {
	if r.data == nil {
		log.Printf("attempted to read from null buffer in data.Read()\n")
		return nil
	}
	// if size < 0 {
	// 	log.Printf("attempted to read <0 bytes\n")
	// 	return 0
	// }

	var idx uint = uint(r.Cursor) + uint(size)
	var idx_max uint = uint(len(r.data))
	var idx_ uint
	if idx <= idx_max {
		idx_ = idx
	} else {
		idx_ = idx_max
	}
	var n_bytes_read int = int(idx_ - uint(r.Cursor))

	value := r.data[r.Cursor:idx_]
	r.Cursor += uint64(n_bytes_read)
	return value
}
