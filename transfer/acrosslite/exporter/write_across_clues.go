package exporter

import (
	"fmt"
	"sort"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// WriteAcrossClues writes the <ACROSS> section
func WriteAcrossClues(pal *al.AcrossLite) string {
	const TAG = "<ACROSS>"
	parts := make([]string, 0)
	clueMap := pal.GetAcrossClues()

	// Must do this in sorted order
	keys := []int{}
	for seq := range clueMap {
		keys = append(keys, seq)
	}
	sort.Ints(keys)
	for _, seq := range keys {
		clue := clueMap[seq]
		parts = append(parts, fmt.Sprintf("    %s", clue))
	}
	section := strings.Join(parts, "\n")
	result := fmt.Sprintf("%s\n%s", TAG, section)
	return result
}
