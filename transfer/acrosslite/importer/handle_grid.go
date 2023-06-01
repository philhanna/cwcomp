package importer

import (
	"fmt"
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleLookingForGrid(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<GRID>":
		pal.Grid = make([]string, 0)
		return READING_GRID, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <GRID>")
	}
}

func HandleReadingGrid(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<ACROSS>" {
		if len(pal.Grid) != pal.Size {
			return UNKNOWN, fmt.Errorf(
				"found %d lines in <GRID> section, expected %d",
				len(pal.Grid), pal.Size)
		}
		return READING_ACROSS, nil
	}
	for len(line) < pal.Size {
		line += " "
	}
	if len(line) != pal.Size {
		return UNKNOWN, fmt.Errorf(
			"found %d characters in grid line, expected %d",
			len(line), pal.Size)
	}
	pal.Grid = append(pal.Grid, line)
	return READING_GRID, nil
}
