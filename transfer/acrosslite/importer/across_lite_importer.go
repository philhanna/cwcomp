package importer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/philhanna/cwcomp"
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
func Parse(reader io.Reader) (*al.AcrossLite, error) {

	var err error

	// Initialize a new AcrossLite structure
	pal := al.NewAcrossLite()

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

	scanner := bufio.NewScanner(reader)
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

	// Since we don't know the word numbers until the parsing is
	// complete, we assigned temporary word numbers to all the clues. We
	// are now in a position to renumber the words and assign the keys
	// in the map correctly.

	patchWordNumbers(pal)

	return pal, nil
}

func patchWordNumbers(pal *al.AcrossLite) {

	// Create an n x n grid that we can renumber
	n := pal.GetSize()
	cells := make([][]byte, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]byte, n)
	}

	// Convert the grid strings into this simple byte matrix
	// so that we can calculate the word numbers.
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			letter := pal.Grid[i][j]
			if letter == '.' {
				letter = cwcomp.BLACK_CELL
			}
			cells[i][j] = letter
		}
	}

	// Now get the actual word numbers using the logic in the main package.
	numberedCells := cwcomp.GetNumberedCells(cells)

	// Internal function to create a new map of word numbers to clues in
	// the given direction.
	remap := func(direction cwcomp.Direction, oldClues map[int]string) map[int]string {

		// Get a sorted list of the old clue indices
		oldSeqs := make([]int, 0)
		for oldIndex := range oldClues {
			oldSeqs = append(oldSeqs, oldIndex)
		}
		sort.Ints(oldSeqs)

		// Internal function to return true if the given word number
		// starts a word in the given direction.
		isWordInThisDirection := func(nc cwcomp.NumberedCell, direction cwcomp.Direction) bool {
			switch direction {
			case cwcomp.ACROSS:
				return nc.StartA
			case cwcomp.DOWN:
				return nc.StartD
			}
			panic(direction)
		}

		// Get a sorted list of the new clue indices
		newSeqs := make([]int, 0)
		for _, nc := range numberedCells {
			if isWordInThisDirection(nc, direction) {
				newSeqs = append(newSeqs, nc.Seq)
			}
		}
		sort.Ints(newSeqs)

		// Now build a new map of sequence numbers to clues
		newClues := make(map[int]string)
		for oldSeq, newSeq := range newSeqs {
			newClues[newSeq] = oldClues[oldSeq]
		}
		return newClues
	}

	// Replace the old clue maps with the new ones
	pal.AcrossClues = remap(cwcomp.ACROSS, pal.AcrossClues)
	pal.DownClues = remap(cwcomp.DOWN, pal.DownClues)

}
