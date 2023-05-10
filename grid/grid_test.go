package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getGoodGrid() *Grid {
	points := []Point{
		{1, 1}, {1, 5},
		{2, 5},
		{3, 5},
		{4, 9},
		{5, 1}, {5, 2}, {5, 3},
	}
	return getTestGrid(points)
}

func getTestGrid(points []Point) *Grid {
	grid := NewGrid(9)
	for _, point := range points {
		grid.AddBlackCell(point)
	}
	grid.RenumberCells()
	return grid
}

func TestGrid_BlackCellIterator(t *testing.T) {
	points := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
	}
	grid := NewGrid(9)
	for _, point := range points {
		grid.AddBlackCell(point)
	}

	expected := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
		{7, 5},
		{9, 9},
	}
	actual := []Point{}
	for bc := range grid.BlackCellIterator() {
		actual = append(actual, bc.point)
	}
	assert.Equal(t, expected, actual)
}

func TestGrid_GetAcrossWordLength(t *testing.T) {
	grid := getGoodGrid()
	point := Point{1, 2}
	nc := grid.GetCell(point).(NumberedCell)
	assert.Equal(t, 3, grid.GetAcrossWordLength(&nc.point))
}

func TestGrid_GetDownWordLength(t *testing.T) {
}

func TestGrid_LetterCellIterator(t *testing.T) {
	grid := getGoodGrid()
	nlc := 0
	for range grid.LetterCellIterator() {
		nlc++
	}
	assert.Equal(t, 9*9-25-16, nlc)
}

func TestGrid_NumberCellIterator(t *testing.T) {
	grid := getGoodGrid()
	nnc := 0
	for range grid.NumberedCellIterator() {
		nnc++
	}
	assert.Equal(t, 25, nnc)
}

func TestGrid_PointIterator(t *testing.T) {
	const n = 3
	grid := NewGrid(n)

	// Make a list of all points received from the iterator
	list1 := make([]Point, n*n)
	it := grid.PointIterator()
	index := 0
	for point := range it {
		list1[index] = point
		index++
	}

	// Make another list of points created from nested loops
	list2 := make([]Point, n*n)
	index = 0
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			list2[index] = Point{r, c}
			index++
		}
	}

	// Should be the same
	assert.Equal(t, list1, list2)
}

func TestGrid_SymmetricPoint(t *testing.T) {
	grid := NewGrid(9)
	tests := []struct {
		p  Point
		sp Point
	}{
		{Point{1, 1}, Point{9, 9}},
		{Point{3, 5}, Point{7, 5}},
	}
	for _, tt := range tests {
		want := tt.sp
		have := grid.SymmetricPoint(tt.p)
		assert.Equal(t, want, have)
	}
}
