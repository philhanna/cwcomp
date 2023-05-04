package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getGoodGrid() *Grid {
	p := NewGrid(9)
	p.bclist = []Point{
		{1, 1}, {1, 5},
		{2, 5},
		{3, 5},
		{4, 9},
		{5, 1}, {5, 2}, {5, 3},
		{5, 7}, {5, 8}, {5, 9},
		{6, 1},
		{7, 5},
		{8, 5},
		{9, 5}, {9, 9},
	}
	p.FindNumberedCells()
	return p
}

func TestGrid_CalculateWordNumbers(t *testing.T) {
	grid := getGoodGrid()
	assert.Equal(t, 25, len(grid.nclist))
}
