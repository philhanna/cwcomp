package importer

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strings"

	"github.com/philhanna/cwcomp"
	al "github.com/philhanna/cwcomp/acrosslite"
	"github.com/philhanna/cwcomp/model"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Signature of a function that handles a string in a particular state
type Handler func(*al.AcrossLite, string) (ParsingState, error)

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

var (
	errNoLines          = errors.New("no lines in file")
	errNoTitle          = errors.New("never found <TITLE>")
	errReadingTitle     = errors.New("unexpected final state READING_TITLE")
	errNoAuthor         = errors.New("never found <AUTHOR>")
	errReadingAuthor    = errors.New("unexpected final state READING_AUTHOR")
	errNoCopyright      = errors.New("never found <COPYRIGHT>")
	errReadingCopyright = errors.New("unexpected final state READING_COPYRIGHT")
	errNoSize           = errors.New("never found <SIZE>")
	errReadingSize      = errors.New("unexpected final state READING_SIZE")
	errNoGrid           = errors.New("never found <GRID>")
	errReadingGrid      = errors.New("unexpected final state READING_GRID")
	errReadingAcross    = errors.New("unexpected final state READING_ACROSS")
)

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
}

// Define the map of invalid final states to their error messages
var FinalStateMap = map[ParsingState]error{
	INIT:                  errNoLines,
	LOOKING_FOR_TITLE:     errNoTitle,
	READING_TITLE:         errReadingTitle,
	LOOKING_FOR_AUTHOR:    errNoAuthor,
	READING_AUTHOR:        errReadingAuthor,
	LOOKING_FOR_COPYRIGHT: errNoCopyright,
	READING_COPYRIGHT:     errReadingCopyright,
	LOOKING_FOR_SIZE:      errNoSize,
	READING_SIZE:          errReadingSize,
	LOOKING_FOR_GRID:      errNoGrid,
	READING_GRID:          errReadingGrid,
	READING_ACROSS:        errReadingAcross,
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
		line = strings.ReplaceAll(line, "_", " ")
		handler, found := StateMap[state]
		if found {
			state, err = handler(pal, line)
			if err != nil {
				return nil, err
			}
		}
	}

	// Check the final state

	err = FinalStateMap[state]
	if err != nil {
		return nil, err
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
	numberedCells := model.GetNumberedCells(cells)

	// Internal function to create a new map of word numbers to clues in
	// the given direction.
	remap := func(direction model.Direction, oldClues map[int]string) map[int]string {

		// Get a sorted list of the old clue indices
		oldSeqs := make([]int, 0)
		for oldIndex := range oldClues {
			oldSeqs = append(oldSeqs, oldIndex)
		}
		sort.Ints(oldSeqs)

		// Internal function to return true if the given word number
		// starts a word in the given direction.
		isWordInThisDirection := func(nc model.NumberedCell, direction model.Direction) bool {
			var result bool
			switch direction {
			case model.ACROSS:
				result = nc.StartA
			case model.DOWN:
				result = nc.StartD
			}
			return result
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
	pal.AcrossClues = remap(model.ACROSS, pal.AcrossClues)
	pal.DownClues = remap(model.DOWN, pal.DownClues)

}
