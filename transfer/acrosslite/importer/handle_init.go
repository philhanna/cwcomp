package importer

import (
	"fmt"
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleInit looks for the valid beginning of an AcrossLite text file,
// which is <ACROSS PUZZLE>.
func HandleInit(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<ACROSS PUZZLE>":
		return LOOKING_FOR_TITLE, nil
	default:
		return UNKNOWN, fmt.Errorf("Valid AcrossLite file must start with <ACROSS PUZZLE>, not %s", line)
	}
}
