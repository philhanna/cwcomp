package svg

import (
	"log"
	"os"
	"path/filepath"
	"testing"
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
	log.Printf("SVG output written to %v\n", filename)
}

func getOutputDirectory() string {
	output := filepath.Join(os.TempDir(), "cwcomp")
	os.MkdirAll(output, 0750)
	return output
}
