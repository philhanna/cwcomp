package cwcomp

import (
	"errors"
	"fmt"

	"github.com/philhanna/stack"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Puzzle contains the cells of a grid.
//
// The grid is constructed with the single parameter n, which is the size
// (n x n) of the grid.
//
// Any of the cells in the puzzle can be "black cells", which act as the
// boundaries of where the words can go. The model automatically takes
// care of matching a black cell with its symmetric twin 180 degrees
// from it.
//
// Wherever an across or down word starts, the puzzle assigns the next
// available word number to the cell and keeps track of the lengths of
// the across and down words.
//
// Puzzle supports a full "undo/redo" capability for the current session
// (from load to save).  Any black cell additions or deletions are
// pushed on an undo stack.
type Puzzle struct {
	n              int                 // Size of the grid (n x n square)
	puzzleName     string              // The puzzle name
	cells          [][]Cell            // Black cells and letter cells
	words          []*Word             // Pointers to the words in this grid
	wordNumbers    []*WordNumber       // Word number pointers
	undoPointStack stack.Stack[Point]  // Undo stack for black cells
	redoPointStack stack.Stack[Point]  // Redo stack for black cells
	undoWordStack  stack.Stack[Doable] // Undo stack for whole words
	redoWordStack  stack.Stack[Doable] // Redo stack for whold words
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewPuzzle creates a puzzle of the specified size.
func NewPuzzle(n int) *Puzzle {
	g := new(Puzzle)
	g.n = n

	// Create an n x n matrix of cell objects
	g.cells = make([][]Cell, n)
	for i := 0; i < n; i++ {
		g.cells[i] = make([]Cell, n)
		for j := 0; j < n; j++ {
			point := NewPoint(i+1, j+1)
			cell := NewLetterCell(point)
			g.cells[i][j] = cell
		}
	}

	g.undoPointStack = stack.NewStack[Point]()
	g.redoPointStack = stack.NewStack[Point]()

	return g
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Equal returns true if this puzzle is essentially equal to the other.
func (puzzle *Puzzle) Equal(other *Puzzle) bool {
	if other == nil {
		return false
	}
	thisString := puzzle.String()
	thatString := other.String()
	return thisString == thatString
}

// GetCell returns the cell at the specified point, which may be a black
// cell or a letter cell.
func (puzzle *Puzzle) GetCell(point Point) Cell {
	err := puzzle.ValidIndex(point)
	if err != nil {
		errmsg := fmt.Sprintf("%s is not a valid point", point.String())
		panic(errmsg)
	}
	x, y := point.ToXY()
	return puzzle.cells[y][x]
}

// GetClue returns the clue for the word.
func (puzzle *Puzzle) GetClue(word *Word) (string, error) {
	err := puzzle.wordPointerIsValid(word)
	if err != nil {
		return "", err
	}
	return word.clue, nil
}

// GetPuzzleName returns the puzzle name
func (puzzle *Puzzle) GetPuzzleName() string {
	return puzzle.puzzleName
}

// GetLength returns the length of the word.
func (puzzle *Puzzle) GetLength(word *Word) (int, error) {
	err := puzzle.wordPointerIsValid(word)
	if err != nil {
		return 0, err
	}
	return word.length, nil
}

// GetLetter returns the value of the cell at this point.  The length of
// the returned value is always 1, unless the point refers to a black
// cell, in which case the length is zero.
func (puzzle *Puzzle) GetLetter(point Point) string {
	letter := ""
	cell := puzzle.GetCell(point)
	switch typedcell := cell.(type) {
	case LetterCell:
		letter = typedcell.letter
		if letter == "" {
			letter = " "
		}
	}
	return letter
}

// GetText returns the text of the word.
func (puzzle *Puzzle) GetText(word *Word) string {
	err := puzzle.wordPointerIsValid(word)
	if err != nil {
		return ""
	}

	length, _ := puzzle.GetLength(word)

	var s string
	var point Point
	for i := 0; i < length; i++ {
		switch word.direction {
		case ACROSS:
			point = NewPoint(word.point.r, word.point.c+i)
		case DOWN:
			point = NewPoint(word.point.r+i, word.point.c)
		}
		letter := puzzle.GetLetter(point)
		s += letter
	}
	return s
}

// ImportPuzzle creates a puzzle from an external source
func ImportPuzzle(source Importer) (*Puzzle, error) {
	n := source.GetSize()
	puzzle := NewPuzzle(n)
	puzzle.puzzleName = source.GetName()
	// TODO what to do about puzzle title?
	for point := range puzzle.PointIterator() {
		r, c := point.r, point.c
		value, err := source.GetCell(r, c)
		if err != nil {
			err := fmt.Errorf("Point(%s): %v", point.String(), err)
			return nil, err
		}

		// Switch to zero-based indices to set cell values
		i, j := r-1, c-1
		if value == BLACK_CELL {
			// Black cell
			puzzle.cells[i][j] = NewBlackCell(point)
		} else {
			 // Letter cell
			lc := NewLetterCell(point)
			lc.letter = string(value)
			puzzle.cells[i][j] = lc
		}
	}
	puzzle.RenumberCells()

	// Now set the clues

	// Across words
	for index, clue := range source.GetAcrossClues() {
		seq := index + 1 // Word numbers are 1-based, not 0-based
		word := puzzle.LookupWordByNumber(seq, ACROSS)
		if word == nil {
			return nil, fmt.Errorf("no %d across word", seq)
		}
		err := puzzle.SetClue(word, clue)
		if err != nil {
			return nil, err
		}
	}

	// Down words
	for index, clue := range source.GetDownClues() {
		seq := index + 1 // Word numbers are 1-based, not 0-based
		word := puzzle.LookupWordByNumber(seq, DOWN)
		if word == nil {
			return nil, fmt.Errorf("no %d down word", seq)
		}
		err := puzzle.SetClue(word, clue)
		if err != nil {
			return nil, err
		}
	}

	return puzzle, nil
}

// LookupWord returns the word containing this point and direction
func (puzzle *Puzzle) LookupWord(point Point, dir Direction) *Word {
	for _, word := range puzzle.words {
		if word.direction == dir {
			for wPoint := range puzzle.WordIterator(word.point, dir) {
				if wPoint == point {
					return word
				}
			}
		}
	}
	return nil
}

// LookupWordByNumber returns the word at this point and direction
func (puzzle *Puzzle) LookupWordByNumber(seq int, dir Direction) *Word {
	wn := puzzle.LookupWordNumber(seq)
	if wn == nil {
		return nil
	}
	for _, word := range puzzle.words {
		if word.point == wn.point {
			if word.direction == dir {
				return word
			}
		}
	}
	return nil // No word for that seq+dir
}

// LookupWordNumber returns the WordNumber for this number
func (puzzle *Puzzle) LookupWordNumber(seq int) *WordNumber {
	for _, wn := range puzzle.wordNumbers {
		if wn.seq == seq {
			return wn
		}
	}
	return nil
}

// LookupWordNumberForStartingPoint returns the WordNumber starting at
// this point.
func (puzzle *Puzzle) LookupWordNumberForStartingPoint(point Point) *WordNumber {
	for _, wn := range puzzle.wordNumbers {
		if wn.point == point {
			return wn
		}
	}
	return nil
}

// RenumberCells assigns the word numbers based on the locations of the
// black cells.
func (puzzle *Puzzle) RenumberCells() {

	// Get the word numbers
	cells := PuzzleToSimpleMatrix(puzzle)
	ncs := GetNumberedCells(cells)
	puzzle.wordNumbers = make([]*WordNumber, len(ncs))
	for i, nc := range ncs {
		wn := new(WordNumber)
		wn.seq = nc.Seq
		wn.point = NewPoint(nc.Row, nc.Col)
		puzzle.wordNumbers[i] = wn
	}

	// Get the words list and add lengths
	puzzle.words = make([]*Word, 0)
	for _, nc := range ncs {
		point := NewPoint(nc.Row, nc.Col)
		if nc.StartA {
			word := NewWord(point, ACROSS, 0, "")
			for range puzzle.WordIterator(point, ACROSS) {
				word.length++
			}
			puzzle.words = append(puzzle.words, word)
		}
		if nc.StartD {
			word := NewWord(point, DOWN, 0, "")
			for range puzzle.WordIterator(point, DOWN) {
				word.length++
			}
			puzzle.words = append(puzzle.words, word)
		}
	}
}

// SetCell sets the cell at the specified point
func (puzzle *Puzzle) SetCell(point Point, cell Cell) {
	x, y := point.ToXY()
	puzzle.cells[y][x] = cell
}

// SetClue sets the specified clue in the specified word.
func (puzzle *Puzzle) SetClue(word *Word, clue string) error {
	if word == nil {
		return fmt.Errorf("word pointer is nil")
	}
	word.clue = clue
	return nil
}

// SetLetter sets the letter value of the cell at the specified point
func (puzzle *Puzzle) SetLetter(point Point, letter string) {
	cell := puzzle.GetCell(point)
	switch typedCell := cell.(type) {
	case LetterCell:
		typedCell.letter = letter
		puzzle.SetCell(point, typedCell)
	}
}

// SetPuzzleName sets the puzzle name
func (puzzle *Puzzle) SetPuzzleName(name string) {
	puzzle.puzzleName = name
}

// SetText sets the text in the puzzle for a specified word.
func (puzzle *Puzzle) SetText(word *Word, text string) error {

	// Make sure this is a valid word pointer
	err := puzzle.wordPointerIsValid(word)
	if err != nil {
		return err
	}

	// Make sure the text is not longer than the word allows
	if len(text) > word.length {
		errmsg := fmt.Sprintf(`Text %q length > expected %d`, len(text), word.length)
		err := errors.New(errmsg)
		return err
	}

	// Create a new Doable for this word and push it on the undoWordStack
	doable := NewDoable(puzzle, word)
	puzzle.undoWordStack.Push(doable)

	// Pad the text with blanks if too short
	for len(text) < word.length {
		text += " "
	}

	// Iterate through the points of the word, storing the text into it
	// letter by letter.
	puzzle.SetTextWithoutPush(word, text)

	// OK
	return nil
}

func (puzzle *Puzzle) SetTextWithoutPush(word *Word, text string) {
	i := 0
	for point := range puzzle.WordIterator(word.point, word.direction) {
		ch := text[i]
		i++
		letter := string(ch)
		puzzle.SetLetter(point, letter)
	}
}

// String returns a string representation of the puzzle
func (puzzle *Puzzle) String() string {
	n := puzzle.n
	sb := ""
	if puzzle.puzzleName == "" {
		sb += "(Untitled)"
	} else {
		sb += puzzle.puzzleName
	}
	sb += "\n"

	// Row of column numbers at the top
	sb += "    " // indent for row numbers
	for c := 1; c <= n; c++ {
		sb += fmt.Sprintf(" %2d ", c)
	}
	sb += "\n"

	// Separator line
	sep := "    " // indent for row numbers
	for c := 1; c <= n; c++ {
		sep += "+---"
	}
	sep += "+"

	// Each row
	for r := 1; r <= n; r++ {
		sb += sep + "\n"
		sb += fmt.Sprintf(" %2d ", r)
		for c := 1; c <= n; c++ {
			point := NewPoint(r, c)
			cell := puzzle.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				sb += "|***"
			case LetterCell:
				letter := puzzle.GetLetter(point)
				sb += fmt.Sprintf("| %s ", letter)
			}
		}
		sb += "|"
		sb += "\n"
	}

	// Bottom separator line
	sb += sep + "\n"

	return sb
}

// wordPointerIsValid returns an error if the word pointer is nil,
// or if it points to a nonexistent point in the puzzle.
func (puzzle *Puzzle) wordPointerIsValid(word *Word) error {
	if word == nil {
		errmsg := "Word number pointer is nil"
		err := errors.New(errmsg)
		return err
	}
	if err := puzzle.ValidIndex(word.point); err != nil {
		return err
	}
	return nil
}
