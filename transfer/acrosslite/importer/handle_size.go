package importer

import (
	"fmt"
	"regexp"
	"strconv"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

func HandleLookingForSize(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<SIZE>":
		return READING_SIZE, nil
	default:
		return UNKNOWN, fmt.Errorf("did not find <SIZE>")
	}
}

func HandleReadingSize(pal *al.AcrossLite, line string) (ParsingState, error) {
	reSize := regexp.MustCompile(`(\d+)x(\d+)`)
	tokens := reSize.FindStringSubmatch(line)
	switch {
	case tokens == nil:
		return UNKNOWN, fmt.Errorf("no <digits>x<digits> size expression found")
	case len(tokens)-1 != 2:
		return UNKNOWN, fmt.Errorf("expected two size integers, found %d", len(tokens)-1)
	case tokens[1] != tokens[2]:
		return UNKNOWN, fmt.Errorf("only square grids allowed, not %sx%s", tokens[1], tokens[2])
	default:
		pal.Size, _ = strconv.Atoi(tokens[1])
		return LOOKING_FOR_GRID, nil
	}

}