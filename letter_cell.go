package cwcomp

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Letter cell is an ordinary point in the grid. It contains:
//   - The location of the cell, a Point(r, c)
//   - The character in the cell
type LetterCell struct {
	point  Point  // Location of this letter cell
	letter string // Character in the cell
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewLetterCell creates a new LetterCell at the specified location.
func NewLetterCell(point Point) LetterCell {
	p := new(LetterCell)
	p.point = point
	return *p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetPoint returns the location of this cell (for the Cell interface)
func (lc LetterCell) GetPoint() Point {
	return lc.point
}

// LetterCellIterator is a generator for all the LetterCells in the grid.
func (puzzle *Puzzle) LetterCellIterator() <-chan LetterCell {
	out := make(chan LetterCell)
	go func() {
		defer close(out)
		for point := range puzzle.PointIterator() {
			cell := puzzle.GetCell(point)
			switch typedCell := cell.(type) {
			case LetterCell:
				out <- typedCell
			}
		}
	}()
	return out
}

// String returns a string representation of this letter cell.
func (lc LetterCell) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf(`point:{%d,%d}`, lc.point.r, lc.point.c))
	parts = append(parts, fmt.Sprintf("letter:%q", lc.letter))
	s := strings.Join(parts, ",")
	return s
}
