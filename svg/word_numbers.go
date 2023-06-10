package svg

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type WordNumber struct {
	Seq int
	Row int // 1, 2, ..., n
	Col int // 1, 2, ..., n
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// WordNumbers generates the word numbers of the grid.
func (svg *SVG) WordNumbers() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Word numbers -->\n")
	ncs := svg.GetNumberedCells()
	for _, nc := range ncs {
		seq := nc.Seq
		xbase := (nc.Col-1)*BOXSIZE + NUMBER_X_OFFSET
		ybase := (nc.Row-1)*BOXSIZE + NUMBER_Y_OFFSET
		fmtstr := `<text x="%d" y="%d" font-size="%s">%d</text>`
		line := fmt.Sprintf(fmtstr, xbase, ybase, NUMBER_FONT_SIZE, seq)
		sb.WriteString(line + "\n")
	}
	return sb.String()
}

// GetNumberedCells determines the points in the grid that are the start
// of an across word and/or a down word.
func (svg *SVG) GetNumberedCells() []WordNumber {
	cells := svg.cells
	var n = len(cells)
	var seq = 0
	ncs := make([]WordNumber, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cells[i][j] != BLACK_CELL {
				startD := (i == 0) || (cells[i-1][j] == BLACK_CELL)
				startA := (j == 0) || (cells[i][j-1] == BLACK_CELL)
				if startA || startD {
					seq++
					nc := WordNumber{
						Seq: seq,
						Row: i + 1,
						Col: j + 1,
					}
					ncs = append(ncs, nc)
				}
			}
		}
	}
	return ncs
}
