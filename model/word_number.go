package model

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// WordNumber is a type that exists for each numbered cell in the grid.
// It contains the word number, the location at which it exists in the
// grid, and the lengths of its across and/or down words, either of
// which can be zero.
type WordNumber struct {
	seq   int    // The word number (1, 2, ...)
	point Point  // The location of the head cell
	aLen  int    // Length of the across word (0 = no across word)
	dLen  int    // Length of the down word (0 = no down word)
	aClue string // Clue for the across word
	dClue string // Clue for the down word
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewWordNumber creates a new WordNumber structure and returns a
// pointer to it.
func NewWordNumber(seq int, point Point, aLen int, dLen int) *WordNumber {
	p := new(WordNumber)
	p.seq = seq
	p.point = point
	p.aLen = aLen
	p.dLen = dLen
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
	parts = append(parts, fmt.Sprintf("aLen:%d", wn.aLen))
	parts = append(parts, fmt.Sprintf("dLen:%d", wn.dLen))
	parts = append(parts, fmt.Sprintf("aClue:%q", wn.aClue))
	parts = append(parts, fmt.Sprintf("dClue:%q", wn.dClue))
	s := strings.Join(parts, ",")
	return s
}
