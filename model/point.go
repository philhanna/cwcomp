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
// Constructor
// ---------------------------------------------------------------------

// Creates a new point with the supplied row and column.
func NewPoint(r, c int) Point {
	return Point{r, c}
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

// PointIterator is a generator for all the points in the grid, from
// top bottom and left to right (i.e, (1, 1), (1, 2), ..., (1, n),
// (2, 1), (2, 2), ..., (2, n), ..., (n, 1) (n, 2), ..., (n, n)).
func (grid *Grid) PointIterator() <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		n := grid.n
		for r := 1; r <= n; r++ {
			for c := 1; c <= n; c++ {
				out <- NewPoint(r, c)
			}
		}
	}()
	return out
}

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.r
	c := point.c
	return NewPoint(grid.n + 1 - r, grid.n + 1 - c)
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
