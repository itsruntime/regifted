package driver

import (
	"regifted/ts"
	"regifted/util/mylog"

	"regifted/giftcollection"

	"flag"
	"fmt"
	"log"
	"os"
)

const LOGGER_NAME = "driver"
const LOGGER_SEVERITY_LEVEL = mylog.SEV_ERROR

func Main() int {

	logger := mylog.CreateLogger(LOGGER_NAME)
	logger.SetSeverityThresh(LOGGER_SEVERITY_LEVEL)

	filename, rv := getFilepath()
	if rv != 0 {
		os.Exit(rv)
	}
	fmt.Printf("Attempting to read file, Run 7 " + filename + "\n")

	fh, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}
	defer func() {
		if err := fh.Close(); err != nil {
			panic(err)
		}
	}()

	tsState := ts.Load(fh)

	Buffer := make([]*ts.AccessUnit, 0)

	var n int = 0

	for true {
		accsessUnit, ok := tsState.GetNextAccessUnit()

		//fmt.Printf("accsessUnit pcr", accsessUnit.Pcr)

		if ok != 0 {
			break
		}

		if n < 30 {

			//fmt.Printf("accsessUnit", accsessUnit.Pcr)

			Buffer = append(Buffer, accsessUnit)
			n++

		}else{
			giftcollection.Regift(Buffer)
			n = 0
			Buffer = make([]*ts.AccessUnit, 0)

			//EXPORT EACH ITERATION GIFTCOLLECTIONS BOXES INTO A BYTE ARRAY OR THEY WILL BE OVERWRITTEN
			//THIS BYTE ARRAY IS TO WRITE OUT TO FILE

		}
	}



	return 0

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
