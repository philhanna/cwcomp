package importer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Import parses the text in an AcrossLite puzzle
func Import(filename string) (*al.AcrossLite, error) {

	// Initialize a new AcrossLite structure
	pal := new(al.AcrossLite)
	pal.Grid = make([]string, 0)
	pal.AcrossClues = make([]string, 0)
	pal.DownClues = make([]string, 0)
	pal.Notepad = make([]string, 0)

	// Describe the parsing states of the finite state machine
	type ParsingState byte
	const (
		UNKNOWN ParsingState = iota
		INIT
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

	// Define the map of parsing states to handler functions
	type Handler func(string) (ParsingState, error)

	stateMap := map[ParsingState]Handler{

		INIT: func(line string) (ParsingState, error) {
			switch line {
			case "<ACROSS PUZZLE>":
				return LOOKING_FOR_TITLE, nil
			default:
				return UNKNOWN, fmt.Errorf("Valid AcrossLite file must start with <ACROSS PUZZLE>, not %s", line)
			}
		},

		LOOKING_FOR_TITLE: func(line string) (ParsingState, error) {
			switch line {
			case "<TITLE>":
				return READING_TITLE, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <TITLE>")
			}
		},

		READING_TITLE: func(line string) (ParsingState, error) {
			pal.Title = line
			return LOOKING_FOR_AUTHOR, nil
		},

		LOOKING_FOR_AUTHOR: func(line string) (ParsingState, error) {
			switch line {
			case "<AUTHOR>":
				return READING_AUTHOR, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <AUTHOR>")
			}
		},

		READING_AUTHOR: func(line string) (ParsingState, error) {
			pal.Author = line
			return LOOKING_FOR_COPYRIGHT, nil
		},

		LOOKING_FOR_COPYRIGHT: func(line string) (ParsingState, error) {
			switch line {
			case "<COPYRIGHT>":
				return READING_COPYRIGHT, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <COPYRIGHT>")
			}
		},

		READING_COPYRIGHT: func(line string) (ParsingState, error) {
			pal.Copyright = line
			return LOOKING_FOR_SIZE, nil
		},

		LOOKING_FOR_SIZE: func(line string) (ParsingState, error) {
			switch line {
			case "<SIZE>":
				return READING_SIZE, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <SIZE>")
			}
		},

		READING_SIZE: func(line string) (ParsingState, error) {
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
		},

		LOOKING_FOR_GRID: func(line string) (ParsingState, error) {
			switch line {
			case "<GRID>":
				pal.Grid = make([]string, 0)
				return READING_GRID, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <GRID>")
			}
		},

		READING_GRID: func(line string) (ParsingState, error) {
			if line == "<ACROSS>" {
				if len(pal.Grid) != pal.Size {
					return UNKNOWN, fmt.Errorf(
						"found %d lines in <GRID> section, expected %d",
						len(pal.Grid), pal.Size)
				}
				return READING_ACROSS, nil
			}
			if len(line) != pal.Size {
				return UNKNOWN, fmt.Errorf(
					"found %d characters in grid line, expected %d",
					len(line), pal.Size)
			}
			pal.Grid = append(pal.Grid, line)
			return READING_GRID, nil
		},

		READING_ACROSS: func(line string) (ParsingState, error) {
			if line == "<DOWN>" {
				return READING_DOWN, nil
			}
			pal.AcrossClues = append(pal.AcrossClues, line)
			return READING_ACROSS, nil
		},

		READING_DOWN: func(line string) (ParsingState, error) {
			if line == "<NOTEPAD>" {
				return READING_NOTEPAD, nil
			}
			pal.DownClues = append(pal.DownClues, line)
			return READING_DOWN, nil
		},

		READING_NOTEPAD: func(line string) (ParsingState, error) {
			pal.Notepad = append(pal.Notepad, line)
			return READING_NOTEPAD, nil
		},

		DONE: func(line string) (ParsingState, error) {
			return DONE, nil
		},
	}

	// Open the file
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// -----------------------------------------------------------------
	// Now run the finite state machine:
	//
	// - Set initial state to INIT
	// - For each line in the file:
	//   - Lookup the handler function
	//   - Call the function with this line
	//   - Set the new state
	// - Check for a valid ending state
	// -----------------------------------------------------------------

	state := INIT

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		handler, found := stateMap[state]
		if found {
			state, err = handler(line)
			if err != nil {
				log.Println(err)
				return nil, err
			}
		}
	}

	// Check the final state

	switch state {
	default:
		return nil, fmt.Errorf("unexpected final state %v", state)
	case INIT:
		return nil, fmt.Errorf("no lines in file")
	case LOOKING_FOR_TITLE:
		return nil, fmt.Errorf("never found <TITLE>")
	case READING_TITLE:
		return nil, fmt.Errorf("unexpected final state READING_TITLE")
	case LOOKING_FOR_AUTHOR:
		return nil, fmt.Errorf("never found <AUTHOR>")
	case READING_AUTHOR:
		return nil, fmt.Errorf("unexpected final state READING_AUTHOR")
	case LOOKING_FOR_COPYRIGHT:
		return nil, fmt.Errorf("never found <COPYRIGHT>")
	case READING_COPYRIGHT:
		return nil, fmt.Errorf("unexpected final state READING_COPYRIGHT")
	case LOOKING_FOR_SIZE:
		return nil, fmt.Errorf("never found <SIZE>")
	case READING_SIZE:
		return nil, fmt.Errorf("unexpected final state READING_SIZE")
	case LOOKING_FOR_GRID:
		return nil, fmt.Errorf("never found <GRID>")
	case READING_GRID:
		return nil, fmt.Errorf("unexpected final state READING_GRID")
	case READING_ACROSS:
		return nil, fmt.Errorf("unexpected final state READING_ACROSS")
	case READING_DOWN, READING_NOTEPAD, DONE:
		// OK
	}

	return pal, nil
}
