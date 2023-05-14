package model

import (
	"strings"
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
	grid := getTestGrid(points)
	return grid
}

func getTestGrid(points []Point) *Grid {
	grid := NewGrid(9)
	for _, point := range points {
		grid.Toggle(point)
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
		grid.Toggle(point)
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

func TestGrid_CountBlackCells(t *testing.T) {
	grid := NewGrid(9)
	assert.Equal(t, 0, grid.CountBlackCells())

	grid = getGoodGrid()
	assert.Equal(t, 16, grid.CountBlackCells())

	grid.Toggle(Point{1, 6})
	assert.Equal(t, 18, grid.CountBlackCells())
	// 18 because we added a point and its symmetric point
}

func TestGrid_GetAcrossWordLength(t *testing.T) {
	grid := getGoodGrid()
	point := Point{1, 2}
	assert.Equal(t, 3, grid.GetAcrossWordLength(&point))
}

func TestGrid_GetAcrossWordText(t *testing.T) {
	grid := getGoodGrid()

	// It should fail if the word sequence number is invalid
	assert.Panics(t, func() {
		grid.GetAcrossWordText(-1)
	})
	assert.Panics(t, func() {
		grid.GetAcrossWordText(1000)
	})

	// Should return a string of the correct length

	wantLength := 3
	want := strings.Repeat(" ", wantLength)
	have := grid.GetAcrossWordText(14)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestGrid_GetDownWordLength(t *testing.T) {
	grid := getGoodGrid()
	point := Point{1, 2}
	assert.Equal(t, 4, grid.GetDownWordLength(&point))
}

func TestGrid_GetDownWordText(t *testing.T) {
	grid := getGoodGrid()

	// It should fail if the word sequence number is invalid
	assert.Panics(t, func() {
		grid.GetDownWordText(-1)
	})
	assert.Panics(t, func() {
		grid.GetDownWordText(1000)
	})

	// Should return a string of the correct length
	wantLength := 9
	want := strings.Repeat(" ", wantLength)
	have := grid.GetDownWordText(3)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestGrid_GetWordText(t *testing.T) {
	tests := []struct {
		name       string
		seq        int
		dir        Direction
		wantLength int
		want       string
	}{
		{"Good across", 21, ACROSS, 4, "    "},
		{"Good down", 19, DOWN, 3, "   "},
		{"No across word", 13, ACROSS, 0, ""},
		{"No down word", 21, DOWN, 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := getGoodGrid()
			want := tt.want
			have := grid.GetWordText(tt.seq, tt.dir)
			assert.Equal(t, tt.wantLength, len(have))
			assert.Equal(t, want, have)
		})
	}
}

func TestGrid_LetterCellIterator(t *testing.T) {
	grid := getGoodGrid()
	nlc := 0
	for range grid.LetterCellIterator() {
		nlc++
	}
	assert.Equal(t, 9*9-16, nlc)
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
