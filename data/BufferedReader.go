package data

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const BUFFER_SIZE int = 4096

type BufferedReader struct {
	buffReader *bufio.Reader
	buff       []byte
	Cursor     int64
	Size       int64
}

func NewBufferedReaderFromStream(fh *os.File) *BufferedReader {
	InitLogger()

	stat, err := fh.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()

	// buffer the stream
	buffReader := bufio.NewReader(fh)
	reader := new(BufferedReader)
	reader.buffReader = buffReader
	reader.buff = make([]byte, BUFFER_SIZE)
	reader.Size = size
	return reader
}

func (reader *BufferedReader) ReadBytes(size uint) []byte {
	logger.Trace("reader.ReadBytes(%u)", size)
	// n, err := reader.buffReader.Read(reader.buff)
	// if err != nil && err != io.EOF { panic(err) }
	// if n == 0 { fmt.Println ("n==0") }

	returnBuffer := make([]byte, size)
	n, err := reader.buffReader.Read(returnBuffer)
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == 0 {
		fmt.Println("n==0")
		return nil
	}
	if n != int(size) {
		buff := make([]byte, n, n)
		copy(buff, returnBuffer)
		returnBuffer = buff
		logger.Critical("%v", len(returnBuffer))

		diff := int(size) - n
		buff2 := make([]byte, diff)
		n, err := reader.buffReader.Read(buff2)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			fmt.Println("n==0")
			return nil
		}
		concat := append(returnBuffer, buff2...)

		logger.Critical("%v", concat)
		return concat
	}
	logger.Trace("num bytes read: %i", n)

	reader.Cursor += int64(size)

	// fmt.Println("READ:")
	// fmt.Println(returnBuffer)
	return returnBuffer
}

func (reader *BufferedReader) Read(size uint) uint {
	logger.Trace("reader.Read(%u)", size)
	// buff := make([]byte, size)
	data := reader.ReadBytes(size)
	var idx uint = uint(reader.Cursor) + size
	var idx_max uint = uint(len(data))
	var idx_ uint
	if idx <= idx_max {
		idx_ = idx
	} else {
		log.Printf("attempted to read past end of buffer\n")
		idx_ = idx_max
	}
	idx_ -= uint(reader.Cursor)

	var value uint = 0
	var i uint = 0

	for ; i < idx_; i++ {
		value |= uint(uint(data[uint64(i)]) << ((idx_ - i - 1) * 8))
	}
	// reader.Cursor += int64(idx_)
	return value
}
