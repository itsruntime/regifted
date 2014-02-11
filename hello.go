package regifted

import (
  "fmt"
  "io"
  "os"
  // "encoding/binary"
  // "encoding/gob"
)

// Lorem Lipsum Dolar Sit Amet
func main() {
  var params = new( ProgramParameters )
  err := params.GetOpt()
  if err != nil {
    fmt.Printf( "failed to parse arguments\n%s\n", err.Error() )
    fmt.Printf( "usage: %s [input file]\n", os.Args[0] )
    os.Exit( 1 )
  }
  fmt.Println( params.input_file )

  fi, err := os.Open( os.Args[1] )
  if err != nil { panic( err ) }
  defer func() {
    if err := fi.Close() ; err != nil {
      panic( err )
    }
  }()

  buff := make( []byte, 188 )
  for {
    n, err := fi.Read( buff )
    if err != nil && err != io.EOF { panic(err) }
    if n == 0 { break }
 
    // fmt.Printf( "%s\n", buff[:n] )
    // fmt.Printf( "%q\n", buff[:n] )
    // fmt.Printf( "%x\n", buff[:n] )

    for idx, el := range buff {
      fmt.Printf( "%04v - %02x %08b", idx, el, el )
      fmt.Printf( "\n" )
    }

    fmt.Printf( "% x", buff[:n] )
    fmt.Printf( "\n\n\n" )

    break
  }
}
