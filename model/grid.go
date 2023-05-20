package model

import (
	"errors"
	"fmt"

	"github.com/philhanna/stack"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Grid contains the cells of a puzzle.
//
// A grid is constructed with the single parameter n, which is the size
// (n x n) of the grid.
//
// Any of the cells in the grid can be "black cells", which act as the
// boundaries of where the words can go. The model automatically takes
// care of matching a black cell with its symmetric twin 180 degrees
// from it.
//
// Wherever an across or down word starts, the grid assigns the next
// available word number to the cell and keeps track of the lengths of
// the across and down words.
//
// Grid supports a full "undo/redo" capability for the current session
// (from load to save).  Any black cell additions or deletions are
// pushed on an undo stack.
type Grid struct {
	n           int                // Size of the grid (n x n square)
	gridName    string             // The grid name
	cells       [][]Cell           // Black cells and letter cells
	words       []*Word            // Pointers to the words in this grid
	wordNumbers []*WordNumber      // Word number pointers
	undoStack   stack.Stack[Point] // Undo stack
	redoStack   stack.Stack[Point] // Redo stack
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewGrid creates a grid of the specified size.
func NewGrid(n int) *Grid {
	g := new(Grid)
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

	g.undoStack = stack.NewStack[Point]()
	g.redoStack = stack.NewStack[Point]()

	return g
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------
// GetCell returns the cell at the specified point, which may be a black
// cell, a letter cell, or a numbered cell.
func (grid *Grid) GetCell(point Point) Cell {
	err := grid.ValidIndex(point)
	if err != nil {
		errmsg := fmt.Sprintf("%s is not a valid point", point.String())
		panic(errmsg)
	}
	x, y := point.ToXY()
	return grid.cells[y][x]
}

// GetClue returns the clue for the word.
func (grid *Grid) GetClue(word *Word) (string, error) {
	err := grid.wordPointerIsValid(word)
	if err != nil {
		return "", err
	}
	return word.clue, nil
}

// GetGridName returns the grid name
func (grid *Grid) GetGridName() string {
	return grid.gridName
}

// GetLength returns the length of the word.
func (grid *Grid) GetLength(word *Word) (int, error) {
	err := grid.wordPointerIsValid(word)
	if err != nil {
		return 0, err
	}
	return word.length, nil
}

// GetLetter returns the value of the cell at this point.  The length of
// the returned value is always 1, unless the point refers to a black
// cell, in which case the length is zero.
func (grid *Grid) GetLetter(point Point) string {
	letter := ""
	cell := grid.GetCell(point)
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
func (grid *Grid) GetText(word *Word) (string, error) {
	err := grid.wordPointerIsValid(word)
	if err != nil {
		return "", err
	}

	length, _ := grid.GetLength(word)

	var s string
	var point Point
	for i := 0; i < length; i++ {
		switch word.direction {
		case ACROSS:
			point = NewPoint(word.point.r, word.point.c+i)
		case DOWN:
			point = NewPoint(word.point.r+i, word.point.c)
		}
		letter := grid.GetLetter(point)
		s += letter
	}
	return s, nil
}

// LookupWord returns the word at this point and direction
func (grid *Grid) LookupWord(point Point, dir Direction) *Word {
	for _, word := range grid.words {
		if word.point == point && word.direction == dir {
			return word
		}
	}
	return nil
}

// LookupWordByNumber returns the word at this point and direction
func (grid *Grid) LookupWordByNumber(seq int, dir Direction) *Word {
	wn := grid.LookupWordNumber(seq)
	if wn == nil {
		return nil
	}
	return grid.LookupWord(wn.point, dir)
}

// LookupWordNumber returns the WordNumber for this number
func (grid *Grid) LookupWordNumber(seq int) *WordNumber {
	for _, wn := range grid.wordNumbers {
		if wn.seq == seq {
			return wn
		}
	}
	return nil
}

// LookupWordNumberByPoint returns the WordNumber starting at this point.
func (grid *Grid) LookupWordNumberByPoint(point Point) *WordNumber {
	for _, wn := range grid.wordNumbers {
		if wn.point == point {
			return wn
		}
	}
	return nil
}

// RenumberCells assigns the word numbers based on the locations of the
// black cells.
func (grid *Grid) RenumberCells() {

	var (
		seq    int         = 0 // Next available word number
		wn     *WordNumber = nil
		aStart bool        = false
		dStart bool        = false
	)

	// Reset the list to empty
	// TODO save clues if word already existed
	grid.wordNumbers = make([]*WordNumber, 0)
	grid.words = make([]*Word, 0)

	// Look through all the letter cells
	for point := range grid.PointIterator() {

		// Skip black cells
		if grid.IsBlackCell(point) {
			continue
		}

		// Determine if this cell is the beginning of an across or a
		// down word, setting a boolean variable for either case.

		aStart = point.c == 1 || grid.IsBlackCell(NewPoint(point.r, point.c-1))
		dStart = point.r == 1 || grid.IsBlackCell(NewPoint(point.r-1, point.c))

		// If either is true, create a new WordNumber
		if aStart || dStart {
			seq++
			wn = NewWordNumber(seq, point)
			grid.wordNumbers = append(grid.wordNumbers, wn)
		}
	}

	// Now calculate the word lengths

	for _, wn := range grid.wordNumbers {

		// Determine if this cell is the beginning of an across or a
		// down word, setting a boolean variable for either case.
		point := wn.point
		aStart := point.c == 1 || grid.IsBlackCell(NewPoint(point.r, point.c-1))
		dStart := point.r == 1 || grid.IsBlackCell(NewPoint(point.r-1, point.c))

		if aStart {
			word := NewWord(point, ACROSS, 0, "")
			for range grid.WordIterator(point, ACROSS) {
				word.length++
			}
			grid.words = append(grid.words, word)
		}
		if dStart {
			word := NewWord(point, DOWN, 0, "")
			for range grid.WordIterator(point, DOWN) {
				word.length++
			}
			grid.words = append(grid.words, word)
		}
	}
}

// SetCell sets the cell at the specified point
func (grid *Grid) SetCell(point Point, cell Cell) {
	x, y := point.ToXY()
	grid.cells[y][x] = cell
}

// SetLetter sets the letter value of the cell at the specified point
func (grid *Grid) SetLetter(point Point, letter string) {
	cell := grid.GetCell(point)
	switch typedCell := cell.(type) {
	case LetterCell:
		typedCell.letter = letter
		grid.SetCell(point, typedCell)
	}
}

// SetGridName sets the grid title
func (grid *Grid) SetGridName(name string) {
	grid.gridName = name
}

// SetText sets the text in the grid for a specified word.
func (grid *Grid) SetText(word *Word, text string) error {

	// Make sure this is a valid word pointer
	err := grid.wordPointerIsValid(word)
	if err != nil {
		return err
	}

	// Make sure the text is not longer than the word allows
	if len(text) > word.length {
		errmsg := fmt.Sprintf(`Text %q length > expected %d`, len(text), word.length)
		err := errors.New(errmsg)
		return err
	}

	// Pad the text with blanks if too short
	for len(text) < word.length {
		text += " "
	}

	// Iterate through the points of the word, storing the text into it
	// letter by letter.
	i := 0
	for point := range grid.WordIterator(word.point, word.direction) {
		ch := text[i]
		i++
		letter := string(ch)
		grid.SetLetter(point, letter)
	}

	// OK
	return nil
}

// String returns a string representation of the grid
func (grid *Grid) String() string {
	n := grid.n
	sb := ""
	if grid.gridName == "" {
		sb += "(Untitled)"
	} else {
		sb += grid.gridName
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
			cell := grid.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				sb += "|***"
			case LetterCell:
				letter := grid.GetLetter(point)
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
// or if it points to a nonexistent point in the grid.
func (grid *Grid) wordPointerIsValid(word *Word) error {
	if word == nil {
		errmsg := "Word number pointer is nil"
		err := errors.New(errmsg)
		return err
	}
	if err := grid.ValidIndex(word.point); err != nil {
		return err
	}
	return nil
}
