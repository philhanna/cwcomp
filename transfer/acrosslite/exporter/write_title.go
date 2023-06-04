package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteTitle writes the <TITLE> entry
func WriteTitle(pal *al.AcrossLite) string {
	const TAG = `<TITLE>`
	return fmt.Sprintf("%s\n    %s", TAG, pal.GetTitle())
}
