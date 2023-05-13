package model

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Letter cell is an ordinary point in the grid. It contains:
//   - A pointer to the numbered cell for the across word (if any)
//   - A pointer to the numbered cell for the down word (if any)
//   - The character in the cell
type LetterCell struct {
	point    Point       // Location of this letter cell
	ncAcross *WordNumber // Pointer to the numbered cell in the across direction
	ncDown   *WordNumber // Pointer to the numbered cell in the down direction
	letter   string      // Character in the cell
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

// String returns a string representation of this letter cell.
func (lc LetterCell) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf(`point:{%d,%d}`, lc.point.Row, lc.point.Col))
	parts = append(parts, fmt.Sprintf(`ncAcross:%v`, lc.ncAcross))
	parts = append(parts, fmt.Sprintf(`ncDown:%v`, lc.ncDown))
	parts = append(parts, fmt.Sprintf("letter:%q", lc.letter))
	s := strings.Join(parts, ",")
	return s
}
