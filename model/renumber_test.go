package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNumberedCellsFromGrid(t *testing.T) {
	grid := getGoodGrid()
	cells := GridToCells(grid)
	ncs := GetNumberedCells(cells)
	for _, nc := range ncs {
		fmt.Printf("%v\n", nc)
	}
	assert.Equal(t, 25, len(ncs))
	tests := []struct {
		name      string // test name
		seq       int    // The word number (1, 2, ...)
		row       int    // The row number (1, 2, ..., n)
		col       int    // The column number (1, 2, ..., n)
		startA    bool   // This is the start of an across word
		startD    bool   // This is the start of a down word
		wantIndex int    // The index into the numbered cells array
	}{
		{"7 down only", 7, 1, 9, false, true, 6},
		{"14 across only", 14, 5, 4, true, false, 13},
		{"20 both", 20, 7, 1, true, true, 19},
	}
	for _, tt := range tests {
		want := ncs[tt.wantIndex]
		have := NumberedCell{tt.wantIndex + 1, tt.row, tt.col, tt.startA, tt.startD}
		assert.Equal(t, want, have)
	}
	for _, nc := range ncs {
		fmt.Println(nc.String())
	}
	dumpGrid(grid)

}

func TestGetNumberedCells(t *testing.T) {
	n := 9
	cells := make([][]byte, n)
	for i := 0; i < n; i++ {
		cells[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			cells[i][j] = ' '
		}
	}
	blackCells := []struct{ r, c int }{
		{1, 1},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 9},
		{5, 1},
		{5, 2},
		{5, 3},
		{5, 7},
		{5, 8},
		{5, 9},
		{6, 1},
		{7, 5},
		{8, 5},
		{9, 5},
		{9, 9},
	}
	for _, bc := range blackCells {
		cells[bc.r-1][bc.c-1] = '\x00'
	}

	ncs := GetNumberedCells(cells)
	for _, nc := range ncs {
		fmt.Println(nc.String())
	}

	assert.Equal(t, 25, len(ncs))
}
