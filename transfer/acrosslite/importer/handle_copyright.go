package importer

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleLookingForCopyright verifies that the current line in the data
// is <COPYRIGHT>.
func HandleLookingForCopyright(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<COPYRIGHT>":
		return READING_COPYRIGHT, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <COPYRIGHT>")
	}
}

// HandleReadingCopyright copies the current line in the data to the
// AcrossLite structure.
func HandleReadingCopyright(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Copyright = line
	return LOOKING_FOR_SIZE, nil
}
