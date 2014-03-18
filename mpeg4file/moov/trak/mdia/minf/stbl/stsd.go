package moov

//The sample description table contains the detailed
//information on the coding type and how to initialize
//it.
type SampleEntry struct { //takes 'format' as an arg
	//extends Box, abstract class..hmmm...
	size					uint32
	largesize				uint64 //size==1 if used
	boxtype					uint32
	const reserved 			[6]uint8 //set to zero
	//dataReferenceIndex is an index where the samples that use this sample description
	//can retrieve data associated with the sample (stored in Data Reference Boxes)
	dataReferenceIndex		uint16
}

//HintSampleEntry takes 'protocol' as an argument
type HintSampleEntry struct {
	SampleEntry
	data			[]uint8
}

//BitRateBox
type BitRateBox struct {
	//extends Box
	size				uint32
	largeSize			uint64
	boxtype				uint32
	//bifferSizeDB is the size of the decoding buffer for the elementary stream in bytes
	bufferSizeDB		uint32
	//maxBitRate is the maximum rate in bits/second over any one second window
	maxBitrate			uint32
	//avgBitRate is the average rate in bits/second over the presentation
	avgBitrate			uint32
}

//MetaDataSampleEntry takes 'codingname' as an arg
type MetaDataSampleEntry struct {
	SampleEntry
}

//XMLMetaDataSampleEntry passes 'metx' to MetaDataSampleEntry.
type XMLMetaDataSampleEntry struct {
	MetaDataSampleEntry
	//contentEncoding is a null-terminted string, UTF-8, expresses
	//a MIME type which identifies the content encoding of the timed
	//metadata.
	contentEncoding			string 	//optional
	//namespace is the namespace of the XML schema; identifies the type of
	//metadata and for decoding XML
	namespace				string
	//schemaLocation can provide a URL for the schema
	schemaLocation			string 	//optional
	BitRateBox						//optional
}

//TextMetaDataSampleEntry passes 'mett' to MetaDataSampleEntry
type TextMetaDataSampleEntry struct {
	MetaDataSampleEntry
	//contentEncoding is a null-terminted string, UTF-8, expresses
	//a MIME type which identifies the content encoding of the timed
	//metadata.
	contentEncoding			string	//optional
	//mimeFormat is the MIME type of the timed metadata
	mimeFormat				string	
	BitRateBox						//optional
}

//URIBox
type URIBox struct {
	//extends FullBox
	size			uint32
	largeSize		uint64
	boxtype			uint32
	version			uint8 		//version set to 0
	flags			[3]uint8
	//theURI is the URI formatted according to 6.2.4
	theURI			string
}

//URIInitBox
type URIInitBox struct {
	//extends FullBox
	size			uint32
	largeSize		uint64
	boxtype			uint32
	version			uint8 		//version set to 0
	flags			[3]uint8
	//URIInitializationData is opaque data defined by URI docs
	uriInitializationData		[]uint8
}

//URIMetaSampleEntry
type URIMetaSampleEntry struct {
	the_label		URIBox
	init			URIInitBox	//optional
	MPEG4BitRateBox				//optional
}

//Visual Sequences section______________________________________

//PixelAspectRatioBox ('pasp')
type PixelAspectRatioBox struct {
	//extends Box
	size			uint32
	largeSize		uint64
	boxtype			uint32
	//hSpacing is the relative width of a pixel
	hSpacing		uint32
	//vSpacing is the relative height of a pixel
	vSpacing		uint32
}

//CleanApertureBox ('clap')
type CleanApertureBox struct {
	//extends Box
	size					uint32
	largeSize				uint64
	boxtype					uint32
	//cleanApertureWidthN is the first part of a fraction that describes
	//the exact clean aperture width in pixels of a video image
	cleanApertureWidthN		uint32
	//cleanApertureWidthD is the second part of a fraction that describes
	//the exact clean aperture width in pixels of a video image
	cleanApertureWidthD		uint32
	//cleanApertureHeightN is the first part of a fraction that describes
	//the exact clean aperture height in pixels of a video image
	cleanApertureHeightN	uint32
	//cleanApertureHeightD is the first part of a fraction that describes
	//the exact clean aperture height in pixels of a video image
	cleanApertureHeightD	uint32
	//horizonOffN is the first part of a fraction that describes the
	//horizontal offset of clean aperture centre minus (width-1)/2. Typically 0.
	horizOffN				uint32
	//horizonOffD is the second part of a fraction that describes the
	//horizontal offset of clean aperture centre minus (width-1)/2. Typically 0.
	horizOffD				uint32
	//vertOffN is the first part of a fraction that describes the
	//vertical offset of clean aperture centre minus (height-1)/2. Typically 0.
	vertOffN				uint32
	//vertOffD is the second part of a fraction that describes the
	//vertical offset of clean aperture centre minus (height-1)/2. Typically 0.
	vertOffD				uint32
}

//ColorInformationBox ('colr'); the if-else series should be added to a
//constructor
type ColorInformationBox struct {
	//extends Box
	size					uint32
	largeSize				uint64
	boxtype					uint32
	//colorType indicates the type of color information provided
	colorType				uint32
	//'nclx' are the 4 bytes defined for PTM_COLOR_INFO in A.7.2 ISO/IEC 29199-2
	if(colorType == 'nclx') {
		colorPrimaries				uint16
		transferCharacteristics		uint16
		matrixCoefficients			uint16
		fullRangeFlag				uint8 //only needs 1 bit
	} else if(colorType == 'rICC) {
		//ICCProfile as defined in ISO 15076-1 or ICC.1:2010 is provided
		ICCProfile 						//restricted ICC profile
	} else if(colorType == 'prof') {
		ICCProfile							//unrestricted ICC profile
	}
}

//VisualSampleEntry ('codingname')
type VisualSampleEntry struct {
	SampleEntry
	preDefined			uint16 		//set to 0 in spec
	const reserved		uint16  	//set to 0 in spec
	preDefined			[3]uint32	//set to 0 in the spec
	//width is the maximum width of the described stream in pixels
	width				uint16
	//height is the maximum height of the descibed stream in pixels
	height				uint16
	//horizresolution is in pixels-per-inch; fixed 16.16
	horizresolution		uint32	//template; set to 0x00480000 (72 dpi)
	//vertresolution is in pixels-per-inch; fixed 16.16
	vertresolution		uint32  //template; set to 0x00480000 (72 dpi)
	const reserved		uint32		//set to 0 in spec
	//frameCount number of frames per sample, default is 1, but can be higher
	frameCount			uint16	//template; set to 1 in spec
	//fixed 32-byte field, first byte is number of bytes displayed, 
	//then the number of bytes of displayable data, followed by
	//padding. The field can be 0.
	compressorName		[32]string
	//depth is the color density the value of 0x0018 means in color 
	//with no alpha and looks to be the only value it is set up to 
	//take at this time.
	depth					uint16	//template; set to 0x0018 in spec
	preDefined			int16		//set to -1 in spec
	clap				CleanApertureBox	//optional
	pasp				PixelAspectRatioBox //optional
}

//Audio sequences section _______________________________________________________________

//AudioSampleEntry ('codingname')
type AudioSampleEntry struct {
	SampleEntry
	const reserved			[2]unit32
	//channelCount is the number of audio channels, e.g. 1 is mono and 2 is stereo
	channelCount			uint16		//template; set to 2 in spec
	//sampleSize is in bits and the default size is 16
	sampleSize				uint16		//template; set to 16 in spec
	preDefined				uint16		//set to 0
	const reserved			uint16		//set to 0
	//sampleRate is the sampling rate in fixed 16.16 (hi.lo) format
	sampleRate				uint32		//template; set to {default samplerate of media}<<16
}

//SampleDescriptionBox (handler_type)
type SampleDescriptionBox struct {
	//extends FullBox('stsd',0,0)
	size			uint32
	largeSize		uint64
	boxtype			uint32
	version			uint8 //looks to be set to 0
	flags			[3]byte //looks to be set to 0
	i				int
	//entryCount is the number of entries in the table; loop control variable
	entryCount		uint32
	for i = 1; i <= entryCount; i++ {
		switch handlerType {
			case 'soun': { // for audio tracks
				AudioSampleEntry
			}
			case 'vide': {
				VisualSampleEntry // for video tracks
			}
			case 'hint': {
				HintSampleEntry // for hint track
			}
			case 'meta': {
				MetadataSampleEntry // for metadata
			}
		}
	}
}