package grid

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// NumberedCell is a letter cell that is the beginning of an across word
// and/or a down word.
type NumberedCell struct {
	LetterCell     // The numbered cell letter values
	seq        int // The word number
	aLen       int // Length of the across word (if any)
	dLen       int // Length of the down word (if any)
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of the structure.
func (nc *NumberedCell) String() string {
	return strings.Join([]string{
		fmt.Sprintf("letterCell:%v", nc.LetterCell),
		fmt.Sprintf("seq:%d", nc.seq),
		fmt.Sprintf("aLen:%d", nc.aLen),
		fmt.Sprintf("dLen:%d", nc.dLen),
	}, ", ")
}