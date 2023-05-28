package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWord_String(t *testing.T) {
	puzzle := getGoodPuzzle()
	word := puzzle.LookupWordByNumber(20, DOWN)
	puzzle.SetText(word, "HOW")
	want := `point:{r:7,c:1},direction:"down",length:3,clue:""`
	have := word.String()
	assert.Equal(t, want, have)
}

func TestPuzzle_WordIterator(t *testing.T) {
	puzzle := getGoodPuzzle()
	expected := []Point{
		NewPoint(5, 4),
		NewPoint(5, 5),
		NewPoint(5, 6),
	}
	actual := []Point{}
	for point := range puzzle.WordIterator(NewPoint(5, 4), ACROSS) {
		actual = append(actual, point)
	}
	assert.Equal(t, expected, actual)
}

func TestWord_GetCrossingWords(t *testing.T) {
	puzzle := getGoodPuzzle()
	tests := []struct {
		name string
		seq  int
		dir  Direction
		want []*Word
	}{
		{"14 across", 14, ACROSS, []*Word{
			puzzle.LookupWordByNumber(3, DOWN),
			puzzle.LookupWordByNumber(13, DOWN),
			puzzle.LookupWordByNumber(4, DOWN),
		}},
		{"2 down", 2, DOWN, []*Word{
			puzzle.LookupWordByNumber(1, ACROSS),
			puzzle.LookupWordByNumber(8, ACROSS),
			puzzle.LookupWordByNumber(10, ACROSS),
			puzzle.LookupWordByNumber(12, ACROSS),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			crossers := word.GetCrossingWords(puzzle)
			assert.Equal(t, tt.want, crossers)
		})

	}
}

func TestPuzzle_GetWordNumber(t *testing.T) {
	puzzle := getGoodPuzzle()
	word := puzzle.LookupWord(NewPoint(8, 8), ACROSS)
	wn := puzzle.GetWordNumber(word)
	assert.NotNil(t, wn)
	assert.Equal(t, 23, wn.seq)
	assert.Equal(t, ACROSS, word.direction)
}

func TestPuzzle_RedoWord(t *testing.T) {
	var (
		word   *Word
		doable Doable
	)
	puzzle := getGoodPuzzle()

	// Stacks should be empty
	assert.Equal(t, 0, puzzle.undoWordStack.Len())
	assert.Equal(t, 0, puzzle.redoWordStack.Len())

	// Now set the text in a word. Undo stack should contain the old
	// value, and redo stack should be empty.
	word = puzzle.LookupWord(NewPoint(8, 8), ACROSS)
	puzzle.SetText(word, "GLOW")
	assert.Equal(t, "GLOW", puzzle.GetText(word))
	assert.Equal(t, 1, puzzle.undoWordStack.Len())
	assert.Equal(t, 0, puzzle.redoWordStack.Len())
	doable, _ = puzzle.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)
	fmt.Printf("DEBUG: doable=%s\n", doable.String())

	// Now do the undo. Undo stack should now be empty, and redo stack
	// should contain the new value.
	puzzle.UndoWord()

	assert.Equal(t, "    ", puzzle.GetText(word))
	assert.Equal(t, 0, puzzle.undoWordStack.Len())
	assert.Equal(t, 1, puzzle.redoWordStack.Len())
	doable, _ = puzzle.redoWordStack.Peek()
	assert.Equal(t, "GLOW", doable.text)

	// Do a redo, and the original grid should be there, with the last
	// operation in the undo stack.
	puzzle.RedoWord()

	assert.Equal(t, "GLOW", puzzle.GetText(word))
	assert.Equal(t, 1, puzzle.undoWordStack.Len())
	assert.Equal(t, 0, puzzle.redoWordStack.Len())
	doable, _ = puzzle.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)

	// Do one more undo, and the grid should be empty
	puzzle.UndoWord()

	for _, word := range puzzle.words {
		text := puzzle.GetText(word)
		// All the letters should be spaces
		for _, rune := range text {
			char := byte(rune)
			assert.Equal(t, byte(' '), char)
		}
	}

	puzzle.UndoWord() // Should do nothing
}

func TestPuzzle_UndoWord(t *testing.T) {
	var (
		word   *Word
		doable Doable
	)
	puzzle := getGoodPuzzle()

	// Stacks should be empty
	assert.Equal(t, 0, puzzle.undoWordStack.Len())
	assert.Equal(t, 0, puzzle.redoWordStack.Len())

	// Now set the text in a word. UndoStack should contain the old
	// value.
	word = puzzle.LookupWord(NewPoint(8, 8), ACROSS)
	puzzle.SetText(word, "GLOW")
	assert.Equal(t, "GLOW", puzzle.GetText(word))
	assert.Equal(t, 1, puzzle.undoWordStack.Len())
	assert.Equal(t, 0, puzzle.redoWordStack.Len())
	doable, _ = puzzle.undoWordStack.Peek()
	assert.Equal(t, "    ", doable.text)
	puzzle.RedoWord() // Should do nothing

	// Now do the undo. Undo stack should now be empty, and redo stack
	// should contain the new value.
	puzzle.UndoWord()

	assert.Equal(t, 0, puzzle.undoWordStack.Len())
	assert.Equal(t, 1, puzzle.redoWordStack.Len())
	doable, _ = puzzle.redoWordStack.Peek()
	assert.Equal(t, "GLOW", doable.text)

}
