package svg

import (
	"os"
	"testing"

	"github.com/philhanna/cwcomp/model"
)

const BLK = byte('\x00')

func getGoodGrid() [][]byte {
	cells := [][]byte{
		{BLK, 'N', 'O', 'W', BLK, ' ', ' ', ' ', 'C'},
		{'B', 'L', 'U', 'E', BLK, ' ', ' ', ' ', 'O'},
		{' ', ' ', ' ', ' ', BLK, ' ', ' ', ' ', 'W'},
		{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', BLK},
		{BLK, BLK, BLK, ' ', ' ', ' ', BLK, BLK, BLK},
		{BLK, ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		{'H', ' ', ' ', ' ', BLK, ' ', ' ', ' ', ' '},
		{'O', ' ', ' ', ' ', BLK, ' ', ' ', ' ', ' '},
		{'W', ' ', ' ', ' ', BLK, ' ', ' ', ' ', BLK},
	}
	return cells
}

func TestSVG_GenerateSVG(t *testing.T) {
	cells := getGoodGrid()
	svg := NewSVG(cells)
	have := svg.GenerateSVG()
	if false { // Change this to true if you want to write the file
		os.WriteFile("testdata/simple_matrix.svg", []byte(have), 0644)
	}
}

func TestSVG_NewSVGFromGrid(t *testing.T) {
	grid := model.NewGrid(9)
	blackCells := [][]int{
		{1, 1}, {1, 5}, {2, 5}, {3, 5}, {4, 9}, {5, 1}, {5, 2}, {5, 3},
	}
	for _, pair := range blackCells {
		r, c := pair[0], pair[1]
		point := model.NewPoint(r, c)
		grid.Toggle(point)
	}
	grid.SetLetter(model.NewPoint(5, 4), "O")
	grid.SetLetter(model.NewPoint(5, 5), "A")
	grid.SetLetter(model.NewPoint(5, 6), "F")
	svg := NewSVGFromGrid(grid)
	have := svg.GenerateSVG()
	if false { // Change this to true if you want to write the file
		os.WriteFile("testdata/from_grid.svg", []byte(have), 0644)
	}
}
