package moov

//The sample description table contains the detailed
//information on the coding type and how to initialize
//it. 


/*type Stsd struct {
	//extends Box, abstract class..hmmm...
	size					uint32
	largesize				uint64 //size==1 if used
	boxtype					uint32
	const reserved 			[6]uint8 //set to zero
	data_reference_index	uint16
}
*/
type SampleEntry struct { //takes 'format' as an arg
	//extends Box, abstract class..hmmm...
	size					uint32
	largesize				uint64 //size==1 if used
	boxtype					uint32
	const reserved 			[6]uint8 //set to zero
	data_reference_index	uint16
}

type HintSampleEntry struct {
	SampleEntry
	data			[]uint8
}