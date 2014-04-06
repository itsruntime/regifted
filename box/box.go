//Package box provides the two highest level boxes in
//the nested hierarchy of the 14496-12 ISO base media file
//format spec. These two boxes are Box and FullBox. FullBox
//inherits from Box, but because of the ubiquity of
//both boxtypes as parent classes in the spec, they are
//both presented here.
package box

//Box variables are all exported;
//easier to work with from outside the package.
type Box struct {
	//Size is always the sum of the sizes of the Box subtype
	//variables plus the size of all child boxes plus the size of
	//the data it or any of its children holds.
	Size uint32
	//LargeSize is rarely used except with mdats and probably
	//not necessary when using Fragments; if used, Size == 1.
	LargeSize uint64
	//Boxtype is the 4-char type of the Box.
	Boxtype uint32
}

//NewBox is a Box factory that can be used to
//include a Box as a field in a struct in another
//package. It is not required to use this since
//the exported variables are sufficient for outside
//access.
func NewBox(size uint32, largeSize uint64, boxtype uint32) Box {
	return Box{size, largeSize, boxtype}
}

//FullBox variables are all exported;
//easier to work with from outside the package.
type FullBox struct {
	Box
	//Version specifies the box version; typically either 0 or 1.
	Version uint8
	//Flags are utilized to indicate a variety of present/absent
	//conditions for a particular FullBox subtype.
	Flags [3]uint8
}

//NewFullBox is a FullBox factory that can be used to
//include a FullBox as a field in a struct in another
//package. It is not required to use this since
//the exported variables are sufficient for outside
//access.
func NewFullBox(size uint32, largeSize uint64, boxtype uint32,
	version uint8, flags [3]uint8) FullBox {
	return FullBox{Box: Box{size, largeSize, boxtype},
		Version: version,
		Flags:   flags}
}

//Converts a string supplied as a boxtype to a uint32 value
//with the most significant 8 bits being the first letter of
//the boxtype.
func BoxTypeFromString(box string) boxType uint32 {
	return 0
}
