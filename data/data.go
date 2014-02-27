package data

type Reader struct {
	data []byte
	cursor uint
}

func NewReader (da []byte) *Reader {
	r := new(Reader)
	r.data = da
	r.cursor = 0
	return r
}

func (r *Reader) Read(size uint) uint {
	
	var (
		value uint = 0
		i uint = 0
	)
	for ; i <size ;i++ {
		value |= (uint(r.data[r.cursor+i]) << ((size-i-1)*8))
	}
	r.cursor += size
	return value
}
