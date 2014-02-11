package regifted

import (
  "fmt"
  "os"
  "flag" // getopt-like parsing for command-line arguments
  // "encoding/binary"
  // "encoding/gob"
)

type ProgramParameters struct {
  input_file string
}

type InvalidInputError struct {
  input_to_error string
}

func (e InvalidInputError) Error() string {
  return fmt.Sprintf( "error: %s", e.input_to_error )
}

func (ProgramParameters) GetOpt() error {
  var flag_var = flag.String( "flagname", "default", "usage messsage" )
  var _ = *flag_var
  flag.Parse() 

  var argc = flag.NArg()
  if argc < 1 {
    return InvalidInputError{ "not enough arguments to do anything" }
  }

  // there is always a race condition for files.
  // todo( mathew guest ) input validation needs verification - it's going to be buggy
  // specifically: directory inputs, invalid permissions,  ...
  var input_file = flag.Arg( 0 )
  if _, err := os.Stat( input_file ); os.IsNotExist( err ) {
    return InvalidInputError{ "no such file" }
    os.Exit( 2 )
  }

  return nil
}