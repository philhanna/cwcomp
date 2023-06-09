package exporter

import (
	"strings"

	al "github.com/philhanna/cwcomp/acrosslite"
)

// WriteGrid writes the <GRID> entry
func WriteGrid(pal *al.AcrossLite) string {
	const TAG = "<GRID>"
	n := pal.GetSize()
	parts := make([]string, n)
	for i, line := range pal.GetGrid() {
		parts[i] = "    " + line
	}
	section := strings.Join(parts, "\n")
	result := TAG + "\n" + section
	return result
}
