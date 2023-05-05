package grid

import "fmt"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Point is a row and column pair
type Point struct {
	Row int `json:"r"`
	Col int `json:"c"`
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Compare returns -1, 0, or 1 depending on whether this point is
// less than, equal to, or greater than another.
func (p *Point) Compare(other Point) int {
	switch {
	case p.Row < other.Row:
		return -1
	case p.Row > other.Row:
		return 1
	case p.Col < other.Col:
		return -1
	case p.Col > other.Col:
		return 1
	default:
		return 0
	}
}

// Equal is true if this point has the same row and column of another
// point.
func (p *Point) Equal(other Point) bool {
	return *p == other
}

// String returns a string representation of this type
func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.Row, p.Col)
}
