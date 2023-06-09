package exporter

import (
	"fmt"
	"sort"
	"strings"

	al "github.com/philhanna/cwcomp/acrosslite"
)

// WriteDownClues writes the <DOWN> section
func WriteDownClues(pal *al.AcrossLite) string {
	const TAG = "<DOWN>"
	parts := make([]string, 0)
	clueMap := pal.GetDownClues()

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
