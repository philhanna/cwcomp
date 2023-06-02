package svg

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/philhanna/cwcomp"
)

const BLK = byte('\x00')

func getMatrix() [][]byte {
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
	cells := getMatrix()
	svg := NewSVG(cells)
	have := svg.GenerateSVG()
	filename := filepath.Join(getOutputDirectory(), "simple_matrix.svg")
	os.WriteFile(filename, []byte(have), 0644)
}

func getOutputDirectory() string {
	output := filepath.Join(os.TempDir(), "cwcomp")
	os.MkdirAll(output, 0750)
	return output
}

func TestSVG_NewSVGFromPuzzle(t *testing.T) {
	puzzle := cwcomp.NewPuzzle(9)
	blackCells := [][]int{
		{1, 1}, {1, 5}, {2, 5}, {3, 5}, {4, 9}, {5, 1}, {5, 2}, {5, 3},
	}
	for _, pair := range blackCells {
		r, c := pair[0], pair[1]
		point := cwcomp.NewPoint(r, c)
		puzzle.Toggle(point)
	}
	puzzle.SetLetter(cwcomp.NewPoint(5, 4), "O")
	puzzle.SetLetter(cwcomp.NewPoint(5, 5), "A")
	puzzle.SetLetter(cwcomp.NewPoint(5, 6), "F")
	svg := NewSVGFromPuzzle(puzzle)
	have := svg.GenerateSVG()
	filename := filepath.Join(getOutputDirectory(), "from_puzzle.svg")
	os.WriteFile(filename, []byte(have), 0644)

}
