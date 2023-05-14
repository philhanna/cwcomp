package model

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

func TestGrid_ToggleBad(t *testing.T) {
	grid := NewGrid(9)

	point := Point{0, 0}
	assert.Panics(t, func() {
		grid.Toggle(point)
	})

	point = Point{1, 0}
	assert.Panics(t, func() {
		grid.Toggle(point)
	})

	point = Point{0, 1}
	assert.Panics(t, func() {
		grid.Toggle(point)
	})
}

func TestGrid_RedoBlackCell(t *testing.T) {
	grid := NewGrid(9)

	// Redo should be a nop if the stack is empty
	assert.Equal(t, 0, grid.CountBlackCells())
	assert.Equal(t, 0, grid.undoStack.Len())
	assert.Equal(t, 0, grid.redoStack.Len())
	grid.RedoBlackCell()
	assert.Equal(t, 0, grid.CountBlackCells())
	assert.Equal(t, 0, grid.undoStack.Len())
	assert.Equal(t, 0, grid.redoStack.Len())

	// Add a black cell and then undo it
	grid.Toggle(Point{1, 1})
	grid.UndoBlackCell()

	// Should be zero cells
	beforeCount := grid.CountBlackCells()
	assert.Equal(t, 0, beforeCount)

	// Now redo the add black cell
	grid.RedoBlackCell()

	// Should be two black cells (symmetric twin, too)
	afterCount := grid.CountBlackCells()
	assert.Equal(t, 2, afterCount)
}

func TestGrid_Toggle(t *testing.T) {
	grid := NewGrid(9)
	points := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
	}
	for _, point := range points {
		grid.Toggle(point)
	}

	expected := []Point{
		{3, 5},
		{5, 5},
		{7, 5},
	}

	actual := []Point{}
	grid.Toggle(points[0])
	for point := range grid.PointIterator() {
		if grid.IsBlackCell(point) {
			actual = append(actual, point)
		}
	}

	assert.Equal(t, expected, actual)
}

func TestGrid_UndoBlackCell(t *testing.T) {
	grid := NewGrid(9)

	// Undo should be a nop if the stack is empty
	assert.Equal(t, 0, grid.CountBlackCells())
	assert.Equal(t, 0, grid.undoStack.Len())
	assert.Equal(t, 0, grid.redoStack.Len())
	grid.UndoBlackCell()
	assert.Equal(t, 0, grid.CountBlackCells())
	assert.Equal(t, 0, grid.undoStack.Len())
	assert.Equal(t, 0, grid.redoStack.Len())

	grid.Toggle(Point{1, 1})
	beforeCount := grid.CountBlackCells()
	assert.Equal(t, 2, beforeCount)
	grid.UndoBlackCell()
	afterCount := grid.CountBlackCells()
	assert.Equal(t, 0, afterCount)

}
