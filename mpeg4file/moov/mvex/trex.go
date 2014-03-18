package moov


//TrackExtendsBox sets up the default values in Movie Fragments(moof).
//The default flags are not made available through the flags variable
//of FullBox. That field is all 0. The defaultSampleFlags are utilized
//by trun and tfhd and are set here according to the following
//bits and their associated flags in the 32-bit defaultSampleFlags variable:
//	reserved					bit4	<-- defined in independent and disposable samples box
// 	isLeading 					uint2	<-- defined in independent and disposable samples box
//	sampleDependsOn 			uint2	<-- defined in independent and disposable samples box
//	sampleIsDependedOn			uint2	<-- defined in independent and disposable samples box
//	sampleHasRedundancy			uint2	<-- defined in independent and disposable samples box
//	samplePaddingValue			bit3	<-- defined in padding bits table
//	sampleIsNonSyncSample		bit1 	<-- provides same information as found in stss
//	sampleDegradationPriority	uint16	<-- defined as for degradation priority table

type Trex struct {
	//extends FullBox('trex',0,0)
	size				uint32
	largeSize			uint64
	boxtype				uint32
	version				uint8
	flags				[3]byte
	//trackID matches a track's trackID found in the Movie Box (moov)
	trackID				uint32
	//defaultSampleDescriptionIndex is the default sample description 
	//index used in the track fragments
	defaultSampleDescriptionIndex		uint32
	//defaultSampleDuration is the default sample duration used in the
	//track fragments
	defaultSampleDuration				uint32
	//defaultSampleSize is the default sample size used in the track
	//fragments
	defaultSampleSize					uint32
	//defaultSampleFlags are the default sample flags used in the track
	//fragments
	defaultSampleFlags					uint32
}