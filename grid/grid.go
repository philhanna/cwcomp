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
	n         int                // Size of the grid (n x n square)
	bclist    []Point            // Black cell points
	nclist    []NumberedCell     // Numbered cells
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
	g.bclist = make([]Point, 0)
	g.nclist = make([]NumberedCell, 0)
	g.undoStack = stack.NewStack[Point]()
	g.redoStack = stack.NewStack[Point]()
	return g
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------
// FindNumberedCells figures out the numbering of the across and down
// words in this word based on the locations of all the black cells.
//
// Algorithm:
//   - Iterate through all cells by row and column numbers (1, 2, ... n).
//   - For each r, c:
//   - Ignore black cells.
//   - See if this is the beginning of an across word.
//     If so, find the row and column of the stopping point,
//     which is either the next black cell or the edge of the puzzle.
//   - Do the same for down words
//   - If either an across or down word is found:
//   - Create a new WordNumber with the next available number
//   - Add it to the list
func (grid *Grid) FindNumberedCells() {
	var n = grid.n

	nclist := make([]NumberedCell, 0)
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {

			// Ignore black cells

			if grid.IsBlackCell(Point{r, c}) {
				continue
			}

			// If this is the beginning of an "across" word, find the
			// row and column of the stopping point, which is either the
			// next black cell, or the edge of the puzzle.

			acrossLength := 0
			if c == 1 || grid.IsBlackCell(Point{r, c - 1}) {
				point := Point{r, c}
				for point.Col < n && !grid.IsBlackCell(point) {
					acrossLength++
					point.Col++
				}

				// Words must have length >= 3. Set length to zero if
				// this is not the case, so that a new word number will
				// not be created for it, unless there is also a down
				// word.

				if acrossLength < 3 {
					acrossLength = 0
				}
			}

			// If this is the beginning of a "down" word, find the
			// row and column of the stopping point, which is either the
			// next black cell, or the edge of the puzzle.

			downLength := 0
			if r == 1 || grid.IsBlackCell(Point{r - 1, c}) {
				point := Point{r, c}
				for point.Row < n && !grid.IsBlackCell(point) {
					downLength++
					point.Row++
				}

				// Words must have length >= 3. Set length to zero if
				// this is not the case, so that a new word number will
				// not be created for it, unless there is also a down
				// word.

				if downLength < 3 {
					downLength = 0
				}
			}

			// If either an across or down word was found, create a new
			// WordNumber with the next available number and add it to
			// the list.

			if acrossLength > 0 || downLength > 0 {
				number := 1 + len(nclist)
				nc := NumberedCell{Point{r, c}, number, acrossLength, downLength}
				nclist = append(nclist, nc)
			}
		}
	}
	grid.nclist = nclist
}

// IndexOfBlackCell returns the integer index of a given point in the
// black cell array, or -1, if not found.
func (g *Grid) IndexOfBlackCell(point Point) int {
	for i, bc := range g.bclist {
		if bc == point {
			return i
		}
	}
	return -1
}

// IsBlackCell returns true if the specified point is a black cell.
func (g *Grid) IsBlackCell(point Point) bool {
	x := g.IndexOfBlackCell(point)
	return x != -1
}
