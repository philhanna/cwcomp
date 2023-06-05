package model

import (
	"testing"
)

func Test_DumpPuzzle(t *testing.T) {
	puzzle := getGoodPuzzle()
	DumpPuzzle(puzzle)
}

func Test_TracePuzzle(t *testing.T) {
	puzzle := getGoodPuzzle()
	TracePuzzle(puzzle)
}
