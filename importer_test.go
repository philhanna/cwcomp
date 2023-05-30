package cwcomp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestImporter struct {
	cells [][]byte
}

func newTestImporter() *TestImporter {
	const blk = BLACK_CELL
	ti := new(TestImporter)
	ti.cells = [][]byte{
		{blk, 'N', 'O', 'W', blk, ' ', ' ', ' ', 'C'},
		{'B', 'L', 'U', 'E', blk, ' ', ' ', ' ', 'O'},
		{' ', ' ', ' ', ' ', blk, ' ', ' ', ' ', 'W'},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', blk},
		{blk, blk, blk, ' ', ' ', ' ', blk, blk, blk},
		{blk, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{'H', ' ', ' ', ' ', blk, ' ', ' ', ' ', ' '},
		{'O', ' ', ' ', ' ', blk, ' ', ' ', ' ', ' '},
		{'W', ' ', ' ', ' ', blk, ' ', ' ', ' ', blk}}

	return ti
}
func (ti *TestImporter) GetSize() int     { return 9 }
func (ti *TestImporter) GetName() string  { return "good9" }
func (ti *TestImporter) GetTitle() string { return "what does this mean?" }
func (ti *TestImporter) GetCell(r, c int) (byte, error) {
	n := ti.GetSize()
	if r < 1 || r > n || c < 1 || c > n {
		return 0, fmt.Errorf("Invalid index: r=%d,c=%d", r, c)
	}
	i, j := r-1, c-1
	return ti.cells[i][j], nil
}
func (ti *TestImporter) GetAcrossClues() []string {
	return []string{
		"Not then",  // 1 across
		"",          // 2 across
		"",          // 3 across
		"",          // 4 across
		"",          // 5 across
		"",          // 6 across
		"",          // 7 across
		"Not green", // 8 across
		"",          // 9 across
		"",          // 10 across
		"",          // 11 across
		"",          // 12 across
		"",          // 13 across
		"",          // 14 across
		"",          // 15 across
		"",          // 16 across
		"",          // 17 across
		"",          // 18 across
		"",          // 19 across
		"",          // 20 across
		"",          // 21 across
		"",          // 22 across
		"",          // 23 across
		"",          // 24 across
		"",          // 25 across
	}
}
func (ti *TestImporter) GetDownClues() []string {
	return []string{
		"", // 1 down
		"", // 2 down
		"", // 3 down
		"", // 4 down
		"", // 5 down
		"", // 6 down
		"Bovine", // 7 down
		"", // 8 down
		"", // 9 down
		"", // 10 down
		"", // 11 down
		"", // 12 down
		"", // 13 down
		"", // 14 down
		"", // 15 down
		"", // 16 down
		"", // 17 down
		"", // 18 down
		"", // 19 down
		"Not why", // 20 down
		"", // 21 down
		"", // 22 down
		"", // 23 down
		"", // 24 down
		"", // 25 down
	}
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

func TestPuzzle_ImportPuzzle(t *testing.T) {
	importer := newTestImporter()
	puzzle, err := ImportPuzzle(importer)
	assert.Nil(t, err)
	fmt.Println(puzzle.String())
}
