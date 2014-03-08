package main

import (
	"testing"
	//"regifted/Nal"

)


func TestReadVideo(t *testing.T) {
	in := []byte{2, 5, 6, 3, 4, 5, 6, 5, 4, 3, 4, 0, 0, 0, 1, 8, 8, 4, 6, 5, 8, 3, 9, 2, 0, 4, 9, 5, 8, 0, 0, 0, 1, 9, 8, 2, 7, 4, 5, 2, 7, 4, 8, 9, 3}
	out := []byte{ 2, 5, 6, 3, 4, 5, 6, 5, 4, 3, 4, 8, 8, 4, 6, 5, 8, 3, 9, 2, 0, 4, 9, 5, 8, 9, 8, 2, 7, 4, 5, 2, 7, 4, 8, 9, 3 }

	n := new(Nal)

	if x := n.readVideo(in); x != out {
		t.Fail()
	}

}