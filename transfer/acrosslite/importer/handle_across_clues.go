package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleReadingAcross(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<DOWN>" {
		return READING_DOWN, nil
	}
	pal.AcrossClues = append(pal.AcrossClues, line)
	return READING_ACROSS, nil
}