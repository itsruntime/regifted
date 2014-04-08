package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	// "regifted/data"
	"regifted/ts"
)

func main() {
	filename, rv := getFilepath()
	if rv != 0 {
		os.Exit(rv)
	}
	fmt.Printf("Attempting to read file, Run 7 " + filename + "\n")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("did not open file\n")
		// os.Exit(66)
		// seems like panic is better?
		panic(err)
	}

	ts := ts.Load(bytes)
	_ = ts
}

// todo( mathew guest ) I think golang wants to use error as return codes but
// it's a little slow so I'm cheating
func getFilepath() (string, int) {
	flag.Parse()
	argc := flag.NArg()
	if argc < 1 {
		log.Printf("Usage: " + os.Args[0] + " [input ts file]\n")
		return "", 66
	}
	if argc > 1 {
		log.Printf("Ignoring all but first argument.\n")
		os.Exit(1)
	}
	filename := os.Args[1]
	return filename, 0
}
