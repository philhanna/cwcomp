package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleDone(pal *al.AcrossLite, line string) (ParsingState, error) {
	return DONE, nil
}
