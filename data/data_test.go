package data

import (
  // "fmt"
  // "io/ioutil"
  // "log"
  // "os"
  "testing"
)

func TestData(t *testing.T) {
  var s []byte
  reader := NewReader(s)
  reader.Read(100)
  _ = reader
}
