package giftcollection

import (
	"fmt"
	"log"
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

func Regift(AccessUnits []*ts.AccessUnit) bool {
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

	// Dave you will need to print to file in reverse order on the
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
		log.Fatal("Not enough data do genertae pcr delta")
		return false
		
	}

	pcrDelta = (videoSamples[len(videoSamples)-1].SampleDuration) - (videoSamples[len(videoSamples)-2].SampleDuration)

	if pcrDelta == 0 {
		log.Fatal("pcrDelta is 0, cannot generate delta")
		return false
		
	}

	fmt.Println("pcrDelta", pcrDelta)

	if (videoSamples[len(videoSamples)-1].SampleDuration % uint32(pcrDelta)) == 0 {

		for i := 0; i < len(videoSamples); i++ {

			videoSamples[i].SampleDuration = pcrDelta

		}

	}

	if (audioSamples[len(audioSamples)-1].SampleDuration % uint32(pcrDelta)) == 0 {

		for i := 0; i < len(audioSamples); i++ {

			audioSamples[i].SampleDuration = pcrDelta

		}

	}

	fmt.Println("\nvideoSamples = ", videoSamples)

	fmt.Println("\naudioSamples = ", audioSamples)

	// Create mdat and add it to boxes array
	payload := append(videoByte, audioByte...)
	mdat := mp4box.NewMdat(uint32(audioSize+videoSize+8), payload)

	Boxes = append(Boxes, mdat)
	// Setting Flags for the trun should be done programatically from the
	// PES data but that can come later
	audioTrunFlags := make([]byte, 0, 3)
	audioTrunFlags = append(audioTrunFlags, 0x00)
	audioTrunFlags = append(audioTrunFlags, 0x0B)
	audioTrunFlags = append(audioTrunFlags, 0x01)
	// Add audio Samples to boxes array. Append to front of boxes array
	audioTrun := mp4box.NewTrun(
		0, //size is calculated later
		0, //version will be zero until we have a reason to do otherwise
		audioTrunFlags,
		0, //dataoffset = MOOF.SIZE + 8, must be calculated later
		0, //no reason for first-sample-flags
		uint32(len(audioSamples)),
		audioSamples)
	// Add audio trun to boxes array. Append to front of boxes array
	Boxes = append(Boxes, audioTrun)

	// Add tfhd to boxes array. Append to front of boxes array
	audioTfhdFlags := make([]byte, 0, 3)
	audioTfhdFlags = append(audioTrunFlags, 0x00)
	audioTfhdFlags = append(audioTrunFlags, 0x00)
	audioTfhdFlags = append(audioTrunFlags, 0x20)
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
	// Add audio traf to boxes array. Append to front of boxes array
	Boxes = append(Boxes, audioTfhd)

	// Add video samples to boxes array. Append to front of boxes array

	// Add video trun to boxes array. Append to front of boxes array

	// Add tfhd to boxes array. Append to front of boxes array

	// Add video traf to boxes array. Append to fron of boxes array

	// Add mfhd to boxes array. Append to front of boxes array

	// Add moof to boxes array. Append to front of boxes array

	// Call the write method for all boxes in boxes array.
	// And append the values to the end of the FileByte array.

	return false

}
