package data

import (
	"log"
	"os"
	"unsafe"

	"fmt"
)

// const DEBUG_SIZE int = 100

type Reader struct {
	data   []byte
	Cursor uint64
	Size   uint64
}

type Writer struct {
	chunks []byte
	size   int
}

// Creates a new Reader for reading the data from the byte array.
func NewReader(da []byte) *Reader {
	InitLogger()
	r := new(Reader)
	r.data = da
	r.Cursor = 0
	r.Size = uint64(len(da))
	return r
}

func NewReaderFromStream(fh *os.File) *Reader {
	InitLogger()

	// read the entire file at once
	stat, err := fh.Stat()
	if err != nil {
		fmt.Println("reader line 37")
		log.Fatal(err)
	}
	size := stat.Size()
	bytes := make([]byte, size)
	// todo( mathew guest ) check allocation?
	n, err := fh.Read(bytes)
	_ = n
	if err != nil {
		fmt.Println("reader line 48")
		log.Fatal(err)
	}
	return NewReader(bytes)
}

// Reads the number of bytes passed in as size from the data byte array in
// the Reader struct. It then concatinates the bytes and returns them as a
// unsigned integer.
func (r *Reader) Read(size uint) uint {
	logger.Trace("reader.Read(%u)", size)
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
	//  log.Printf("attempted to read <0 bytes\n")
	//  return 0
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
	logger.Trace("reader.ReadBytes(%u)", size)
	if r.data == nil {
		log.Printf("attempted to read from null buffer in data.Read()\n")
		return nil
	}
	// if size < 0 {
	//  log.Printf("attempted to read <0 bytes\n")
	//  return 0
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
