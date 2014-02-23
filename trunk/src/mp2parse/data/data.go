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

func ReadBytes(curser int, size int, bytes []byte) []byte {

	return bytes[curser:(curser+size)]

}
