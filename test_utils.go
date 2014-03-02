package main

import (
  "bytes"
  "encoding/hex"
  "log"
  "unicode"
)

// reads human-readable hex string (src) into a []byte (dest).
//
// The standard library provides functions to do this but they (specifically
// encoding/hex.Decode*()) does not gracefully handle input with whitespace.
// This is a wrapper around encoding/hex.Decode so that white space can be
// handled.
//
// *dest is invariably wiped
//
// Typical Usage:
//  var packetBytes []byte
//  var packetString string
//  packetString = `
//    47 02 62 17 80 70 24 F1 31 D4 33 F3 88 8B 23 5A
//    94 13 D4 1D DD CD 61 D8 73 08 82 88 B1 13 85 BB
//    D0 93 C4 41 54 BE 68 70 70 EA 3D CF 70 1F 50 CD
//    4E 26 A5 00 D2 1C C4 4C EE 71 C4 68 38 2B 91 45
//    44 40 F7 50 85 5D 7F DD 62 58 28 72 D6 4F 39 4A
//    A1 86 83 28 00 9E 2E 01 F0 AE 79 CC 16 F1 3C 19
//    7F FE A3 25 F0 1A 5B 0E CC 45 B6 1C D8 EC 2C E3
//    48 E7 78 9F F1 E1 62 30 78 1C 80 D5 D0 CF DA C0
//    6A A0 01 F4 03 FA AF 80 B2 34 DE 01 DC 14 46 72
//    0F 23 CB 70 21 DE AE 70 1D A0 F1 22 E1 10 79 1C
//    58 EA D4 4F 50 07 D1 3E D8 77 E4 63 65 2C E6 D0
//    9A 11 82 26 CC 62 D6 2E 00 1F DA C3`
//  generateBytesFromString( &packetBytes, &packetString )
//  fmt.Println( packetBytes )
//
// Params:
//  dest pointer to the destination []byte
//  src pointer to the incoming string
func generateBytesFromString(dest *[]byte, src *string) {
  // todo( mathew guest ) this should check string input size. e.g. src
  // must be a multiple of 2 or it must break down correctly
  var stringBytes []byte
  stringBytes = deFormatStringForHexDecode(src)

  // size is divided by 2 because two characters in hexadecimal (e.g. AF)
  // equate only one byte
  dest_size := len(stringBytes) / 2
  *dest = make([]byte, dest_size, dest_size)
  n_bytes_read, err := hex.Decode(*dest, stringBytes)
  if n_bytes_read != 188 {
    log.Printf( "EE: packet size not 188 bytes!\n" )
    // todo( mathew guest ) return error code and check it in ts_test
  }
  _ = n_bytes_read
  log.Printf( "\nbytes read: %d\n", n_bytes_read )
  if err != nil {
    log.Println(err)
  }
}

// removes white space from a string and also converts into byte array
//
// The two-thing function is arguably messy, but encoding/hex.Decode* is going
// to convert it anyway so the alternative becomes to put it into a string and
// then convert it right back out of a string. This function just needs more
// appropriately named.
func deFormatStringForHexDecode(s *string) []byte {
  var buff bytes.Buffer
  for _, ch := range *s {
    if !unicode.IsSpace(ch) {
      buff.WriteRune(ch)
    }
  }
  return buff.Bytes()
}
