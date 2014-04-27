package giftcollection

import (
	"fmt"
	"log"
	"os"
	"regifted/mp4boxes"
	"regifted/ts"
)

type sample struct {
	size     int
	duration uint32
	flags    uint32
}

//Need a byte array to hold the boxes in a byte
//form for printing from the driver later
//type GiftCollection struct{
//    FileByte []byte
//}

// This file could be used to add the ftyp box to the beginning of the
// FileByte array. This may not be necessary based on the files provided
// by Niell
//func InitializeFileByte(){}

const (
	AUDIO_STREAM_TYPE uint = 15
	VIDEO_STREAM_TYPE uint = 27
)

var sequenceNumber uint32 = 1

// Regift takes a section of the TS stream and repackages into 
// a MPEG-4 box structure. 
func Regift(AccessUnits []*ts.AccessUnit) []byte {
	fmt.Println("\nRegift()\n\n")

	//fmt.Println("AccessUnits[0]:")
	//fmt.Println(AccessUnits[0])
	//fmt.Println("AccessUnits[0].PesMap:")
	//fmt.Println(AccessUnits[0].PesMap)

	audioByte := make([]byte, 0)
	videoByte := make([]byte, 0)
	audioSamples := make([]mp4box.Sample, 0)
	videoSamples := make([]mp4box.Sample, 0)
	trackID := 1
	// Need a array of boxes to hold the boxes
	// until they are ready to print
	// boxes = make([]mpeg4boxes, 0)
	// IMPORTANT NOTE: To have a array of the boxes they all have to
	// be in the same interface. I think this means all the box files
	// will need to be in the same package.

	// Dave you will need to print to file in reverse order from driver.go
	// Boxes is exported so that this is convenient
	Boxes := make([]mp4box.Box, 0)

	var audioSize int = 0
	var videoSize int = 0
	var pcrDelta uint32 = 0

	for _, AccessUnit := range AccessUnits {
		// fmt.Println( "for _, AccessUnit := " )
		// fmt.Println( AccessUnit )

		delta := 0

		// fmt.Println("VIDEO_STREAM_TYPE = ", AccessUnit.PesMap[VIDEO_STREAM_TYPE])

		//fmt.Println("for_, pes := AUDIO_STREAM_TYPE")
		for _, pes := range AccessUnit.PesMap[AUDIO_STREAM_TYPE] {
			//fmt.Println("audio pes payload= ", pes.Payload)
			audioByte = append(audioByte, pes.Payload...)
		}
		//fmt.Println("AFTER for_, pes := AUDIO_STREAM_TYPE")

		delta = len(audioByte) - audioSize

		audioSize = len(audioByte)

		audioSamples = append(audioSamples, mp4box.Sample{uint32(AccessUnit.Pcr), uint32(delta), 0, 0})

		// fmt.Println("audioSamples = ", audioSamples)

		for _, pes := range AccessUnit.PesMap[VIDEO_STREAM_TYPE] {
			videoByte = append(videoByte, pes.Payload...)
		}

		delta = len(videoByte) - videoSize

		videoSize = len(videoByte)

		videoSamples = append(videoSamples, mp4box.Sample{uint32(AccessUnit.Pcr), uint32(delta), 0, 0})
	}

	if len(videoSamples) < 2 {
		log.Fatal("Not enough data to genertae pcr delta")
		os.Exit(71)

	}

	pcrDelta = (videoSamples[len(videoSamples)-1].SampleDuration) - (videoSamples[len(videoSamples)-2].SampleDuration)

	if pcrDelta == 0 {
		log.Fatal("pcrDelta is 0, cannot generate delta")
		os.Exit(71)

	}

	fmt.Println("video pcrDelta", pcrDelta)

	if (videoSamples[len(videoSamples)-1].SampleDuration % uint32(pcrDelta)) == 0 {

		for i := 0; i < len(videoSamples); i++ {

			videoSamples[i].SampleDuration = pcrDelta

		}

	}

	if len(audioSamples) < 2 {
		log.Fatal("Not enough data to genertae pcr delta")
		os.Exit(71)

	}

	pcrDelta = (audioSamples[len(audioSamples)-1].SampleDuration) - (audioSamples[len(audioSamples)-2].SampleDuration)

	if pcrDelta == 0 {
		log.Fatal("pcrDelta is 0, cannot generate delta")
		os.Exit(71)

	}

	fmt.Println("audio pcrDelta", pcrDelta)

	if (audioSamples[len(audioSamples)-1].SampleDuration % uint32(pcrDelta)) == 0 {

		for i := 0; i < len(audioSamples); i++ {

			audioSamples[i].SampleDuration = pcrDelta

		}

	}

	fmt.Println("\nvideoSamples = ", videoSamples)

	fmt.Println("\naudioSamples = ", audioSamples)

	// Create mdat and add it to boxes array
	payload := append(audioByte, videoByte...)
	mdat := mp4box.NewMdat(uint32(audioSize+videoSize+8), payload)

	Boxes = append(Boxes, mdat)
	// Setting Flags for the trun should be done programatically from the
	// PES data but that can come later
	audioTrunFlags := make([]byte, 0, 3)
	audioTrunFlags = append(audioTrunFlags, 0x00)
	audioTrunFlags = append(audioTrunFlags, 0x0B)
	audioTrunFlags = append(audioTrunFlags, 0x00)
	// Add audio Samples to boxes array. Appended to rear of boxes array
	fmt.Println(uint32(len(audioSamples)))
	audioTrun := mp4box.NewTrun(
		0, //size is calculated later
		0, //version will be zero until we have a reason to do otherwise
		audioTrunFlags,
		uint32(len(audioSamples)),
		0, //dataoffset = MOOF.SIZE + 8, points to start of mdat, must be calculated later
		0, //no reason for first-sample-flags
		audioSamples)
	audioTrun.CalculateSize()
	// Add audio trun to boxes array. Appended to rear of boxes array
	Boxes = append(Boxes, audioTrun)

	// Add tfhd to boxes array. Appended to rear of boxes array
	audioTfhdFlags := make([]byte, 0, 3)
	audioTfhdFlags = append(audioTfhdFlags, 0x00)
	audioTfhdFlags = append(audioTfhdFlags, 0x00)
	audioTfhdFlags = append(audioTfhdFlags, 0x20)
	audioTfhd := mp4box.NewTfhd(
		0, //size is calculated later
		0, //version is typically 0
		audioTfhdFlags,
		uint32(trackID),
		0, //base-data-offset not obsevred in sample fragments
		0, //sample-description-index not observed in sample fragments
		0, //default-sample-duration not observed in sample fragments
		0, //default-sample-size not observed in sample fragments
		0) //default-sample-flags not observed in sample fragments
	trackID++
	audioTfhd.CalculateSize()
	Boxes = append(Boxes, audioTfhd)
	// Add audio traf to boxes array. Appended to rear of boxes array
	audioTraf := mp4box.NewTraf(0) //Size = 8 + audioTfhd.size + audioTrun.Size, calculated later
	var totalTrun uint32 = 0
	for i := len(Boxes) - 1; i > 0; i-- {
		if Boxes[i].GetBoxType() == uint32(0x74726166) {
			break
		}
		if Boxes[i].GetBoxType() == uint32(0x7472756E) {
			totalTrun += Boxes[i].GetSize()
		}
	}
	audioTraf.CalculateSize(totalTrun, audioTfhd.Size)
	Boxes = append(Boxes, audioTraf)
	// Add video samples to boxes array. Appended to rear of boxes array
	// Setting Flags for the trun should be done programatically from the
	// PES data but that can come later
	videoTrunFlags := make([]byte, 0, 3)
	videoTrunFlags = append(videoTrunFlags, 0x00)
	videoTrunFlags = append(videoTrunFlags, 0x0B)
	videoTrunFlags = append(videoTrunFlags, 0x00)
	fmt.Println("Num videoSamples is: ", len(videoSamples))
	// Add video Samples to boxes array. Appended to rear of boxes array
	videoTrun := mp4box.NewTrun(
		0, //size is calculated later
		0, //version will be zero until we have a reason to do otherwise
		videoTrunFlags,
		uint32(len(videoSamples)),
		0, //dataoffset = MOOF.SIZE + 8 + len(audioByte),
		//points to end of audio data in mdat, must be calculated later
		0, //no reason for first-sample-flags
		videoSamples)
	videoTrun.CalculateSize()
	fmt.Println("Num samplecount is: ", videoTrun.SampleCount)
	// Add video trun to boxes array.
	Boxes = append(Boxes, videoTrun)

	// Add tfhd to boxes array. Appended to rear of boxes array
	videoTfhdFlags := make([]byte, 0, 3)
	videoTfhdFlags = append(videoTfhdFlags, 0x00)
	videoTfhdFlags = append(videoTfhdFlags, 0x00)
	videoTfhdFlags = append(videoTfhdFlags, 0x20)
	videoTfhd := mp4box.NewTfhd(
		0, //size is calculated later
		0, //version is typically 0
		videoTfhdFlags,
		uint32(trackID),
		0, //base-data-offset not observed in sample fragments
		0, //sample-description-index not observed in sample fragments
		0, //default-sample-duration not observed in sample fragments
		0, //default-sample-size not observed in sample fragments
		0) //default-sample-flags not observed in sample fragments
	trackID++
	videoTfhd.CalculateSize()
	Boxes = append(Boxes, videoTfhd)
	// Add video traf to boxes array. Append to front of boxes array
	videoTraf := mp4box.NewTraf(0) //Size = 8 + audioTfhd.size + audioTrun.Size, calculated later
	totalTrun = 0
	for i := len(Boxes) - 1; i > 0; i-- {
		if Boxes[i].GetBoxType() == uint32(0x74726166) {
			break
		}
		if Boxes[i].GetBoxType() == uint32(0x7472756E) {
			totalTrun += Boxes[i].GetSize()
		}
	}
	videoTraf.CalculateSize(totalTrun, videoTfhd.Size)
	Boxes = append(Boxes, videoTraf)
	// Add mfhd to boxes array. Appended to rear of boxes array
	mfhdFlags := make([]byte, 0, 3)
	mfhdFlags = append(mfhdFlags, 0x00)
	mfhdFlags = append(mfhdFlags, 0x00)
	mfhdFlags = append(mfhdFlags, 0x00)
	mfhd := mp4box.NewMfhd(
		16, //Size = 16 always
		0,
		mfhdFlags,
		sequenceNumber)
	sequenceNumber++ //advance sequenceNumber for the next moof-mdat fragment
	Boxes = append(Boxes, mfhd)
	// Add moof to boxes array. Appended to rear of boxes array
	moof := mp4box.NewMoof(0) //Size is 8 + MFHD.SIZE + TRAFs.SIZE, calculated later
	var totalTraf uint32 = 0
	for i := len(Boxes) - 1; i > 0; i-- {
		if Boxes[i].GetBoxType() == uint32(0x74726166) {
			totalTraf += Boxes[i].GetSize()
		}
	}
	moof.CalculateSize(totalTraf, mfhd.Size)
	Boxes = append(Boxes, moof)
	

	boxesBytes := make([]byte, 0)
	for i := (len(Boxes) - 1); i >= 0; i-- {
		boxesBytes = append(boxesBytes, Boxes[i].Write()...)
	}

	return boxesBytes
}
