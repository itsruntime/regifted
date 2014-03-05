package main

import (
  "fmt"
)


type ElementaryStreamPacket struct {
  byteChunk []byte
  payload   []byte

  unitStart bool

  pid           uint
  hasAdaptation bool
}

//ElementaryStreamPacket Dispatch
//if unitstart, dump current PES and construct a new one,
//else append the es payload
func (elementaryStreamPacket *ElementaryStreamPacket) Dispatch() {

  var pesData Pes

  pesData = pesCollector[elementaryStreamPacket.pid]

  if elementaryStreamPacket.unitStart {

    if pesData, ok := pesCollector[elementaryStreamPacket.pid]; ok {

      pesData.pid = elementaryStreamPacket.pid
      pesData.streamType = types[elementaryStreamPacket.pid]
      pesData.Read()
      pesData.Print()

    }
    pesData = Pes{}

  }

  pesData.byteChunk = append(pesData.byteChunk, elementaryStreamPacket.payload...)

  pesCollector[elementaryStreamPacket.pid] = pesData

}


func (elementaryStreamPacket *ElementaryStreamPacket) Print() {
  fmt.Println("\n:::ES:::\n")
  fmt.Println("payload = ", elementaryStreamPacket.payload)
}

