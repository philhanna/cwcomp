package importer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Signature of a function that handles a string in a particular state
type Handler func(*al.AcrossLite, string) (ParsingState, error)

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

// Define the map of parsing states to handler functions
var StateMap = map[ParsingState]Handler{
	INIT:                  HandleInit,
	LOOKING_FOR_TITLE:     HandleLookingForTitle,
	READING_TITLE:         HandleReadingTitle,
	LOOKING_FOR_AUTHOR:    HandleLookingForAuthor,
	READING_AUTHOR:        HandleReadingAuthor,
	LOOKING_FOR_COPYRIGHT: HandleLookingForCopyright,
	READING_COPYRIGHT:     HandleReadingCopyright,
	LOOKING_FOR_SIZE:      HandleLookingForSize,
	READING_SIZE:          HandleReadingSize,
	LOOKING_FOR_GRID:      HandleLookingForGrid,
	READING_GRID:          HandleReadingGrid,
	READING_ACROSS:        HandleReadingAcross,
	READING_DOWN:          HandleReadingDown,
	READING_NOTEPAD:       HandleReadingNotepad,
	DONE:                  HandleDone,
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Parse parses the text in an AcrossLite puzzle
func Parse(filename string) (*al.AcrossLite, error) {

	// Initialize a new AcrossLite structure
	pal := al.NewAcrossLite()

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
		handler, found := StateMap[state]
		if found {
			state, err = handler(pal, line)
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
