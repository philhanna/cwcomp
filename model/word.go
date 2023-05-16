package model

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Word consists of a point and a direction
type Word struct {
	point     Point     // Point at which the word starts
	direction Direction // Across or down
	length    int       // Length of the word
	clue      string    // The text of the clue
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// NewWord is the constructor for Word
func NewWord(point Point, dir Direction, length int, clue string) *Word {
	p := new(Word)
	p.point = point
	p.direction = dir
	p.length = length
	p.clue = clue
	return p
}

// String returns a string representation of this object.
func (word *Word) String() string {
	parts := []string{
		fmt.Sprintf("point:%s", word.point.String()),
		fmt.Sprintf("direction:%q", word.direction.String()),
		fmt.Sprintf("length:%d", word.length),
		fmt.Sprintf("clue:%q", word.clue),
	}
	return strings.Join(parts, ",")
}

// WordIterator iterates through the points in a word, stopping when it
// encounters a black cell or the edge of the grid.
func (grid *Grid) WordIterator(point Point, dir Direction) <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		pp := NewPoint(point.r, point.c)
		for {
			if err := grid.ValidIndex(pp); err != nil {
				return
			}
			if grid.IsBlackCell(pp) {
				return
			}
			out <- pp
			switch dir {
			case ACROSS:
				pp.c++
			case DOWN:
				pp.r++
			}
		}
	}()
	return out
}
