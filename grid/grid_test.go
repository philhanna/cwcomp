package grid

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dumpGrid(grid *Grid) {
	n := grid.n
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				log.Printf("BlackCell    at %v is %T %v\n", point, cell, cell)
			case LetterCell:
				log.Printf("LetterCell   at %v is %T %+v\n", point, cell, cell)
			case NumberedCell:
				log.Printf("NumberedCell at %v is %T %v\n", point, cell, cell)
			default:
				log.Printf("???????????  at %v is %T %v\n", point, cell, cell)
			}
		}
	}
}

func getBadGrid() *Grid {
	points := []Point{
		{1, 1}, {1, 5},
		{2, 5},
		{3, 5},
		{4, 6},
		{4, 9},
		{5, 1}, {5, 2}, {5, 3},
	}
	return getTestGrid(points)
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
	grid.FindNumberedCells()
	return grid
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
		symPoint := tt.sp
		grid.AddBlackCell(point)

		cell := grid.GetCell(point)
		switch cellType := cell.(type) {
		case BlackCell: // OK
		default:
			t.Errorf("Point %v should be black cell, not %v", point, cellType)
		}

		symCell := grid.GetCell(symPoint)
		switch cellType := symCell.(type) {
		case BlackCell: // OK
		default:
			t.Errorf("Symmetric point %v should be black cell, not %v", symPoint, cellType)
		}
	}

}

func TestGrid_FindNumberedCells_BadGrid(t *testing.T) {
	grid := getBadGrid()
	dumpGrid(grid)
	n := grid.n
	countBlackCells := 0
	countLetterCells := 0
	countNumberedCells := 0
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cellType := cell.(type) {
			case BlackCell:
				countBlackCells++
			case LetterCell:
				countLetterCells++
			case NumberedCell:
				countNumberedCells++
			default:
				t.Fatalf("Unrecognized type %v\n", cellType)
			}
		}
	}
	assert.Equal(t, 18, countBlackCells)
	assert.Equal(t, 35, countLetterCells)
	assert.Equal(t, 28, countNumberedCells)
}

func TestGrid_FindNumberedCells_GoodGrid(t *testing.T) {
	grid := getGoodGrid()
	n := grid.n
	countBlackCells := 0
	countLetterCells := 0
	countNumberedCells := 0
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cellType := cell.(type) {
			case BlackCell:
				countBlackCells++
			case LetterCell:
				countLetterCells++
			case NumberedCell:
				countNumberedCells++
			default:
				t.Fatalf("Unrecognized type %v\n", cellType)
			}
		}
	}
	assert.Equal(t, 16, countBlackCells)
	assert.Equal(t, 40, countLetterCells)
	assert.Equal(t, 25, countNumberedCells)
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
