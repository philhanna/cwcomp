package grid

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
	point    Point  // Location of this letter cell
	ncAcross *Point // Pointer to the numbered cell in the across direction
	ncDown   *Point // Pointer to the numbered cell in the down direction
	letter   string // Character in the cell
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

// LetterCell returns a string representation of this letter cell.
func (lc *LetterCell) String() string {
	parts := make([]string, 0)
	parts = append(parts, fmt.Sprintf(`point:{%d,%d}`, lc.point.Row, lc.point.Col))
	if lc.ncAcross == nil {
		parts = append(parts, fmt.Sprintf(`ncAcross:%v`, nil))
	} else {
		parts = append(parts, fmt.Sprintf(`ncAcross:%v`, lc.ncAcross.String()))
	}
	if lc.ncDown == nil {
		parts = append(parts, fmt.Sprintf(`ncDown:%v`, nil))
	} else {
		parts = append(parts, fmt.Sprintf(`ncDown:%v`, lc.ncDown.String()))
	}
	parts = append(parts, fmt.Sprintf("letter:%q", lc.letter))
	s := strings.Join(parts, ",")
	return s
}
