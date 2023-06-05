package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlackCell_String(t *testing.T) {
	bc := NewBlackCell(NewPoint(0, 0))
	want := "{r:0,c:0}"
	have := bc.String()
	assert.Equal(t, want, have)
}

func TestPuzzle_BlackCellIterator(t *testing.T) {
	points := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
	}
	puzzle := NewPuzzle(9)
	for _, point := range points {
		puzzle.Toggle(point)
	}

	expected := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
		{7, 5},
		{9, 9},
	}
	actual := []Point{}
	for bc := range puzzle.BlackCellIterator() {
		actual = append(actual, bc.point)
	}
	assert.Equal(t, expected, actual)
}

func TestPuzzle_CountBlackCells(t *testing.T) {
	puzzle := NewPuzzle(9)
	assert.Equal(t, 0, puzzle.CountBlackCells())

	puzzle = getGoodPuzzle()
	assert.Equal(t, 16, puzzle.CountBlackCells())

	puzzle.Toggle(NewPoint(1, 6))
	assert.Equal(t, 18, puzzle.CountBlackCells())
	// 18 because we added a point and its symmetric point
}

func TestBlackCell_GetPoint(t *testing.T) {
	bc := NewBlackCell(NewPoint(1, 2))
	want := NewPoint(1, 2)
	have := bc.GetPoint()
	assert.Equal(t, want, have)

	bc = *new(BlackCell)
	assert.Equal(t, Point{0, 0}, bc.GetPoint())
}

func TestPuzzle_RedoBlackCell(t *testing.T) {
	puzzle := NewPuzzle(9)

	// Redo should be a nop if the stack is empty
	assert.Equal(t, 0, puzzle.CountBlackCells())
	assert.Equal(t, 0, puzzle.undoPointStack.Len())
	assert.Equal(t, 0, puzzle.redoPointStack.Len())
	puzzle.RedoBlackCell()
	assert.Equal(t, 0, puzzle.CountBlackCells())
	assert.Equal(t, 0, puzzle.undoPointStack.Len())
	assert.Equal(t, 0, puzzle.redoPointStack.Len())

	// Add a black cell and then undo it
	puzzle.Toggle(NewPoint(1, 1))
	puzzle.UndoBlackCell()

	// Should be zero cells
	beforeCount := puzzle.CountBlackCells()
	assert.Equal(t, 0, beforeCount)

	// Now redo the add black cell
	puzzle.RedoBlackCell()

	// Should be two black cells (symmetric twin, too)
	afterCount := puzzle.CountBlackCells()
	assert.Equal(t, 2, afterCount)
}

func TestPuzzle_Toggle(t *testing.T) {
	puzzle := NewPuzzle(9)
	points := []Point{
		{1, 1},
		{3, 5},
		{5, 5},
	}
	for _, point := range points {
		puzzle.Toggle(point)
	}

	expected := []Point{
		{3, 5},
		{5, 5},
		{7, 5},
	}

	actual := []Point{}
	puzzle.Toggle(points[0])
	for point := range puzzle.PointIterator() {
		if puzzle.IsBlackCell(point) {
			actual = append(actual, point)
		}
	}

	assert.Equal(t, expected, actual)
}

func TestPuzzle_Toggle_Bad(t *testing.T) {
	puzzle := NewPuzzle(9)

	point := NewPoint(0, 0)
	assert.Panics(t, func() {
		puzzle.Toggle(point)
	})

	point = NewPoint(1, 0)
	assert.Panics(t, func() {
		puzzle.Toggle(point)
	})

	point = NewPoint(0, 1)
	assert.Panics(t, func() {
		puzzle.Toggle(point)
	})
}

func TestPuzzle_UndoBlackCell(t *testing.T) {
	puzzle := NewPuzzle(9)

	// Undo should be a nop if the stack is empty
	assert.Equal(t, 0, puzzle.CountBlackCells())
	assert.Equal(t, 0, puzzle.undoPointStack.Len())
	assert.Equal(t, 0, puzzle.redoPointStack.Len())
	puzzle.UndoBlackCell()
	assert.Equal(t, 0, puzzle.CountBlackCells())
	assert.Equal(t, 0, puzzle.undoPointStack.Len())
	assert.Equal(t, 0, puzzle.redoPointStack.Len())

	puzzle.Toggle(NewPoint(1, 1))
	beforeCount := puzzle.CountBlackCells()
	assert.Equal(t, 2, beforeCount)
	puzzle.UndoBlackCell()
	afterCount := puzzle.CountBlackCells()
	assert.Equal(t, 0, afterCount)
}
