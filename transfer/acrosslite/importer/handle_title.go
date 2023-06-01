package importer

import (
	"fmt"
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleLookingForTitle(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<TITLE>":
		return READING_TITLE, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <TITLE>")
	}
}

func HandleReadingTitle(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Title = line
	return LOOKING_FOR_AUTHOR, nil
}