package svg

import (
	"os"
	"testing"
)

func getGoodGrid() [][]byte {
	blk := byte('\x00')
	cells := [][]byte{
		{blk, 'N', 'O', 'W', blk, ' ', ' ', ' ', 'C'},
		{'B', 'L', 'U', 'E', blk, ' ', ' ', ' ', 'O'},
		{' ', ' ', ' ', ' ', blk, ' ', ' ', ' ', 'W'},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', blk},
		{blk, blk, blk, ' ', ' ', ' ', blk, blk, blk},
		{blk, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{'H', ' ', ',', ',', blk, ' ', ' ', ' ', ' '},
		{'O', ' ', ',', ',', blk, ' ', ' ', ' ', ' '},
		{'W', ' ', ',', ',', blk, ' ', ' ', ' ', blk},
	}
	return cells
}
func TestSVG_GenerateSVG(t *testing.T) {
	cells := getGoodGrid()
	svg := NewSVG(cells)
	have := svg.GenerateSVG()
	if false { // Change this to true if you want to write the file
		os.WriteFile("test.svg", []byte(have), 0644)
	}
}
