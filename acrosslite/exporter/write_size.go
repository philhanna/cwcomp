package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/acrosslite"
)

// WriteSize writes the <SIZE> entry
func WriteSize(pal *al.AcrossLite) string {
	const TAG = `<SIZE>`
	n := pal.GetSize()
	return fmt.Sprintf("%s\n    %dx%d", TAG, n, n)
}
