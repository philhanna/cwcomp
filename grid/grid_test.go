package grid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dumpGrid(grid *Grid) {
	n := grid.n
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cellType := cell.(type) {
			case BlackCell:
				bc := cell.(BlackCell)
				fmt.Printf("BlackCell    at %v has value %v\n", point, bc.String())
			case LetterCell:
				lc := cell.(LetterCell)
				fmt.Printf("LetterCell   at %v has value %v\n", point, lc.String())
			case NumberedCell:
				nc := cell.(NumberedCell)
				fmt.Printf("NumberedCell at %v has value %v\n", point, nc.String())
			default:
				fmt.Printf("???????????  at %v is type %s, value %v\n", point, cellType, cell)
			}
		}
	}
}

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

func TestGrid_RemoveBlackCell(t *testing.T) {
	grid := NewGrid(9)
	points := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
	}
	for _, point := range points {
		grid.AddBlackCell(point)
	}

	expected := []Point{
		{3, 5},
		{5, 5},
		{7, 5},
	}

	actual := []Point{}
	grid.RemoveBlackCell(points[0])
	for point := range grid.PointIterator() {
		if grid.IsBlackCell(point) {
			actual = append(actual, point)
		}
	}

	assert.Equal(t, expected, actual)
}

func TestGrid_AddBlackCell(t *testing.T) {
	tests := []struct {
		p  Point
		sp Point
	}{
		{Point{1, 1}, Point{9, 9}},
		{Point{3, 5}, Point{7, 5}},
		{Point{5, 5}, Point{5, 5}},
	}
	for _, tt := range tests {
		grid := NewGrid(9)
		point := tt.p
		grid.AddBlackCell(point)

		cell := grid.GetCell(point)
		switch cellType := cell.(type) {
		case BlackCell: // OK
		default:
			t.Errorf("Point %v should be black cell, not %v", point, cellType)
		}

		symPoint := tt.sp
		symCell := grid.GetCell(symPoint)
		switch cellType := symCell.(type) {
		case BlackCell: // OK
		default:
			t.Errorf("Symmetric point %v should be black cell, not %v", symPoint, cellType)
		}
	}

}

func TestGrid_FindNumberedCells_Bad(t *testing.T) {
	tests := []struct {
		name       string
		blackCells []Point
		nBC        int
		nLC        int
		nNC        int
	}{
		{
			"Bad",
			[]Point{
				{1, 1}, {1, 5},
				{2, 5},
				{3, 5},
				{4, 6},
				{4, 9},
				{5, 1}, {5, 2}, {5, 3},
				{7, 1},
			},
			20, 32, 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := getTestGrid(tt.blackCells)
			have_nBC := 0
			have_nLC := 0
			have_nNC := 0
			for r := 1; r <= grid.n; r++ {
				for c := 1; c <= grid.n; c++ {
					point := Point{r, c}
					cell := grid.GetCell(point)
					switch cell.(type) {
					case BlackCell:
						have_nBC++
					case LetterCell:
						have_nLC++
					case NumberedCell:
						have_nNC++
					}
				}
			}
			//dumpGrid(grid)
			assert.Equal(t, tt.nBC, have_nBC)
			assert.Equal(t, tt.nLC, have_nLC)
			assert.Equal(t, tt.nNC, have_nNC)
		})
	}

}

func TestGrid_FindNumberedCells_Good(t *testing.T) {
	tests := []struct {
		name       string
		blackCells []Point
		nBC        int
		nLC        int
		nNC        int
	}{
		// Add test cases
		{
			"Good",
			[]Point{
				{1, 1}, {1, 5},
				{2, 5},
				{3, 5},
				{4, 9},
				{5, 1}, {5, 2}, {5, 3},
			},
			16, 40, 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := getTestGrid(tt.blackCells)
			dumpGrid(grid)
			have_nBC := 0
			have_nLC := 0
			have_nNC := 0
			for r := 1; r <= grid.n; r++ {
				for c := 1; c <= grid.n; c++ {
					point := Point{r, c}
					cell := grid.GetCell(point)
					switch cell.(type) {
					case BlackCell:
						have_nBC++
					case LetterCell:
						have_nLC++
					case NumberedCell:
						have_nNC++
					}
				}
			}
			assert.Equal(t, tt.nBC, have_nBC, "Black cell count")
			assert.Equal(t, tt.nLC, have_nLC, "Letter cell count")
			assert.Equal(t, tt.nNC, have_nNC, "Numbered cell count")
		})
	}

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
