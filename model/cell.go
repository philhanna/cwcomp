package model

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Cell can be either a black cell or a letter cell.
type Cell interface {
	GetPoint() Point
	String() string
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// CellIterator is a generator for all the cells in the grid, from top
// to bottom, left to right (same as PointIterator).
func (puzzle *Puzzle) CellIterator() <-chan Cell {
	out := make(chan Cell)
	go func() {
		defer close(out)
		for point := range puzzle.PointIterator() {
			cell := puzzle.GetCell(point)
			out <- cell
		}
	}()
	return out
}
