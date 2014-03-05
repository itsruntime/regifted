package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regifted/data"
	"regifted/moof"
)

func main() {
	args := os.Args
	fullFile, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error reading file")
	} else {
		reader := data.NewReader(fullFile)
		moof := new(moof.Moof)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		moof.Read(reader)
		fmt.Println(moof)
	}

}
