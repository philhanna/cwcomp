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
	r int
	c int
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Compare returns -1, 0, or 1 depending on whether this point is
// less than, equal to, or greater than another.
func (p *Point) Compare(other Point) int {
	switch {
	case p.r < other.r:
		return -1
	case p.r > other.r:
		return 1
	case p.c < other.c:
		return -1
	case p.c > other.c:
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
	r := point.r
	c := point.c
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}

// ToXY converts a point (that uses 1-based coordinates) to a pair (x,
// y) that uses zero-based ones.
func (p *Point) ToXY() (int, int) {
	x := p.c - 1
	y := p.r - 1
	return x, y
}

// ValidIndex whether a point is a valid index in this grid.
func (grid *Grid) ValidIndex(point Point) error {
	r, c := point.r, point.c
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
	return fmt.Sprintf("{r:%d,c:%d}", p.r, p.c)
}