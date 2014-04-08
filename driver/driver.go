package main

import (
	"flag"
	"fmt"
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

	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
		// os.Exit(66)
		panic(err)
	}
	ts := ts.Load(file)
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
