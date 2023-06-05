package model

import (
	"testing"
)

func Test_DumpPuzzle(t *testing.T) {
	puzzle := getGoodPuzzle()
	word := puzzle.LookupWordByNumber(1, ACROSS)
	puzzle.SetClue(word, "The meaning is 42")
	DumpPuzzle(puzzle)
}

func Test_TracePuzzle(t *testing.T) {
	puzzle := getGoodPuzzle()
	TracePuzzle(puzzle)
}
