package exporter

import (
	"fmt"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteClues writes the <ACROSS> or <DOWN> section
func WriteClues(
	pal *al.AcrossLite,
	TAG string,
	f func(*al.AcrossLite, string) map[int]string) string {
	parts := make([]string, 0)
	for _, clue := range f(pal, TAG) {
		parts = append(parts, fmt.Sprintf("\t%s", clue))
	}
	section := strings.Join(parts, "\n")
	result := fmt.Sprintf("%s\n\t%s", TAG, section)
	return result
}
