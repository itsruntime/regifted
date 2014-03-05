package moof

import (
	"fmt"
	"io/ioutil"
	"os"
	"regifted/data"
)

func main() {
	args := os.Args
	fullFile, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading file")
	} else {
		reader := data.NewReader(fullFile)
		moof := new(Moof)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		moof.Read(reader)
		fmt.Println(moof)
	}

}
