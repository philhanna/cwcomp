package cwcomp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDoable(t *testing.T) {
	var (
		doable Doable
	)
	puzzle := getGoodPuzzle()
	word := puzzle.LookupWordByNumber(8, ACROSS)

	// Check it before setting the text of the word

	doable = NewDoable(puzzle, word)
	assert.Equal(t, NewPoint(2, 1), doable.word.point)
	assert.Equal(t, ACROSS, doable.word.direction)
	assert.Equal(t, 4, doable.word.length)
	assert.Equal(t, "", doable.word.clue)
	assert.Equal(t, "    ", doable.text)

	puzzle.SetText(word, "BLUE")

	// Check it after setting the text of the word

	doable = NewDoable(puzzle, word)
	assert.Equal(t, NewPoint(2, 1), doable.word.point)
	assert.Equal(t, ACROSS, doable.word.direction)
	assert.Equal(t, 4, doable.word.length)
	assert.Equal(t, "", doable.word.clue)
	assert.Equal(t, "BLUE", doable.text)
}
