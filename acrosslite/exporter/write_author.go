package exporter

import (
	"fmt"
	al "github.com/philhanna/cwcomp/acrosslite"
)

// WriteAuthor writes the <AUTHOR> entry
func WriteAuthor(pal *al.AcrossLite) string {
	const TAG = `<AUTHOR>`
	return fmt.Sprintf("%s\n    %s", TAG, pal.GetAuthor())
}
