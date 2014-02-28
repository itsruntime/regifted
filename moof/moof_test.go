
package moof

import (
	"moof/data"
	"testing"
	"fmt"
)

func TestRead(t *testing.T) {
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

//Test that flags are set to true
func TestReadTfhdFalse(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("Size = %d\n", x.size)
	if x.size != 16909060 {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", x.boxtype)
	if x.boxtype != 84281096 {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", x.version)
	if x.version != 9 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 658188 {
		t.Fail()
	}
	fmt.Printf("baseDataOffsetPresent = %t\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent != true && x.baseDataOffset !=0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %t\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent != true && x.sampleDescription != 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %t\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent != true && x.defaultSampleDuration!=0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %t\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent != true && x.defaultSampleSize!=0{
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %t\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent != true && x.defaultSampleFlag!=0 {
		t.Fail()
	}
}

//Test that flags are set to false
func TestReadTfhdTrue(t *testing.T) {
	testData := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i:=1;i<=40;i++{
		testData.append(i-1)
	}
	reader := data.NewReader(testData)
	x := new (Tfhd)
	x.Read(reader)
	fmt.Printf("Size = %d\n", x.size)
	if x.size != 16909060 {
		t.Fail()
	}
	fmt.Printf("Type = %d\n", x.boxtype)
	if x.boxtype != 84281096 {
		t.Fail()
	}
	fmt.Printf("Version = %d\n", x.version)
	if x.version != 9 {
		t.Fail()
	}
	fmt.Printf("Flags = %d\n", x.flags)
	if x.flags != 658188 {
		t.Fail()
	}
	fmt.Printf("baseDataOffsetPresent = %t\n", x.baseDataOffsetPresent)
	if x.baseDataOffsetPresent == true && x.baseDataOffset ==0 {
		t.Fail()
	}
	fmt.Printf("sampleDescriptionPresent = %t\n", x.sampleDescriptionPresent)
	if x.sampleDescriptionPresent == true && x.sampleDescription == 0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleDurationPresent = %t\n", x.defaultSampleDurationPresent)
	if x.defaultSampleDurationPresent == true && x.defaultSampleDuration==0 {
		t.Fail()
	}
	fmt.Printf("defaultSampleSizePresent = %t\n", x.defaultSampleSizePresent)
	if x.defaultSampleSizePresent == true && x.defaultSampleSize==0{
		t.Fail()
	}
	fmt.Printf("defaultSampleFlagsPresent = %t\n", x.defaultSampleFlagsPresent)
	if x.defaultSampleFlagsPresent == true && x.defaultSampleFlag==0 {
		t.Fail()
	}
}