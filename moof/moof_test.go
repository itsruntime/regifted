
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
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// sampledescriptionindex
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleduration
		0x00, 0x00, 0x00, 0x00,
		// defaultsamplesize
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleflags
		0x00, 0x00, 0x00, 0x00}
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
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// sampledescriptionindex
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleduration
		0x00, 0x00, 0x00, 0x00,
		// defaultsamplesize
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleflags
		0x00, 0x00, 0x00, 0x00}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 658188 {
		t.Fail()
	}
	fmt.Printf("baseDataOffsetPresent = %t\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent != 0 && x.baseDataOffset !=0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %t\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent != 0 && x.sampleDescriptionIndex != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %t\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent != 0 && x.defaultSampleDuration!=0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %t\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent != 0 && x.defaultSampleSize!=0{
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %t\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent != 0 && x.defaultSampleFlags!=0 {
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
		0x01, 0x00, 0x3b,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// sampledescriptionindex
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleduration
		0x00, 0x00, 0x00, 0x00,
		// defaultsamplesize
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleflags
		0x00, 0x00, 0x00, 0x00}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)	
	fmt.Printf("baseDataOffsetPresent = %t\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent == 0 && x.baseDataOffset ==0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %t\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent == 0 && x.sampleDescriptionIndex == 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %t\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent == 0 && x.defaultSampleDuration==0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %t\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent == 0 && x.defaultSampleSize==0{
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %t\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent == 0 && x.defaultSampleFlags==0 {
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
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// sampledescriptionindex
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleduration
		0x00, 0x00, 0x00, 0x00,
		// defaultsamplesize
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleflags
		0x00, 0x00, 0x00, 0x00}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("baseDataOffset = %d\n", x.baseDataOffset)
	if x.baseDataOffset != 16909060 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionIndex = %d\n", x.sampleDescriptionIndex)
	if x.sampleDescriptionIndex != 84281096 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDuration = %d\n", x.defaultSampleDuration)
	if x.defaultSampleDuration != 9 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSize = %d\n", x.defaultSampleSize)
	if x.defaultSampleSize != 9 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlags = %d\n", x.defaultSampleFlags)
	if x.defaultSampleFlags != 9 {
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
		0x01, 0x00, 0x3b,
		// trackid
		0x00, 0x00, 0x00, 0x01,				  
		// basedataoffset 
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		// sampledescriptionindex
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleduration
		0x00, 0x00, 0x00, 0x00,
		// defaultsamplesize
		0x00, 0x00, 0x00, 0x00,
		// defaultsampleflags
		0x00, 0x00, 0x00, 0x00}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("baseDataOffset = %d\n", x.baseDataOffset)
	if x.baseDataOffset != 16909060 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionIndex = %d\n", x.sampleDescriptionIndex)
	if x.sampleDescriptionIndex != 84281096 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDuration = %d\n", x.defaultSampleDuration)
	if x.defaultSampleDuration != 9 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSize = %d\n", x.defaultSampleSize)
	if x.defaultSampleSize != 9 {
		t.Fail()
	}
	fmt.Printf("defaultSampleFlags = %d\n", x.defaultSampleFlags)
	if x.defaultSampleFlags != 9 {
		t.Fail()
	}
}


//Tests for SampleInformation

//Tests for Trun

//Tests for Traf

//Tests for Box

//Tests for Moof 

//Tests for Mdat

