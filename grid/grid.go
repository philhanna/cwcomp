package grid

import "github.com/philhanna/stack"

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
	N           int                 `json:"n"`   // Size of the grid (n x n square)
	BCS         []Point             `json:"bcs"` // Black cell points
	wordNumbers []Point             // Numbered cells (word number is index + 1)
	undoStack   stack.Stack[Doable] // Undo stack
	redoStack   stack.Stack[Doable] // Redo stack
}

// A record of an event that can be undone and redone
type Doable interface {
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewGrid creates a grid of the specified size.
func NewGrid(n int) *Grid {
	g := new(Grid)
	g.N = n
	g.BCS = make([]Point, 0)
	g.wordNumbers = make([]Point, 0)
	g.undoStack = stack.NewStack[Doable]()
	g.redoStack = stack.NewStack[Doable]()
	return g
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------
