package importer

import (
	al "github.com/philhanna/cwcomp/acrosslite"
)

// HandleReadingDown parses the down clues in the data, line by line.
// This is a valid final state, because <NOTEPAD> is an optional
// element.
//
// Note that this function does not yet know the word numbering scheme,
// so we just assign consecutive word number. This will be patched up
// when the parsing is complete.
func HandleReadingDown(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<NOTEPAD>" {
		return READING_NOTEPAD, nil
	}
	tempKey := len(pal.DownClues)
	pal.DownClues[tempKey] = line
	return READING_DOWN, nil
}
