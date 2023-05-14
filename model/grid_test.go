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

func TestGrid_GetClue(t *testing.T) {
	tests := []struct {
		name string
		seq  int
		dir  Direction
		want string
	}{
		{"Good across", 21, ACROSS, ""},
		{"Good down", 19, DOWN, ""},
		{"No across word", 13, ACROSS, ""},
		{"No down word", 21, DOWN, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := getGoodGrid()
			have := grid.GetClue(tt.seq, tt.dir)
			assert.Equal(t, tt.want, have)
		})
	}
}

func TestGrid_GetClue_Bad(t *testing.T) {
	grid := getGoodGrid()
	assert.Panics(t, func() {
		grid.GetClue(-1, ACROSS)
	})
}

func TestGrid_GetLength(t *testing.T) {
	grid := getGoodGrid()
	tests := []struct {
		name string
		seq  int
		dir  Direction
		want int
	}{
		{"Good both(a)", 4, ACROSS, 4},
		{"Good both(d)", 4, DOWN, 9},
		{"Good across", 25, ACROSS, 3},
		{"Good down", 19, DOWN, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, grid.GetLength(tt.seq, tt.dir))
		})
	}
}

func TestGrid_GetLength_Bad(t *testing.T) {
	grid := getGoodGrid()
	assert.Panics(t, func() {
		grid.GetLength(-3, ACROSS)
	})
}

func TestGrid_GetText(t *testing.T) {
	type test struct {
		name       string
		seq        int
		dir        Direction
		wantLength int
		want       string
	}
	tests := []test{
		{"Good across", 21, ACROSS, 4, "    "},
		{"Good down", 19, DOWN, 3, "   "},
		{"No across word", 13, ACROSS, 0, ""},
		{"No down word", 21, DOWN, 0, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grid := getGoodGrid()
			want := tt.want
			have := grid.GetText(tt.seq, tt.dir)
			assert.Equal(t, tt.wantLength, len(have))
			assert.Equal(t, want, have)
		})
	}
}

func TestGrid_GetText_Across(t *testing.T) {
	grid := getGoodGrid()

	// It should fail if the word sequence number is invalid
	assert.Panics(t, func() {
		grid.GetText(-1, ACROSS)
	})
	assert.Panics(t, func() {
		grid.GetText(1000, ACROSS)
	})

	// Should return a string of the correct length

	wantLength := 3
	want := strings.Repeat(" ", wantLength)
	have := grid.GetText(14, ACROSS)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestGrid_GetText_Down(t *testing.T) {
	grid := getGoodGrid()

	// It should fail if the word sequence number is invalid
	assert.Panics(t, func() {
		grid.GetText(-1, DOWN)
	})
	assert.Panics(t, func() {
		grid.GetText(1000, DOWN)
	})

	// Should return a string of the correct length
	wantLength := 9
	want := strings.Repeat(" ", wantLength)
	have := grid.GetText(3, DOWN)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestGrid_GetTextWithLetters(t *testing.T) {
	type test struct {
		name       string
		seq        int
		dir        Direction
		wantLength int
		want       string
	}
	grid := getGoodGrid()
	grid.SetLetter(NewPoint(5, 4), "O")
	grid.SetLetter(NewPoint(5, 5), "A")
	grid.SetLetter(NewPoint(5, 6), "F")

	tests := []test{
		{"14 across", 14, ACROSS, 3, "OAF"},
		{"3 down", 3, DOWN, 9, "    O    "},
		{"13 down", 13, DOWN, 3, " A "},
		{"4 down", 4, DOWN, 9, "    F    "},
		{"No change to others", 15, ACROSS, 8, "        "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := grid.GetText(tt.seq, tt.dir)
			assert.Equal(t, tt.wantLength, len(have))
			assert.Equal(t, want, have)
		})

	}

}

func TestGrid_String(t *testing.T) {
	grid := getGoodGrid()
	grid.SetLetter(NewPoint(5, 4), "O")
	grid.SetLetter(NewPoint(5, 5), "A")
	grid.SetLetter(NewPoint(5, 6), "F")
	gridString := grid.String()
	substring := "| O | A | F |"
	assert.Contains(t, gridString, substring)
}
