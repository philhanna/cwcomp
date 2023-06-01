package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleReadingDown parses the down clues in the data, line by line.
// This is a valid final state, because <NOTEPAD> is an optional
// element.
func HandleReadingDown(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<NOTEPAD>" {
		return READING_NOTEPAD, nil
	}
	pal.DownClues = append(pal.DownClues, line)
	return READING_DOWN, nil
}
