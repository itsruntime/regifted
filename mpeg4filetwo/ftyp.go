package mpeg4file

import (
	"strconv"
	"mpeg4file/box"
)

type Ftyp struct {
	*box.Box
	majorBrand uint32
	minorVersion uint32
	compatibleBrands []uint32
}

func NewFtyp(s uint64, box uint32, mBrand uint32, ver uint32, compBrands []uint32) Ftyp {
	Ftyp:= {&box.Box{0,0,box},
			mBrand,
			ver,
			compBrands}
	Ftyp.SetSize(s)
	return Ftyp
}

func (f *Ftyp) SetSize(s uint64){
	if s==0{
		f.Size=0
	} else {
		if s>4294967295 {
			f.Size = 1
			f.LargeSize = s
		}else{
			f.Size = uint32(s)
		}
	}
}

func (f *Ftyp) String() string{
	return strconv.FormatUint(uint64(f.Size),10)
}