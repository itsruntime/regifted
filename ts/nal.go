package ts

import "bytes"
//import "fmt"

type Nal struct {

	streamType		uint32
	units			[]byte
	
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

	//Looking for this sequence in byteArr
	sep := []byte{0, 0, 0, 1}

	//item is a [][]byte and to append to units we need byte type
	for _, item := range bytes.Split(byteArr, sep) {

		if len(item) > 0 {

			//get each byte from item and append it to units
			for i := 0; i < len(item) ; i++ {
				n.units = append(n.units, item[i])
			}//end for
			
		}//end if
	}//end for
}//end func

func (n *Nal) readAudio(bytes []byte) {

	n.units = append(n.units, bytes...)
	
}

/* FOR TESTING SAKE
  func main() {
	var units = []byte{}
	in := []byte{2, 5, 6, 3, 4, 5, 6, 5, 4, 3, 4, 0, 0, 0, 1, 8, 8, 4, 6, 5, 8, 3, 9, 2, 0, 4, 9, 5, 8, 0, 0, 0, 1, 9, 8, 2, 7, 4, 5, 2, 7, 4, 8, 9, 3}
	
	//SHOULD BE THE RESULT OF THE PRINT OUT
	//  out := []byte{ 2, 5, 6, 3, 4, 5, 6, 5, 4, 3, 4, 8, 8, 4, 6, 5, 8, 3, 9, 2, 0, 4, 9, 5, 8, 9, 8, 2, 7, 4, 5, 2, 7, 4, 8, 9, 3 }


	//Looking for this sequence in byteArr
	sep := []byte{0, 0, 0, 1}

	//item is a [][]byte and to append to units we need byte type
	for _, item := range bytes.Split(in, sep) {

		if len(item) > 0 {

			//get each byte from item and append it to units
			for i := 0; i < len(item) ; i++ {
				units = append(units, item[i])
			}//end for
			
		}//end if
	}//end for

	fmt.Println(units)
}*/