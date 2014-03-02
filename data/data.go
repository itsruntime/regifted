package data


var DEBUG_SIZE int = 100


type Reader struct {
	data []byte
	Cursor uint64
	Size uint64
}

func NewReader (da []byte) *Reader {
	r := new(Reader)
	r.data = da
	r.Cursor = 0
	r.Size = uint64(len(da))
	return r
}


func (r *Reader) Read(size uint) uint {

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

func (r *Reader) ReadBytes(size uint64) []byte {
	value:=r.data[r.Cursor:r.Cursor+size]
	r.Cursor += size
	return value
}


