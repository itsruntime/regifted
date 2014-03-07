package main

//import "strings"
//import "bytes"
//import bufio

type Nal struct {
	streamType uint32
	units      []byte
}

func (n *Nal) read(bytes []byte) {

	//var H264_VIDEO uint32 = 0x1b

	if n.streamType == 0x1b {
		n.readVideo(bytes)
	} else {
		n.readAudio(bytes)
	}

}

func (n *Nal) readVideo(byteArr []byte) {

	var i int = 0
	//var ba []byte

	//items := byteArr //bytes.Split(byteArr, "0\0\0\1", 0)	//*** ??? ***

	for ; i < len(byteArr); i++ {
		value := byteArr[i]

		ba := []byte{value}

		n.units = append(n.units, ba...)

	}

}

func (n *Nal) readAudio(bytes []byte) {

	n.units = append(n.units, bytes...)

}
