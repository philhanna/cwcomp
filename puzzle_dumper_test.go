package cwcomp

import (
	"testing"
)

func Test_DumpPuzzle(t *testing.T) {
	puzzle := getGoodPuzzle()
	DumpPuzzle(puzzle)
}
