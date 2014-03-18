package moov


//ChunkOffsetBox is a table that gives the index of each chunk where it
//occurs in the containing file. There are 32-bit and 63-nit variants; this
//is the 32-bit variant. These are file offsets and are not evaluated 
//against any of the boxes.
type Stco package {
	//extends FullBox('stco', version = 0, 0)
	size				uint32
	largeSize			uint64
	boxtype				uint32
	version				uint8
	flags				[3]byte
	//entryCount is the number of chunk offsets in the table; loop control variable
	entryCount			uint32
	for i:=1; i <=entryCount; i++ {
		//chunkOffset is the offset of the start of the chunk into its containing
		//media file.
		chunkOffset			uint32
	}
}