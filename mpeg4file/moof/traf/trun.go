package traf

import "strconv"

type trun struct{
	size uint32
	largeSize uint64
	boxType uint32
	version uint8
	flags [3]byte
	sampleCount uint32
	//optional fields
	dataOffset int32
	firstSampleFlag uint32
	samples []sample
}

type sample struct{
	sampleDuration uint32
	sampleSize uint32
	sampleFlags uint32
    sampleCompositionTimeOffset uint32 //when version is 0
    sampleCompositionTimeOffsetNormal int32 // when version is 1
}

func NewTrun(s uint64){
	newTrun := new(trun)
	newTrun.SetSize(s)
	newTrun.boxType = 0x7472756E
	return newTrun
}

func (t *trun) SetSize(s uint64) {
	if s==0{
		t.size = 0
	}else if s > 4294967295 {
		t.size = 1
		t.largeSize = s
	} else {
		t.size = uint32(s)
	}
}

func (t *trun) String() string{
	return strconv.FormatUint(uint64(t.size),10)
}

func (m *trun) Write() []byte{
	buf := new(bytes.Buffer)
	var err error
	// Size
	err=binary.Write(buf, binary.BigEndian, m.size)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	// BoxType
	err = binary.Write(buf,binary.BigEndian,m.boxType)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	//version
	err = binary.Write(buf,binary.BigEndian,m.version)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	//flags
	err = binary.Write(buf,binary.BigEndian,m.flags)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	//sample count
	err = binary.Write(buf,binary.BigEndian,m.sampleCount)
	if err!=nil{
		fmt.Println("binary.Write failed:", err)
	}
	// data offset
	if m.dataOffset != 0 {
		err = binary.Write(buf,binary.BigEndian,m.dataOffset)
		if err!=nil{
		fmt.Println("binary.Write failed:", err)
		}
	}
	// first sample flags
	if m.firstSampleFlag != 0 {
		err = binary.Write(buf,binary.BigEndian,m.firstSampleFlag)
		if err!=nil{
		fmt.Println("binary.Write failed:", err)
		}
	}
	for i:=0;i<len(m.samples)-1;i++ {
		if m.sample[i].sampleDuration != 0 {
			err = binary.Write(buf,binary.BigEndian,m.sample[i].sampleDuration)
			if err!=nil{
			fmt.Println("binary.Write failed:", err)
			}
		}
		if m.sample[i].sampleSize != 0 {
			err = binary.Write(buf,binary.BigEndian,m.sample[i].sampleSize)
			if err!=nil{
			fmt.Println("binary.Write failed:", err)
			}
		}
		if m.sample[i].sampleFlags!= 0 {
			err = binary.Write(buf,binary.BigEndian,m.sample[i].sampleFlags)
			if err!=nil{
			fmt.Println("binary.Write failed:", err)
			}
		}
		if version == 0{
			if m.sample[i].sampleDuration != 0 {
				err = binary.Write(buf,binary.BigEndian,m.sample[i].sampleCompositionTimeOffset)
				if err!=nil{
				fmt.Println("binary.Write failed:", err)
				}
			}
		} else {
			f m.sample[i].sampleDuration != 0 {
				err = binary.Write(buf,binary.BigEndian,m.sample[i].sampleCompositionTimeOffsetNormal)
				if err!=nil{
				fmt.Println("binary.Write failed:", err)
				}
			}
		}
	}
	return buf.Bytes()
}




