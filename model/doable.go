package model

import "fmt"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Doable is an entry in the undoWord/redoWord stacks
type Doable struct {
	word Word
	text string
}

// NewDoable creates a Doable from the specified word in the grid
func NewDoable(puzzle *Puzzle, word *Word) Doable {
	p := new(Doable)
	p.word = *word
	p.text = puzzle.GetText(word)
	return *p
}

// String returns a string representation of the Doable
func (d Doable) String() string {
	return fmt.Sprintf("%s,text=%q", d.word.String(), d.text)
}
