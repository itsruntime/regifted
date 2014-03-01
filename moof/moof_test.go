
package moof

import (
	"moof/data"
	"testing"
	"fmt"
)

//Tests setting of Mfhd struct fields by Mfhd.Reader(reader Reader)

func TestReadMfhdBoxFields(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	reader := data.NewReader(testData)
	m := new (Mfhd)
	m.Read(reader)
	fmt.Printf("Size = %d\n", m.size)
	if m.size != 16909060 {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", m.boxtype)
	if m.boxtype != 84281096 {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", m.version)
	if m.version != 9 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", m.flags)
	if m.flags != 658188 {
		t.Fail()
	}
	fmt.Printf("Sequence = %d\n", m.sequence)
	if m.sequence != 219025168{
		t.Fail()
	}
}

// Test setting of fullbox variables in Tfhd by Tfhd.Reader(reader Reader)

func TestReadTfhdBoxFields(t *testing.T) {
	testData := []byte{
		// size
		0x00, 0x11, 0x22, 0x33, 
		// boxtype
		0x74, 0x66, 0x68, 0x64, 
		// version
		0x01, 
		// flags
		0x01, 0x00, 0x3b,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		// sampledescriptionindex
		0x11, 0x11, 0x11, 0x11,
		// defaultsampleduration
		0x22, 0x22, 0x22, 0x22,
		// defaultsamplesize
		0x33, 0x33, 0x33, 0x33,
		// defaultsampleflags
		0x44, 0x44, 0x44, 0x44}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("Size = %d\n", x.size)
	if x.size != 1122867 {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", x.boxtype)
	if x.boxtype != 1952868452 {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", x.version)
	if x.version != 1 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 65595 {
		t.Fail()
	}
	fmt.Printf("TrackId = %d\n", x.trackId)
	if x.trackId != 1 {
		t.Fail()
	}	
}


//Simple test that flags are set to TRUE in Tfhd.Read(reader Reader)

func TestReadTfhdFlagsTrue(t *testing.T) {
	testData := []byte{
		// size
		0x00, 0x11, 0x22, 0x33, 
		// boxtype
		0x74, 0x66, 0x68, 0x64, 
		// version
		0x01, 
		// flags
		0x01, 0x00, 0x3b,
		// trackid
		0x00, 0x00, 0x00, 0x00,				  
		// basedataoffset 
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		// sampledescriptionindex
		0x11, 0x11, 0x11, 0x11,
		// defaultsampleduration
		0x22, 0x22, 0x22, 0x22,
		// defaultsamplesize
		0x33, 0x33, 0x33, 0x33,
		// defaultsampleflags
		0x44, 0x44, 0x44, 0x44}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 65595 {
		t.Fail()
	}
	fmt.Printf("baseDataOffsetPresent = %d\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent != 1 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %d\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent != 2 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %d\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent != 8 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %d\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent != 16 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %d\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent != 32 {
		t.Fail()
	}
	fmt.Printf("durationIsEmpty = %d\n", x.durationIsEmpty)
	if x.durationIsEmpty != 65536 {
		t.Fail()
	}
}

//Simple test that flags are set to FALSE in Tfhd.Read(reader Reader)

func TestReadTfhdFlagsFalse(t *testing.T) {
	testData := []byte{
		// size
		0x00, 0x11, 0x22, 0x33, 
		// boxtype
		0x74, 0x66, 0x68, 0x64, 
		// version
		0x01, 
		// flags
		0x00, 0x00, 0x00,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		// sampledescriptionindex
		0x11, 0x11, 0x11, 0x11,
		// defaultsampleduration
		0x22, 0x22, 0x22, 0x22,
		// defaultsamplesize
		0x33, 0x33, 0x33, 0x33,
		// defaultsampleflags
		0x44, 0x44, 0x44, 0x44}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)	
	fmt.Printf("baseDataOffsetPresent = %d\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent != 0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %d\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %d\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %d\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %d\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent != 0 {
		t.Fail()
	}
}

//Simple test of setting flagged fields in Tfhd.Read(reader Reader) when all flags
//TRUE

func TestReadTfhdFlagFieldsSetTrue(t *testing.T) {
	testData := []byte{
		// size
		0x00, 0x11, 0x22, 0x33, 
		// boxtype
		0x74, 0x66, 0x68, 0x64, 
		// version
		0x01, 
		// flags
		0x01, 0x00, 0x3b,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		// sampledescriptionindex
		0x11, 0x11, 0x11, 0x11,
		// defaultsampleduration
		0x22, 0x22, 0x22, 0x22,
		// defaultsamplesize
		0x33, 0x33, 0x33, 0x33,
		// defaultsampleflags
		0x44, 0x44, 0x44, 0x44}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("baseDataOffset = %d\n", x.baseDataOffset)
	if x.baseDataOffset != 72340172838076673 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionIndex = %d\n", x.sampleDescriptionIndex)
	if x.sampleDescriptionIndex != 286331153 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDuration = %d\n", x.defaultSampleDuration)
	if x.defaultSampleDuration != 572662306 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSize = %d\n", x.defaultSampleSize)
	if x.defaultSampleSize != 858993459 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlags = %d\n", x.defaultSampleFlags)
	if x.defaultSampleFlags != 1145324612 {
		t.Fail()
	}
}

//Simple test of setting flagged fields in Tfhd.Read(reader Reader) when all flags
//FALSE

func TestReadTfhdFlagFieldsSetFalse(t *testing.T) {
	testData := []byte{
		// size
		0x00, 0x11, 0x22, 0x33, 
		// boxtype
		0x74, 0x66, 0x68, 0x64, 
		// version
		0x01, 
		// flags
		0x00, 0x00, 0x00,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		// sampledescriptionindex
		0x11, 0x11, 0x11, 0x11,
		// defaultsampleduration
		0x22, 0x22, 0x22, 0x22,
		// defaultsamplesize
		0x33, 0x33, 0x33, 0x33,
		// defaultsampleflags
		0x44, 0x44, 0x44, 0x44}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("baseDataOffset = %d\n", x.baseDataOffset)
	if x.baseDataOffset != 0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionIndex = %d\n", x.sampleDescriptionIndex)
	if x.sampleDescriptionIndex != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDuration = %d\n", x.defaultSampleDuration)
	if x.defaultSampleDuration != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSize = %d\n", x.defaultSampleSize)
	if x.defaultSampleSize != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlags = %d\n", x.defaultSampleFlags)
	if x.defaultSampleFlags != 0 {
		t.Fail()
	}
}


//Tests for SampleInformation when the flags are false
func TestSampleInformationFlagsFalse(t *testing.T) {
	testData := []byte{
		// duration
		0x00, 0x11, 0x22, 0x33,
		// size
		0x00, 0x22, 0x33, 0x44,
		// flags
		0x00, 0x55, 0x66, 0x77,
		// offset
		0x00, 0x88, 0x99, 0x11}
	reader := data.NewReader(testData)
	y := new(Trun)
	y.sampleDurationPresent = 0
	y.sampleSizePresent = 0
	y.sampleFlagsPresent = 0
	y.sampleOffsetPresent = 0
	x := new(SampleInformation)
	
	x.Read(reader, y)
	fmt.Printf("Duration = %d\n", x.duration)
	if x.duration != 0 {
		t.Fail()
	}
	fmt.Printf("Size = %d\n", x.size)
	if x.size != 0 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 0 {
		t.Fail()
	}
	fmt.Printf("Offset = %d\n", x.offset)
	if x.offset != 0 {
		t.Fail()
	}
}

//Tests for SampleInformation when the flags are true
func TestSampleInformationFlagsTrue(t *testing.T) {
	testData := []byte {
		// duration
		0x00, 0x11, 0x22, 0x33,
		// size
		0x00, 0x22, 0x33, 0x44,
		// flags
		0x00, 0x55, 0x66, 0x77,
		// offset
		0x00, 0x88, 0x99, 0x11}
	reader := data.NewReader(testData)
	x := new(SampleInformation)
	y := new(Trun)
	y.sampleDurationPresent = 1
	y.sampleSizePresent = 1
	y.sampleFlagsPresent = 1
	y.sampleOffsetPresent = 1
	x.Read(reader, y)
	fmt.Printf("Duration = %d\n", x.duration)
	if x.duration != 1122867 {
		t.Fail()
	}
	fmt.Printf("Size = %d\n", x.size)
	if x.size != 2241348 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 5596791 {
		t.Fail()
	}
	fmt.Printf("Offset = %d\n", x.offset)
	if x.offset != 8952081 {
		t.Fail()
	}
}

//Tests for Trun.Read in assigning values to the
// fields inherited from Box and Fullbox
func TestReadTrunBoxFields(t *testing.T) {
	testData := []byte{}
	reader := data.NewReader(testData)
	x := new(Trun)
	x.Read(reader)
	fmt.Printf("Size = %d\n", x.size)
	if x.size !=  {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", x.boxtype)
	if x.boxtype !=  {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", x.version)
	if x.version !=  {
		t.Fail()
	}
}

//Simple test that flags are set to TRUE in Trun.Read(reader Reader)

//Simple test that flags are set to FALSE in Trun.Read(reader Reader)

//Simple test of setting flagged fields in Trun.Read(reader Reader) 
//when all flags TRUE

//Simple test of setting flagged fields in Trun.Read(reader Reader) 
//when all flags FALSE

//Tests for Traf.Read in assigning values to the
//fields inherited from Box and the []Box <-- NEED GOOD TEST!

//Tests for Box.Read in assigning values to the Box fields and
//returning a child of Box of the appropriate subtype. To test
//whether this is happening, make tests for each possible returned
//boxtype (traf, mfhd, tfhd, trun) and MAKE SURE TO CREATE
//[]BYTE THAT HAS ALL NECESSARY INFORMATION TO DO BOTH BOX.READ
//AND THE SUBSEQUENT SUBTYPE_BOX.READ 

//Tests for Moof 

//Tests for Mdat

