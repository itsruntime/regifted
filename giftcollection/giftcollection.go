package giftcollection

import (
    "regifted/ts"
    //"regifted/mpeg4file"
    //"regifted/mpeg4file/moof"
    //"regifted/mpeg4file/moof/traf"
)

type sample struct{

    size int
    duration uint32
    flags uint32
}

const (
AUDIO_STREAM_TYPE = 15
VIDEO_STREAM_TYPE = 27

)
func Regift(tsArray []*ts.TSState) bool{

        audioByte := make([]byte, 0)
        videoByte := make([]byte, 0)
        audioSamples := make([]sample, 0)
        videoSamples := make([]sample, 0)

    var audioSize int = 0
    var videoSize int = 0

    for _, ts := range(tsArray){

        delta := 0

        for _, pes := range(ts.PesMap[AUDIO_STREAM_TYPE]){
            audioByte = append(audioByte, pes.Payload...)
        }

        delta = len(audioByte) - audioSize

        audioSize = len(audioByte)

        audioSamples = append(audioSamples, sample{delta, 0 , 0})

        for _, pes := range(ts.PesMap[VIDEO_STREAM_TYPE]){
            videoByte = append(videoByte, pes.Payload...)
        }


        delta = len(videoByte) - videoSize

        videoSize = len(videoByte)

        videoSamples = append(videoSamples, sample{delta, 0 , 0})



    }



     return false

}
