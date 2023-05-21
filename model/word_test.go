package model

import (
	"fmt"
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

func TestGrid_RedoWord(t *testing.T) {
	var (
		word   *Word
		doable Doable
	)
	grid := getGoodGrid()

	// Stacks should be empty
	assert.Equal(t, 0, grid.undoWordStack.Len())
	assert.Equal(t, 0, grid.redoWordStack.Len())

	// Now set the text in a word. Undo stack should contain the old
	// value, and redo stack should be empty.
	word = grid.LookupWord(NewPoint(8, 8), ACROSS)
	grid.SetText(word, "GLOW")
	assert.Equal(t, "GLOW", grid.GetText(word))
	assert.Equal(t, 1, grid.undoWordStack.Len())
	assert.Equal(t, 0, grid.redoWordStack.Len())
	doable, _ = grid.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)
	fmt.Printf("DEBUG: doable=%s\n", doable.String())

	// Now do the undo. Undo stack should now be empty, and redo stack
	// should contain the new value.
	grid.UndoWord()

	assert.Equal(t, "    ", grid.GetText(word))
	assert.Equal(t, 0, grid.undoWordStack.Len())
	assert.Equal(t, 1, grid.redoWordStack.Len())
	doable, _ = grid.redoWordStack.Peek()
	assert.Equal(t, "GLOW", doable.text)

	// Do a redo, and the original grid should be there, with the last
	// operation in the undo stack.
	grid.RedoWord()

	assert.Equal(t, "GLOW", grid.GetText(word))
	assert.Equal(t, 1, grid.undoWordStack.Len())
	assert.Equal(t, 0, grid.redoWordStack.Len())
	doable, _ = grid.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)

	// Do one more undo, and the grid should be empty
	grid.UndoWord()

	for _, word := range grid.words {
		text := grid.GetText(word)
		// All the letters should be spaces
		for _, rune := range text {
			char := byte(rune)
			assert.Equal(t, byte(' '), char)
		}
	}

	grid.UndoWord() // Should do nothing
}

func TestGrid_UndoWord(t *testing.T) {
	var (
		word   *Word
		doable Doable
	)
	grid := getGoodGrid()

	// Stacks should be empty
	assert.Equal(t, 0, grid.undoWordStack.Len())
	assert.Equal(t, 0, grid.redoWordStack.Len())

	// Now set the text in a word. UndoStack should contain the old
	// value.
	word = grid.LookupWord(NewPoint(8, 8), ACROSS)
	grid.SetText(word, "GLOW")
	assert.Equal(t, "GLOW", grid.GetText(word))
	assert.Equal(t, 1, grid.undoWordStack.Len())
	assert.Equal(t, 0, grid.redoWordStack.Len())
	doable, _ = grid.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)
	grid.RedoWord() // Should do nothing

	// Now do the undo. Undo stack should now be empty, and redo stack
	// should contain the new value.
	grid.UndoWord()

	assert.Equal(t, 0, grid.undoWordStack.Len())
	assert.Equal(t, 1, grid.redoWordStack.Len())
	doable, _ = grid.redoWordStack.Peek()
	assert.Equal(t, "GLOW", doable.text)

}
