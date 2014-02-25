package data

import "io/ioutil"

var DEBUG_SIZE int = 100

func Read(fileName string, curser int) []byte {

	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return b
}

func ReadSegemnt(bytes []byte) uint32 {

	var segment uint32

	l := len(bytes)
	for i, b := range bytes {
		shift := uint32((l - i - 1) * 8)

		segment |= uint32(b) << shift
	}
	//fmt.Print("segment:", segment)

	return segment
}

func ReadBytes(curser int, size int, bytes []byte) []byte {

	return bytes[curser:(curser + size)]

}

func TruncateBytes(curser int, bytes []byte) []byte {

	return bytes[curser:len(bytes)]

}

func ReadHeaderData(bytes []byte) uint32 {

	var curser int = 0

	var a uint32 = (ReadSegemnt(ReadBytes(curser, 1, bytes)) >> 1) & 0x07
	curser++
	var b uint32 = ReadSegemnt(ReadBytes(curser, 1, bytes))
	curser++
	var c uint32 = (ReadSegemnt(ReadBytes(curser, 1, bytes)) >> 1) & 0x7f
	curser++
	var d uint32 = ReadSegemnt(ReadBytes(curser, 1, bytes))
	curser++
	var e uint32 = (ReadSegemnt(ReadBytes(curser, 1, bytes)) >> 1) & 0x7f
	curser++
	var timestamp uint32 = (a << 30) | (b << 22) | (c << 15) | (d << 7) | e

	return timestamp

}
