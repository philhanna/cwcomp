package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord_String(t *testing.T) {
	grid := getGoodGrid()
	word := grid.LookupWordByNumber(20, DOWN)
	grid.SetText(word, "HOW")
	want := `point:{r:7,c:1},direction:"down",length:3,clue:""`
	have := word.String()
	assert.Equal(t, want, have)
}

func TestGrid_WordIterator(t *testing.T) {
	grid := getGoodGrid()
	expected := []Point{
		NewPoint(5, 4),
		NewPoint(5, 5),
		NewPoint(5, 6),
	}
	actual := []Point{}
	for point := range grid.WordIterator(NewPoint(5, 4), ACROSS) {
		actual = append(actual, point)
	}
	assert.Equal(t, expected, actual)
}

func TestWord_GetCrossingWords(t *testing.T) {
	grid := getGoodGrid()
	tests := []struct {
		name string
		seq  int
		dir  Direction
		want []*Word
	}{
		{"14 across", 14, ACROSS, []*Word{
			grid.LookupWordByNumber(3, DOWN),
			grid.LookupWordByNumber(13, DOWN),
			grid.LookupWordByNumber(4, DOWN),
		}},
		{"2 down", 2, DOWN, []*Word{
			grid.LookupWordByNumber(1, ACROSS),
			grid.LookupWordByNumber(8, ACROSS),
			grid.LookupWordByNumber(10, ACROSS),
			grid.LookupWordByNumber(12, ACROSS),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := grid.LookupWordByNumber(tt.seq, tt.dir)
			crossers := word.GetCrossingWords(grid)
			assert.Equal(t, tt.want, crossers)
		})

	}
}

func TestGrid_GetWordNumber(t *testing.T) {
	grid := getGoodGrid()
	word := grid.LookupWord(NewPoint(8, 8), ACROSS)
	wn := grid.GetWordNumber(word)
	assert.NotNil(t, wn)
	assert.Equal(t, 23, wn.seq)
	assert.Equal(t, ACROSS, word.direction)
}
