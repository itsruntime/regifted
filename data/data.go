/*
In order to recreate data, it seems like it is necessary to create
two structs: one for Reader and one for Writer. The entire program
should run on functions since everything here is like a helper method.
*/
package data

// struct for Reader
type Reader struct {
	data byte[]
	cursor int
}

// struct for Writer
type Writer struct {
	chunks byte[] // unclear from the python
	size int
}

// def read(self, size)

// def readBytes(self, size)

// def readString(self)

// def readAll(self)

// def hasBytes(self)

// def __str__(self)

// def write(self, size, value

// def writeBytes(self, bytes)

// def writeString(self, s)

// def getBytes(self)

// def getSize(self)

// def __str__(self)