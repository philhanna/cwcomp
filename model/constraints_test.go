package model

import (
	"fmt"
	"testing"
)

func TestPuzzle_GetConstraints(t *testing.T) {

	var (
		word        *Word
		constraints []*Constraint
		overall     string
	)

	puzzle := getGoodPuzzle()

	// Set the text of some words
	someWords := []struct {
		seq  int
		dir  Direction
		text string
	}{
		{1, ACROSS, "NOW"},
		{8, ACROSS, "BLUE"},
		{10, ACROSS, "RA  "},
	}
	for _, sw := range someWords {
		word = puzzle.LookupWordByNumber(sw.seq, sw.dir)
		puzzle.SetText(word, sw.text)
	}

	// Now find the constraints of a word in that puzzle
	word = puzzle.LookupWordByNumber(2, DOWN)
	constraints = puzzle.GetConstraints(word)
	fmt.Printf("Constraints for 2 down:\n")
	overall = ""
	for i, constraint := range constraints {
		fmt.Printf("%d: %v\n", i, constraint.ToJSON())
		overall += constraint.Pattern
	}
	fmt.Printf("Overall pattern: %q\n", overall)

	// Try 10 across
	word = puzzle.LookupWordByNumber(10, ACROSS)
	constraints = puzzle.GetConstraints(word)
	fmt.Printf("Constraints for 10 across:\n")
	overall = ""
	for i, constraint := range constraints {
		fmt.Printf("%d: %v\n", i, constraint.ToJSON())
		overall += constraint.Pattern
	}
	fmt.Printf("Overall pattern: %q\n", overall)

}
