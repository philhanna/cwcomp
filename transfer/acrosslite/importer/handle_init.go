package importer

import (
	"fmt"
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleInit(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<ACROSS PUZZLE>":
		return LOOKING_FOR_TITLE, nil
	default:
		return UNKNOWN, fmt.Errorf("Valid AcrossLite file must start with <ACROSS PUZZLE>, not %s", line)
	}
}
