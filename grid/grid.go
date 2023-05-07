package grid

import (
	"log"

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

// AddBlack cell sets the cell at the specifed point and at the
// symmetric point to be black cells.
// TODO: push this cell onto the undo stack
func (grid *Grid) AddBlackCell(point Point) {

	cell := new(BlackCell)
	symCell := new(BlackCell)

	grid.SetCell(point, *cell)
	grid.SetCell(grid.SymmetricPoint(point), *symCell)
}

// FindNumberedCells figures out the numbering of the across and down
// words in this word based on the locations of all the black cells.
//
// Algorithm:
//   - Iterate through all cells by row and column numbers (1, 2, ..., n).
//   - For each r, c:
//   - Ignore black cells.
//   - See if this is the beginning of an across word.
//     If so, find the row and column of the stopping point,
//     which is either the next black cell or the edge of the puzzle,
//     then change this to a numbered cell
//   - Do the same for down words
//   - If either an across or down word is found:
//   - Create a new WordNumber with the next available number
//   - Add it to the list
func (grid *Grid) FindNumberedCells() {

	n := grid.n

	// The numbered cell counter
	seq := 0

	// Iterate through all the cells of the grid by row, column
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {

			// Get the current point
			startingPoint := Point{r, c}

			// Ignore black cells
			if grid.IsBlackCell(startingPoint) {
				continue
			}

			// Create a structure to hold the potential new numbered
			// cell, and save its previous contents as a letter cell. We
			// do this because this grid calculation may be done when
			// some letters have already been entered in the grid
			// (EDITING_PUZZLE).
			nc := new(NumberedCell)
			nc.LetterCell = grid.GetCell(startingPoint).(LetterCell)
			nc.LetterCell.ncAcross = nil
			nc.LetterCell.ncDown = nil

			// This is the beginning point of an across word if either:
			//  - It is in the first column
			//  - The previous column on the left is a black cell
			if startingPoint.Col == 1 || grid.IsBlackCell(Point{r, c - 1}) {

				// Find the ending point of this across word
				thisPoint := startingPoint
				for thisPoint.Col < n && !grid.IsBlackCell(thisPoint) {
					thisPoint.Col++
				}

				// From this, we can calculate the length. Word lengths
				// must be at least three to qualify.
				nc.aLen = thisPoint.Col - startingPoint.Col
				if nc.aLen < 3 {
					nc.aLen = 0
				}

				if nc.aLen > 0 {

					// This is a good across word. We need to set a
					// pointer back to the starting cell in each of its
					// cells (except the starting point itself).
					thisPoint = startingPoint

					for i := 1; i < nc.aLen; i++ {
						thisPoint.Col++
						cell := grid.GetCell(thisPoint).(LetterCell)
						cell.ncAcross = &startingPoint
						grid.SetCell(thisPoint, cell)
					}
				}
			}

			// This is the beginning point of a down word if either:
			//  - It is in the first row
			//  - The previous row above is a black cell
			if startingPoint.Row == 1 || grid.IsBlackCell(Point{r - 1, c}) {

				// Find the ending point of this down word
				thisPoint := startingPoint
				for thisPoint.Row < n && !grid.IsBlackCell(thisPoint) {
					thisPoint.Row++
				}

				// From this, we can calculate the length. Word lengths
				// must be at least three to qualify.
				nc.dLen = thisPoint.Row - startingPoint.Row
				if nc.dLen < 3 {
					log.Printf("WARNING: at (%d,%d), down word is invalid\n", r, c)
					nc.dLen = 0
				}

				if nc.dLen > 0 {

					// This is a good down word. We need to set a
					// pointer back to the starting cell in each of its
					// cells (except the starting point itself).
					thisPoint = startingPoint

					for i := 1; i < nc.dLen; i++ {
						thisPoint.Row++
						cell := grid.GetCell(thisPoint).(LetterCell)
						cell.ncDown = &startingPoint
						grid.SetCell(thisPoint, cell)
					}
				}
			}

			// If either across or down words exist (length > 0) then
			// treat this starting point as a numbered cell.
			if nc.aLen > 0 || nc.dLen > 0 {

				// Assign the word number and update the cell in the
				// grid to be a numbered cell.
				seq++
				nc.seq = seq
				grid.SetCell(startingPoint, *nc)
			}
		}
	}
}

// GetCell returns the cell at the specified point, which may be a black
// cell, a letter cell, or a numbered cell.
func (grid *Grid) GetCell(point Point) Cell {
	x := point.Col - 1
	y := point.Row - 1
	return grid.cells[y][x]
}

// IsBlackCell returns true if the specified point is a black cell.
func (g *Grid) IsBlackCell(point Point) bool {
	cell := g.GetCell(point)
	switch cell.(type) {
	case BlackCell:
		return true
	default:
		return false
	}
}

// SetCell sets the cell at the specified point
func (grid *Grid) SetCell(point Point, cell Cell) {
	x := point.Col - 1
	y := point.Row - 1
	grid.cells[y][x] = cell
}

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.Row
	c := point.Col
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}
