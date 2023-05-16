package model

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// WordNumber is a type that exists for each numbered cell in the grid.
// It contains the word number and the location at which it exists in the
// grid
type WordNumber struct {
	seq   int   // The word number (1, 2, ...)
	point Point // The location of the head cell
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewWordNumber creates a new WordNumber structure and returns a
// pointer to it.
func NewWordNumber(seq int, point Point) *WordNumber {
	p := new(WordNumber)
	p.seq = seq
	p.point = point
	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of this word number
func (wn *WordNumber) String() string {
	parts := []string{}
	parts = append(parts, fmt.Sprintf("seq:%d", wn.seq))
	parts = append(parts, fmt.Sprintf("point:%v", wn.point.String()))
	s := strings.Join(parts, ",")
	return s
}
