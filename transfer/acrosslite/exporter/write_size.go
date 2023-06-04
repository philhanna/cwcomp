package exporter

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteSize writes the <SIZE> entry
func WriteSize(pal *al.AcrossLite) string {
	const TAG = `<SIZE>`
	n := pal.GetSize()
	return fmt.Sprintf("%s\n\t%dx%d\n", TAG, n, n)
}
