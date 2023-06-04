package exporter

import (
	"fmt"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteAcrossClues writes the <ACROSS> section
func WriteAcrossClues(pal *al.AcrossLite) string {
	const TAG = "<ACROSS>"
	parts := make([]string, 0)
	for _, clue := range pal.GetAcrossClues() {
		parts = append(parts, fmt.Sprintf("\t%s", clue))
	}
	section := strings.Join(parts, "\n")
	result := fmt.Sprintf("%s\n\t%s", TAG, section)
	return result
}
