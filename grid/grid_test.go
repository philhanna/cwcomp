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
			switch cellType := cell.(type) {
			case BlackCell:
				bc := cell.(BlackCell)
				log.Printf("BlackCell    at %v has value %v\n", point, bc.String())
			case LetterCell:
				lc := cell.(LetterCell)
				log.Printf("LetterCell   at %v has value %v\n", point, lc.String())
			case NumberedCell:
				nc := cell.(NumberedCell)
				lc := nc.LetterCell
				log.Printf("NumberedCell at %v has value %v, seq:%d, aLen:%d, dLen:%d\n", point, lc.String(), nc.seq, nc.aLen, nc.dLen)
			default:
				log.Printf("???????????  at %v is type %s, value %v\n", point, cellType, cell)
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

func TestGrid_FindNumberedCells(t *testing.T) {
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
			20, 34, 27,
		},
	}
	for i, tt := range tests {
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
			if i == 0 {
				dumpGrid(grid)
			}
			assert.Equal(t, tt.nBC, have_nBC)
			assert.Equal(t, tt.nLC, have_nLC)
			assert.Equal(t, tt.nNC, have_nNC)
		})
	}

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
