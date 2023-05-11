package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberedCell_assignNumberedCells(t *testing.T) {
	// Create map of the expected word numbers to cell points
	expected := make(map[int]Point)
	expected[1] = Point{1, 2}
	expected[2] = Point{1, 3}
	expected[3] = Point{1, 4}
	expected[4] = Point{1, 6}
	expected[5] = Point{1, 7}
	expected[6] = Point{1, 8}
	expected[7] = Point{1, 9}
	expected[8] = Point{2, 1}
	expected[9] = Point{2, 6}
	expected[10] = Point{3, 1}
	expected[11] = Point{3, 6}
	expected[12] = Point{4, 1}
	expected[13] = Point{4, 5}
	expected[14] = Point{5, 4}
	expected[15] = Point{6, 2}
	expected[16] = Point{6, 3}
	expected[17] = Point{6, 7}
	expected[18] = Point{6, 8}
	expected[19] = Point{6, 9}
	expected[20] = Point{7, 1}
	expected[21] = Point{7, 6}
	expected[22] = Point{8, 1}
	expected[23] = Point{8, 6}
	expected[24] = Point{9, 1}
	expected[25] = Point{9, 6}

	grid := getGoodGrid()
	actual := make(map[int]Point)
	for nc := range grid.NumberedCellIterator() {
		actual[nc.wordNumber] = nc.point
	}

	assert.Equal(t, expected, actual)
}

func TestGrid_RenumberCells_Bad(t *testing.T) {
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
			assert.Equal(t, tt.nBC, have_nBC)
			assert.Equal(t, tt.nLC, have_nLC)
			assert.Equal(t, tt.nNC, have_nNC)
		})
	}

}

func TestGrid_RenumberCells_Good(t *testing.T) {
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

func TestNumberedCell_String(t *testing.T) {
	tests := []struct {
		name string
		lc   LetterCell
		seq  int
		aLen int
		dLen int
		want string
	}{
		{
			"simple",
			LetterCell{
				point:    Point{1, 2},
				ncAcross: nil,
				ncDown:   nil,
				letter:   "O",
			},
			1,
			3,
			4,
			`LetterCell:{point:{1,2},ncAcross:<nil>,ncDown:<nil>,letter:"O"},seq:1,aLen:3,dLen:4`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NumberedCell{tt.lc, tt.seq, tt.aLen, tt.dLen}
			assert.Equal(t, tt.want, nc.String())
		})
	}
}
