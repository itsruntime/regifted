package main

import (
	"regifted/driver"

	"os"
)

func main() {
	rc := driver.Main()
	if rc != 0 {
		os.Exit(rc)
	}
}
