package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleReadingNotepad(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Notepad = append(pal.Notepad, line)
	return READING_NOTEPAD, nil
}
