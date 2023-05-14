package model

import (
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
	n             int                 // Size of the grid (n x n square)
	cells         [][]Cell            // Black cells and letter cells
	wordNumberMap map[int]*WordNumber // Word number pointers
	undoStack     stack.Stack[Point]  // Undo stack
	redoStack     stack.Stack[Point]  // Redo stack
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
			point := Point{r: i + 1, c: j + 1}
			cell := NewLetterCell(point)
			g.cells[i][j] = cell
		}
	}

	g.wordNumberMap = make(map[int]*WordNumber)
	g.undoStack = stack.NewStack[Point]()
	g.redoStack = stack.NewStack[Point]()
	return g
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetAcrossWordLength returns the length of the across word for this
// numbered cell.
func (grid *Grid) GetAcrossWordLength(pnc *Point) int {
	n := grid.n
	length := 0
	point := Point{pnc.r, pnc.c}
	for point.c <= n && !grid.IsBlackCell(point) {
		length++
		point.c++
	}
	return length
}

// GetDownWordLength returns the length of the down word for this
// numbered cell.
func (grid *Grid) GetDownWordLength(pnc *Point) int {
	n := grid.n
	length := 0
	point := Point{pnc.r, pnc.c}
	for point.r <= n && !grid.IsBlackCell(point) {
		length++
		point.r++
	}
	return length

}

// GetCell returns the cell at the specified point, which may be a black
// cell, a letter cell, or a numbered cell.
func (grid *Grid) GetCell(point Point) Cell {
	x, y := point.ToXY()
	return grid.cells[y][x]
}

// PointIterator is a generator for all the points in the grid, from
// top bottom and left to right (i.e, (1, 1), (1, 2), ..., (1, n),
// (2, 1), (2, 2), ..., (2, n), ..., (n, 1) (n, 2), ..., (n, n)).
func (grid *Grid) PointIterator() <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		n := grid.n
		for r := 1; r <= n; r++ {
			for c := 1; c <= n; c++ {
				out <- Point{r, c}
			}
		}
	}()
	return out
}

// RenumberCells assigns the word numbers based on the locations of the
// black cells.
func (grid *Grid) RenumberCells() {
	seq := 0 // Next available word number

	// Reset the list to empty
	grid.wordNumberMap = make(map[int]*WordNumber)

	// Look through all the letter cells
	for lc := range grid.LetterCellIterator() {

		point := lc.GetPoint()

		var wn *WordNumber
		var aStart bool
		var dStart bool

		// Determine if this cell is the beginning of an across or a
		// down word, setting a boolean variable for either case.
		aStart = point.c == 1 || grid.IsBlackCell(Point{point.r, point.c - 1})
		dStart = point.r == 1 || grid.IsBlackCell(Point{point.r - 1, point.c})

		// If not a new word, skip to the next cell
		if !aStart && !dStart {
			continue
		}

		// If either is true, create a new WordNumber
		if aStart || dStart {
			seq++
			wn = NewWordNumber(seq, point, 0, 0)
		}

		// Then if this is the start of an across word, calculate the
		// length and store that in the WordNumber
		if aStart {
			wn.aLen = 0
			aPoint := Point{point.r, point.c}
			for aPoint.c <= grid.n && !grid.IsBlackCell(aPoint) {
				cell := grid.GetCell(aPoint).(LetterCell)
				cell.ncAcross = wn
				grid.SetCell(aPoint, cell)
				wn.aLen++
				aPoint.c++
			}
		}

		// Or if this is the start of a down word, calculate the
		// length and store that in the WordNumber
		if dStart {
			wn.dLen = 0
			dPoint := Point{point.r, point.c}
			for dPoint.r <= grid.n && !grid.IsBlackCell(dPoint) {
				cell := grid.GetCell(dPoint).(LetterCell)
				cell.ncDown = wn
				grid.SetCell(dPoint, cell)
				wn.dLen++
				dPoint.r++
			}
		}

		// Store the new word number in the word number map
		grid.wordNumberMap[seq] = wn
	}
}

// SetCell sets the cell at the specified point
func (grid *Grid) SetCell(point Point, cell Cell) {
	x, y := point.ToXY()
	grid.cells[y][x] = cell
}

// String returns a string representation of the grid
func (grid *Grid) String() string {
	n := grid.n

	// Row of column numbers at the top
	sb := "    " // indent for row numbers
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
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				sb += "|***"
			case LetterCell:
				sb += "|   "
			}
		}
		sb += "|"
		sb += "\n"
	}

	// Bottom separator line
	sb += sep + "\n"

	return sb
}
