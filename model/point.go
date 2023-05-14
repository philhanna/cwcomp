package model

import (
	"errors"
	"fmt"
)

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

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.Row
	c := point.Col
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}

// ToXY converts a point (that uses 1-based coordinates) to a pair (x,
// y) that uses zero-based ones.
func (p *Point) ToXY() (int, int) {
	x := p.Col - 1
	y := p.Row - 1
	return x, y
}

// ValidIndex whether a point is a valid index in this grid.
func (grid *Grid) ValidIndex(point Point) error {
	r, c := point.Row, point.Col
	validRow := r >= 1 && r <= grid.n
	validCol := c >= 1 && c <= grid.n
	if validRow && validCol {
		return nil
	}

	var errmsg string

	if !validRow && !validCol {
		errmsg = fmt.Sprintf("Invalid row %d and column %d\n", r, c)
	} else if !validRow {
		errmsg = fmt.Sprintf("Invalid row %d\n", r)
	} else if !validCol {
		errmsg = fmt.Sprintf("Invalid column %d\n", c)
	}

	err := errors.New(errmsg)
	return err
}

// String returns a string representation of this type
func (p *Point) String() string {
	return fmt.Sprintf("{Row:%d,Col:%d}", p.Row, p.Col)
}
