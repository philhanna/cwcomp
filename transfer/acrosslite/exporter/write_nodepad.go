package exporter

import (
	"fmt"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteNotepad writes the <NOTEPAD> entry with the creation and
// modification dates
func WriteNotepad(pal *al.AcrossLite) string {
	const TAG = `<NOTEPAD>`
	parts := make([]string, 2)
	parts[0] = fmt.Sprintf("%q:%q", "created", pal.GetCreatedDate().Format(al.ISO8601))
	parts[1] = fmt.Sprintf("%q:%q", "modified", pal.GetModifiedDate().Format(al.ISO8601))
	jsonstr := "{" + strings.Join(parts, ",") + "}"
	result := fmt.Sprintf("%s\n%s", TAG, jsonstr)
	return result
}
