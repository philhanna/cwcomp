package importer

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleReadingNotepad takes each line and appends it to the AcrossLite
// structure notepad list.  Note that this is an optional section, and
// it doesn't require any validation.
func HandleReadingNotepad(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Notepad += line
	return READING_NOTEPAD, nil
}
