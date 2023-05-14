package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlackCell_String(t *testing.T) {
	bc := NewBlackCell(Point{0, 0})
	want := "{r:0,c:0}"
	have := bc.String()
	assert.Equal(t, want, have)
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
