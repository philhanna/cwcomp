package importer

import (
	"fmt"
	"regexp"
	"strconv"

	al "github.com/philhanna/cwcomp/acrosslite"
)

// HandleLookingForSize verifies that the current line in the data
// is <SIZE>.
func HandleLookingForSize(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<SIZE>":
		return READING_SIZE, nil
	default:
		return UNKNOWN, errNoSize
	}
}

// HandleReadingSize examines the current line, and verifies that it has
// the form <digits>x<digits>.  Note that the only grids I handle are
// square ones.
func HandleReadingSize(pal *al.AcrossLite, line string) (ParsingState, error) {
	reSize := regexp.MustCompile(`(\d+)x(\d+)`)
	tokens := reSize.FindStringSubmatch(line)
	switch {
	case tokens == nil:
		return UNKNOWN, fmt.Errorf("no <digits>x<digits> size expression found")
	case tokens[1] != tokens[2]:
		return UNKNOWN, fmt.Errorf("only square grids allowed, not %sx%s", tokens[1], tokens[2])
	default:
		pal.Size, _ = strconv.Atoi(tokens[1])
		return LOOKING_FOR_GRID, nil
	}

}
