package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleReadingAcross parses the across clues in the data, line by
// line, and and appends them to the list of across clues in the
// AcrossLite structure.
//
// Note that this function does not yet know the word numbering scheme,
// so we just assign consecutive word number. This will be patched up
// when the parsing is complete.
func HandleReadingAcross(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<DOWN>" {
		return READING_DOWN, nil
	}
	tempKey := len(pal.AcrossClues)
	pal.AcrossClues[tempKey] = line
	return READING_ACROSS, nil
}
