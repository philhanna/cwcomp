package export

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Import parses the text in an AcrossLite puzzle
func Import(filename string) (*AcrossLite, error) {
	pal := new(AcrossLite)

	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	type ParsingState byte
	const (
		INIT ParsingState = iota
		LOOKING_FOR_TITLE
		READING_TITLE
		LOOKING_FOR_AUTHOR
		READING_AUTHOR
		LOOKING_FOR_COPYRIGHT
		READING_COPYRIGHT
		LOOKING_FOR_SIZE
		READING_SIZE
		LOOKING_FOR_GRID
		READING_GRID
		READING_ACROSS
		READING_DOWN
		READING_NOTEPAD
		DONE
	)
	state := INIT
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		switch state {

		case INIT:
			if line == "<ACROSS PUZZLE>" {
				state = LOOKING_FOR_TITLE
			} else {
				return nil, fmt.Errorf("Not a valid AcrossLite text file: %s", line)
			}

		case LOOKING_FOR_TITLE:
			if line == "<TITLE>" {
				state = READING_TITLE
			} else {
				return nil, fmt.Errorf("did not find <TITLE>")
			}

		case READING_TITLE:
			pal.title = line
			state = LOOKING_FOR_AUTHOR

		case LOOKING_FOR_AUTHOR:
			if line == "<AUTHOR>" {
				state = READING_AUTHOR
			} else {
				return nil, fmt.Errorf("did not find <AUTHOR>")
			}

		case READING_AUTHOR:
			pal.author = line
			state = LOOKING_FOR_COPYRIGHT

		case LOOKING_FOR_COPYRIGHT:
			if line == "<COPYRIGHT>" {
				state = READING_COPYRIGHT
			} else {
				return nil, fmt.Errorf("did not find <COPYRIGHT>")
			}

		case READING_COPYRIGHT:
			pal.copyright = line
			state = LOOKING_FOR_SIZE

		case LOOKING_FOR_SIZE:
			if line == "<SIZE>" {
				state = READING_SIZE
			} else {
				return nil, fmt.Errorf("did not find <SIZE>")
			}

		case READING_SIZE:
			reSize := regexp.MustCompile(`(\d+)x(\d)`)
			tokens := reSize.FindStringSubmatch(line)
			if tokens == nil {
				return nil, fmt.Errorf("no <digits>x<digits> size expression found")
			}
			if len(tokens) != 2 {
				return nil, fmt.Errorf("expected two size integers, found %d", len(tokens))
			}
			if tokens[0] != tokens[1] {
				return nil, fmt.Errorf("only square grids allowed, not %sx%s", tokens[0], tokens[1])
			}
			pal.size, _ = strconv.Atoi(tokens[0])
			state = LOOKING_FOR_GRID

		case LOOKING_FOR_GRID:
			if line == "<GRID>" {
				state = READING_GRID
				pal.grid = make([]string, 0)
			} else {
				return nil, fmt.Errorf("did not find <GRID>")
			}

		case READING_GRID:
			if line == "<ACROSS>" {
				if len(pal.grid) != pal.size {
					return nil, fmt.Errorf("Must be %d lines in grid, not %d", pal.size, len(pal.grid))
				}
				state = READING_ACROSS
				pal.acrossClues = make([]string, 0)
			} else {
				return nil, fmt.Errorf("did not find <ACROSS>")
			}

		case READING_ACROSS:
		}
	}

	return pal, nil
}
