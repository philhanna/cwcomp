package export

import (
	"bufio"
	"fmt"
	"log"
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

	// Initialize a new AcrossLite structure
	pal := new(AcrossLite)
	pal.grid = make([]string, 0)
	pal.acrossClues = make([]string, 0)
	pal.downClues = make([]string, 0)
	pal.notepad = make([]string, 0)

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
			pal.title = line
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
			pal.author = line
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
			pal.copyright = line
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
			reSize := regexp.MustCompile(`(\d+)x(\d)`)
			tokens := reSize.FindStringSubmatch(line)
			switch {
			case tokens == nil:
				return UNKNOWN, fmt.Errorf("no <digits>x<digits> size expression found")
			case len(tokens) != 2:
				return UNKNOWN, fmt.Errorf("expected two size integers, found %d", len(tokens))
			case tokens[0] != tokens[1]:
				return UNKNOWN, fmt.Errorf("only square grids allowed, not %sx%s", tokens[0], tokens[1])
			default:
				pal.size, _ = strconv.Atoi(tokens[0])
				return LOOKING_FOR_GRID, nil
			}
		},

		LOOKING_FOR_GRID: func(line string) (ParsingState, error) {
			switch line {
			case "<GRID>":
				pal.grid = make([]string, 0)
				return READING_GRID, nil
			default:
				return UNKNOWN, fmt.Errorf("did not find <GRID>")
			}
		},

		READING_GRID: func(line string) (ParsingState, error) {
			if line == "<ACROSS>" {
				if len(pal.grid) != pal.size {
					return UNKNOWN, fmt.Errorf(
						"found %d lines in <GRID> section, expected %d",
						len(pal.grid), pal.size)
				}
				return READING_ACROSS, nil
			}
			if len(line) != pal.size {
				return UNKNOWN, fmt.Errorf(
					"found %d characters in grid line, expected %d",
					len(line), pal.size)
			}
			pal.grid = append(pal.grid, line)
			return READING_GRID, nil
		},

		READING_ACROSS: func(line string) (ParsingState, error) {
			if line == "<DOWN>" {
				return READING_DOWN, nil
			}
			pal.acrossClues = append(pal.acrossClues, line)
			return READING_ACROSS, nil
		},

		READING_DOWN: func(line string) (ParsingState, error) {
			if line == "<NOTEPAD>" {
				return READING_NOTEPAD, nil
			}
			pal.downClues = append(pal.downClues, line)
			return READING_DOWN, nil
		},

		READING_NOTEPAD: func(line string) (ParsingState, error) {
			pal.notepad = append(pal.notepad, line)
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

	if state != DONE {
		// TODO change this to a switch statement for all states
		return nil, fmt.Errorf("Expected final state to be DONE, not %v", state)
	}

	return pal, nil
}
