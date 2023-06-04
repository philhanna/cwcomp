package exporter

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteAcrossPuzzle writes the <ACROSS PUZZLE> line
func WriteAcrossPuzzle(pal *al.AcrossLite) string {
	const TAG = `<ACROSS PUZZLE>`
	return TAG
}
