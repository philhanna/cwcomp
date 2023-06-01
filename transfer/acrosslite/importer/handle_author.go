package importer

import (
	"fmt"
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleLookingForAuthor(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<AUTHOR>":
		return READING_AUTHOR, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <AUTHOR>")
	}
}

func HandleReadingAuthor(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Author = line
	return LOOKING_FOR_COPYRIGHT, nil
}
