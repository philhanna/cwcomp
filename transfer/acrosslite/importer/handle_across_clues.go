package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleReadingAcross parses the across clues in the data, line by
// line, and and appends them to the list of across clues in the
// AcrossLite structure.
func HandleReadingAcross(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<DOWN>" {
		return READING_DOWN, nil
	}
	pal.AcrossClues = append(pal.AcrossClues, line)
	return READING_ACROSS, nil
}
