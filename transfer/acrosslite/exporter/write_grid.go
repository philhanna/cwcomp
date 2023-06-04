package exporter

import (
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteGrid writes the <GRID> entry
func WriteGrid(pal *al.AcrossLite) string {
	const TAG = "<GRID>"
	n := pal.GetSize()
	parts := make([]string, n)
	for i, line := range pal.GetGrid() {
		parts[i] = "\t" + line
	}
	section := strings.Join(parts, "\n")
	result := TAG + "\n" + section
	return result
}
