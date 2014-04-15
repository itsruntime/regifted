//Package box provides the two highest level boxes in
//the nested hierarchy of the 14496-12 ISO base media file
//format spec. These two boxes are Box and FullBox. FullBox
//inherits from Box, but because of the ubiquity of
//both boxtypes as parent classes in the spec, they are
//both presented here.
package mp4box

type Box interface {
	Write() []byte
	GetSize() uint32    //Sorry Go
	GetBoxType() uint32 //Sorry Again, Go
}

//Box variables are all exported;
//easier to work with from outside the package.
type BoxFields struct {
	//Size is always the sum of the sizes of the Box subtype
	//variables plus the size of all child boxes plus the size of
	//the data it or any of its children holds.
	Size uint32
	//LargeSize is rarely used except with mdats and probably
	//not necessary when using Fragments; if used, Size == 1.
	//LargeSize uint64
	//Boxtype is the 4-char type of the Box.
	BoxType uint32
}

//NewBox is a Box factory that can be used to
//include a Box as a field in a struct in another
//package. It is not required to use this since
//the exported variables are sufficient for outside
//access.
func NewBoxFields(size uint32, boxtype uint32) BoxFields {
	return BoxFields{size, boxtype}
}

//FullBox variables are all exported;
//easier to work with from outside the package.
type FullBoxFields struct {
	*BoxFields
	//Version specifies the box version; typically either 0 or 1.
	Version uint8
	//Flags are utilized to indicate a variety of present/absent
	//conditions for a particular FullBox subtype.
	Flags []byte
}

//NewFullBox is a FullBox factory that can be used to
//include a FullBox as a field in a struct in another
//package. It is not required to use this since
//the exported variables are sufficient for outside
//access.
func NewFullBoxFields(size uint32, boxtype uint32,
	version uint8, flags []byte) FullBoxFields {
	return FullBoxFields{&BoxFields{size, boxtype},
		version,
		flags}
}
