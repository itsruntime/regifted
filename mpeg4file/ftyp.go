package mpeg4file

import "strconv"

type Level1 interface {
	String() string
}

type Ftyp struct {
	size uint32
	largeSize uint64
	boxType uint32
	majorBrand uint32
	minorVersion uint32
	compatibleBrands []uint32
}

func NewFtyp(s uint64, box uint32, mBrand uint32, ver uint32, compBrands []uint32){
	newFtyp:=new(Ftyp)
	newFtyp.SetSize(s)
	newFtyp.boxType = box
	newFtyp.majorBrand = mBrand
	newFtyp.minorVersion = ver
	newFtyp.compatibleBrands = compBrands
}

func (f *Ftyp) SetSize(s uint64){
	if s==0{
		f.size=0
	} else {
		if s>4294967295 {
			f.size = 1
			f.largeSize = s
		}else{
			f.size = uint32(s)
		}
	}
}

func (f *Ftyp) String() string{
	return strconv.FormatUint(uint64(f.size),10)
}