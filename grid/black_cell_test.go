package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlackCell_String(t *testing.T) {
	tests := []struct {
		name string
		bc   *BlackCell
		want string
	}{
		{"simple", &BlackCell{}, "(0,0)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlackCell{}
			if got := bc.String(); got != tt.want {
				t.Errorf("BlackCell.String() = %v, want %v", got, tt.want)
			}
		})
	}
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
