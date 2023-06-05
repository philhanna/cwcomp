package model

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
func (ti *TestImporter) GetAcrossClues() map[int]string {
	return map[int]string{
		1: "Not then",
		8: "Not green",
	}
}
func (ti *TestImporter) GetDownClues() map[int]string {
	return map[int]string{
		7:  "Bovine",
		20: "Not why",
	}
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

func TestPuzzle_ImportPuzzle(t *testing.T) {
	importer := newTestImporter()
	puzzle, err := ImportPuzzle(importer)
	assert.Nil(t, err)
	DumpPuzzle(puzzle)
}
