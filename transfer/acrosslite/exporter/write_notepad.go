package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteNotepad writes the <NOTEPAD> entry
func WriteNotepad(pal *al.AcrossLite) string {
	const TAG = "<NOTEPAD>"
	return fmt.Sprintf("%s\n%s", TAG, pal.GetNotepad())
}
