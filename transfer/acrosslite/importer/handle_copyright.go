package importer

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleLookingForCopyright(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<COPYRIGHT>":
		return READING_COPYRIGHT, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <COPYRIGHT>")
	}
}

func HandleReadingCopyright(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Copyright = line
	return LOOKING_FOR_SIZE, nil
}
