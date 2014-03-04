package data

<<<<<<< HEAD
import (
	"log"
)

var DEBUG_SIZE int = 100


type Reader struct {
	data []byte
	Cursor uint64
	Size uint64
}

=======
// Contains a file saved as byte array. 
type Reader struct {
	data []byte
	Cursor uint64
}

// Creates a new Reader for reading the data from the byte array.
>>>>>>> moof
func NewReader (da []byte) *Reader {
	r := new(Reader)
	r.data = da
	r.Cursor = 0
<<<<<<< HEAD
	r.Size = uint64(len(da))
	return r
}


func (r *Reader) Read(size uint) uint {
	if r.data == nil {
		log.Printf( "attempted to read from null buffer in data.Read()\n" )
		return 0
	}
=======
	return r
}

// Reads the number of bytes passed in as size from the data byte array in
// the Reader struct. It then concatinates the bytes and returns them as a 
// unsigned integer. 
func (r *Reader) Read(size uint) uint {
	
>>>>>>> moof
	var (
		value uint = 0
		i uint64 = 0
	)
	for ; i <uint64(size) ;i++ {
		value |= (uint(r.data[r.Cursor+i]) << ((uint64(size)-i-1)*8))
	}
	r.Cursor += uint64(size)
	return value
}

<<<<<<< HEAD
=======
// Reads the number of bytes passed in as size from the data byte array in the Reader 
// struct. It returns a byte array from the cursors current position to the cursor 
// plus the size. 
>>>>>>> moof
func (r *Reader) ReadBytes(size uint64) []byte {
	value:=r.data[r.Cursor:r.Cursor+size]
	r.Cursor += size
	return value
<<<<<<< HEAD
}


=======
}
>>>>>>> moof
