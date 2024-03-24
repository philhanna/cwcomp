package importer

import (
	al "github.com/philhanna/cwcomp/acrosslite"
)

// HandleLookingForTitle verifies that the current line in the data is
// <TITLE>.
func HandleLookingForTitle(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<TITLE>":
		return READING_TITLE, nil
	default:
		return UNKNOWN, errNoTitle
	}
}

// HandleReadingTitle copies the line into the Title element of the
// AcrossLite structure.
func HandleReadingTitle(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Title = line
	return LOOKING_FOR_AUTHOR, nil
}
