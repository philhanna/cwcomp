package grid

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
	n         int                // Size of the grid (n x n square)
	cells     [][]Cell           // Black cells, Letter cells, Numbered cells
	undoStack stack.Stack[Point] // Undo stack
	redoStack stack.Stack[Point] // Redo stack
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
			point := Point{Row: i + 1, Col: j + 1}
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

// BlackCellIterator is a generator for all the black cells in the grid.
func (grid *Grid) BlackCellIterator() <-chan BlackCell {
	out := make(chan BlackCell)
	go func() {
		defer close(out)
		for point := range grid.PointIterator() {
			cell := grid.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				bc := cell.(BlackCell)
				out <- bc
			}
		}
	}()
	return out
}

// CellIterator is a generator for all the cells in the grid, from top
// to bottom, left to right (same as PointIterator).
func (grid *Grid) CellIterator() <-chan Cell {
	out := make(chan Cell)
	go func() {
		defer close(out)
		for point := range grid.PointIterator() {
			cell := grid.GetCell(point)
			out <- cell
		}
	}()
	return out
}

// GetCell returns the cell at the specified point, which may be a black
// cell, a letter cell, or a numbered cell.
func (grid *Grid) GetCell(point Point) Cell {
	x, y := point.ToXY()
	return grid.cells[y][x]
}

// LetterCellIterator is a generator for all the LetterCells in the grid.
func (grid *Grid) LetterCellIterator() <-chan LetterCell {
	out := make(chan LetterCell)
	go func() {
		defer close(out)
		for point := range grid.PointIterator() {
			cell := grid.GetCell(point)
			switch cell.(type) {
			case LetterCell:
				lc := cell.(LetterCell)
				out <- lc
			}
		}
	}()
	return out
}

// NumberedCellIterator is a generator for all the NumberedCells in the grid.
func (grid *Grid) NumberedCellIterator() <-chan NumberedCell {
	out := make(chan NumberedCell)
	go func() {
		defer close(out)
		for point := range grid.PointIterator() {
			cell := grid.GetCell(point)
			switch cell.(type) {
			case NumberedCell:
				nc := cell.(NumberedCell)
				out <- nc
			}
		}
	}()
	return out
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
		sb += fmt.Sprintf(" %2d", c)
	}
	sb += "\n"

	// Separator line
	sep := "    " // indent for row numbers
	for c := 1; c <= n; c++ {
		sep += "+--"
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
				sb += "|xx"
			case LetterCell:
				sb += "|  "
			case NumberedCell:
				nc := cell.(NumberedCell)
				sb += fmt.Sprintf("|%2d", nc.wordNumber)
			}
		}
		sb += "|"
		sb += "\n"
	}

	// Bottom separator line
	sb += sep + "\n"

	return sb
}
