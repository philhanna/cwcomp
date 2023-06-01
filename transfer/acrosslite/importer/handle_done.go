package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleDone is a no-op that continues in the DONE state.
func HandleDone(pal *al.AcrossLite, line string) (ParsingState, error) {
	return DONE, nil
}
