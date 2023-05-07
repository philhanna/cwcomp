package grid

import (
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
			cell := new(LetterCell)
			g.cells[i][j] = *cell
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
	x, y := point.ToXY()
	return grid.cells[y][x]
}

// SetCell sets the cell at the specified point
func (grid *Grid) SetCell(point Point, cell Cell) {
	x, y := point.ToXY()
	grid.cells[y][x] = cell
}
