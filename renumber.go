package cwcomp

import (
	"fmt"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type NumberedCell struct {
	Seq    int  // The word number (1, 2, ...)
	Row    int  // The row number (1, 2, ..., n)
	Col    int  // The column number (1, 2, ..., n)
	StartA bool // This is the start of an across word
	StartD bool // This is the start of a down word
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// GetNumberedCells determines the points in the grid that are the start
// of an across word and/or a down word.
func GetNumberedCells(cells [][]byte) []NumberedCell {
	var n = len(cells)
	var seq = 0
	ncs := make([]NumberedCell, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cells[i][j] != 0 {
				startD := (i == 0) || (cells[i-1][j] == 0)
				startA := (j == 0) || (cells[i][j-1] == 0)
				if startA || startD {
					seq++
					nc := NumberedCell{
						Seq:    seq,
						Row:    i + 1,
						Col:    j + 1,
						StartA: startA,
						StartD: startD,
					}
					ncs = append(ncs, nc)
				}

			}
		}
	}
	return ncs
}

// PuzzleToSimpleMatrix builds a simple representation of a grid as an n
// x n matrix of bytes, where '\x00' represents a black cell, and the
// rest are the letters in that cell.
func PuzzleToSimpleMatrix(puzzle *Puzzle) [][]byte {
	n := puzzle.n
	cells := make([][]byte, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			point := NewPoint(i+1, j+1)
			if puzzle.IsBlackCell(point) {
				cells[i][j] = '\x00'
			} else {
				letter := puzzle.GetLetter(point)
				cells[i][j] = letter[0]
			}
		}
	}
	return cells
}

// String returns a string representation of a numbered cell.
func (nc NumberedCell) String() string {
	return fmt.Sprintf("seq:%d,r=%d,c=%d,across=%t,down=%t",
		nc.Seq, nc.Row, nc.Col, nc.StartA, nc.StartD)
}
