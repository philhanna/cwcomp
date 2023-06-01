package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleReadingDown(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<NOTEPAD>" {
		return READING_NOTEPAD, nil
	}
	pal.DownClues = append(pal.DownClues, line)
	return READING_DOWN, nil
}