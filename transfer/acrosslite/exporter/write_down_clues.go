package exporter

import (
	"fmt"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteDownClues writes the <DOWN> section
func WriteDownClues(pal *al.AcrossLite) string {
	const TAG = "<DOWN>"
	parts := make([]string, 0)
	for _, clue := range pal.GetDownClues() {
		parts = append(parts, fmt.Sprintf("    %s", clue))
	}
	section := strings.Join(parts, "\n")
	result := fmt.Sprintf("%s\n%s", TAG, section)
	return result
}
