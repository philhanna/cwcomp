package importer

import (
	"fmt"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// HandleLookingForGrid verifies that the current line in the data is
// <GRID>.
func HandleLookingForGrid(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<GRID>":
		pal.Grid = make([]string, 0)
		return READING_GRID, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <GRID>")
	}
}

// HandleReadingGrid stores grid lines in the AcrossLite structure.  It
// verifies that each line is of the right length, and that the final
// number of lines agrees with the declared size.
func HandleReadingGrid(pal *al.AcrossLite, line string) (ParsingState, error) {
	if line == "<ACROSS>" {

		// Verify that the number of grid lines agrees with the declared
		// size.
		if len(pal.Grid) != pal.Size {
			return UNKNOWN, fmt.Errorf(
				"found %d lines in <GRID> section, expected %d",
				len(pal.Grid), pal.Size)
		}
		return READING_ACROSS, nil
	}

	// Because I allow partially completed puzzles to be imported, it is
	// necessary to pad the line with spaces if it is not as long as the
	// declared size.
	for len(line) < pal.Size {
		line += " "
	}

	// But if the line is too long, that will be a fatal error
	if len(line) != pal.Size {
		return UNKNOWN, fmt.Errorf(
			"found %d characters in grid line, expected %d",
			len(line), pal.Size)
	}

	// Append the grid line to the AcrossLite grid list
	pal.Grid = append(pal.Grid, line)
	return READING_GRID, nil
}
