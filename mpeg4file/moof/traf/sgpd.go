package traf

import "strconv"

type sgpd struct{
	// box
	size uint32
	largeSize uint64
	boxType uint32
	// fullbox
	version uint8
	flags uint8
	// sgpd box
	groupType uint32
	defaultLength uint32
	entryCount uint32
	descriptionLength []uint32
}

func (s *sgpd) String() string{
	return strconv.FormatUint(uint64(s.size),10)
}

/*
int i;
   for (i = 1 ; i <= entry_count ; i++){
      if (version==1) {
         if (default_length==0) {
            unsigned int(32) description_length;
         }
      }
      switch (handler_type){
} }
}
case ‘vide’: // for video tracks VisualSampleGroupEntry (grouping_type); break;
case ‘soun’: // for audio tracks AudioSampleGroupEntry(grouping_type); break;
case ‘hint’: // for hint tracks HintSampleGroupEntry(grouping_type); break;
*/