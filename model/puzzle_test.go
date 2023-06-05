package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------
// Internal functions
// ---------------------------------------------------------------------

func getGoodPuzzle() *Puzzle {
	points := []Point{
		{1, 1},
		{1, 5},
		{2, 5},
		{3, 5},
		{4, 9},
		{5, 1},
		{5, 2},
		{5, 3},
	}
	puzzle := getTestPuzzle(points)
	return puzzle
}

func getTestPuzzle(points []Point) *Puzzle {
	puzzle := NewPuzzle(9)
	for _, point := range points {
		puzzle.Toggle(point)
	}
	puzzle.RenumberCells()
	return puzzle
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

func TestPuzzle_Equal(t *testing.T) {
	puzzle := getGoodPuzzle()
	assert.False(t, puzzle.Equal(nil))
	assert.True(t, puzzle.Equal(puzzle))
	bogusPuzzle := getTestPuzzle([]Point{
		{1, 2},
		{3, 4},
	})
	assert.False(t, puzzle.Equal(bogusPuzzle))
}

func TestPuzzle_GetCell(t *testing.T) {
	puzzle := getGoodPuzzle()
	puzzle.GetCell(NewPoint(4, 6)) // Good point
	assert.Panics(t, func() {
		puzzle.GetCell(NewPoint(10, -1))
	})
}

func TestPuzzle_GetClue(t *testing.T) {
	tests := []struct {
		name string
		seq  int
		dir  Direction
		want string
	}{
		{"Good across", 21, ACROSS, "21 across clue"},
		{"Good down", 19, DOWN, "19 down clue"},
		{"No across word", 13, ACROSS, ""},
		{"No down word", 21, DOWN, ""},
	}
	puzzle := getGoodPuzzle()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			if word != nil {
				word.clue = tt.want
				want := tt.want
				have, err := puzzle.GetClue(word)
				assert.Nil(t, err)
				assert.Equal(t, want, have)
			}
		})
	}
}

func TestPuzzle_GetClue_Bad(t *testing.T) {
	puzzle := getGoodPuzzle()
	word := new(Word)
	_, err := puzzle.GetClue(word)
	assert.NotNil(t, err)
}

func TestPuzzle_GetLength(t *testing.T) {
	puzzle := getGoodPuzzle()
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
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			have, _ := puzzle.GetLength(word)
			want := tt.want
			assert.Equal(t, want, have, "Failed word at %v", word.point)
		})
	}
}

func TestPuzzle_GetLength_Bad(t *testing.T) {
	var err error
	var word *Word

	puzzle := getGoodPuzzle()

	// Try a nil pointer
	word = nil
	_, err = puzzle.GetLength(word)
	assert.NotNil(t, err)

	// There is no 9 down
	word = puzzle.LookupWordByNumber(9, DOWN)
	assert.Nil(t, word)
	_, err = puzzle.GetLength(word)
	assert.NotNil(t, err)

}

// TestPuzzle_GetLength_All checks the expected length of all words in the
// grid.
func TestPuzzle_GetLength_All(t *testing.T) {
	puzzle := getGoodPuzzle()

	type test struct {
		seq  int
		dir  Direction
		want int
	}
	tests := []test{
		{1, ACROSS, 3},
		{4, ACROSS, 4},
		{8, ACROSS, 4},
		{9, ACROSS, 4},
		{10, ACROSS, 4},
		{11, ACROSS, 4},
		{12, ACROSS, 8},
		{14, ACROSS, 3},
		{15, ACROSS, 8},
		{20, ACROSS, 4},
		{21, ACROSS, 4},
		{22, ACROSS, 4},
		{23, ACROSS, 4},
		{24, ACROSS, 4},
		{25, ACROSS, 3},
		{1, DOWN, 4},
		{2, DOWN, 4},
		{3, DOWN, 9},
		{4, DOWN, 9},
		{5, DOWN, 4},
		{6, DOWN, 4},
		{7, DOWN, 3},
		{8, DOWN, 3},
		{13, DOWN, 3},
		{15, DOWN, 4},
		{16, DOWN, 4},
		{17, DOWN, 4},
		{18, DOWN, 4},
		{19, DOWN, 3},
		{20, DOWN, 3},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%d %v\n", tt.seq, tt.dir)
		t.Run(name, func(t *testing.T) {
			want := tt.want
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			have, err := puzzle.GetLength(word)
			assert.Nil(t, err)
			assert.Equal(t, want, have)
		})
	}
}

func TestPuzzle_GetText(t *testing.T) {
	type test struct {
		name       string
		seq        int
		dir        Direction
		wantLength int
		want       string
		expectOK   bool
	}
	tests := []test{
		{"Good across", 21, ACROSS, 4, "    ", true},
		{"Good down", 19, DOWN, 3, "   ", true},
		{"No across word", 13, ACROSS, 0, "", false},
		{"No down word", 21, DOWN, 0, "", false},
	}
	puzzle := getGoodPuzzle()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			have := puzzle.GetText(word)
			switch tt.expectOK {
			case true:
				assert.Equal(t, tt.wantLength, len(have))
				assert.Equal(t, want, have)
			case false:
			}
		})
	}
}

func TestPuzzle_GetText_Across(t *testing.T) {
	var word *Word

	puzzle := getGoodPuzzle()

	// It should fail if the word pointer is invalid
	word = new(Word)
	word.point = NewPoint(-1, 3)
	word.direction = ACROSS
	text := puzzle.GetText(word)
	assert.Equal(t, "", text)

	word = new(Word)
	word.point = NewPoint(100, 3)
	word.direction = ACROSS
	text = puzzle.GetText(word)
	assert.Equal(t, "", text)

	// Should return a string of the correct length
	wantLength := 3
	want := strings.Repeat(" ", wantLength)
	word = puzzle.LookupWordByNumber(14, ACROSS)
	have := puzzle.GetText(word)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestPuzzle_GetText_Bad(t *testing.T) {
	grid := getGoodPuzzle()
	grid.GetText(nil)

}
func TestPuzzle_GetText_Down(t *testing.T) {

	var word *Word
	puzzle := getGoodPuzzle()

	// It should fail if the word point is invalid
	word = new(Word)
	word.point = NewPoint(-1, 3)
	word.direction = DOWN
	text := puzzle.GetText(word)
	assert.Equal(t, "", text)

	// Should return a string of the correct length
	wantLength := 9
	word = puzzle.LookupWordByNumber(3, DOWN)

	want := strings.Repeat(" ", wantLength)
	have := puzzle.GetText(word)
	assert.Equal(t, wantLength, len(have))
	assert.Equal(t, want, have)
}

func TestPuzzle_GetTextWithLetters(t *testing.T) {
	type test struct {
		name       string
		seq        int
		dir        Direction
		wantLength int
		want       string
	}
	puzzle := getGoodPuzzle()
	puzzle.SetLetter(NewPoint(5, 4), "O")
	puzzle.SetLetter(NewPoint(5, 5), "A")
	puzzle.SetLetter(NewPoint(5, 6), "F")
	tests := []test{
		{"14 across", 14, ACROSS, 3, "OAF"},
		{"3 down", 3, DOWN, 9, "    O    "},
		{"13 down", 13, DOWN, 3, " A "},
		{"4 down", 4, DOWN, 9, "    F    "},
		{"No change to others", 15, ACROSS, 8, "        "},
		{"5 down?", 5, DOWN, 4, "    "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			have := puzzle.GetText(word)
			assert.Equal(t, tt.wantLength, len(have))
			assert.Equal(t, want, have)
		})
	}
}

func TestPuzzle_LookupWord(t *testing.T) {
	puzzle := getGoodPuzzle()
	type Test struct {
		name       string
		point      Point
		wantAcross int
		wantDown   int
		wantOK     bool
	}
	tests := []Test{
		{"middle point", NewPoint(7, 8), 21, 18, true},
		{"on word number", NewPoint(6, 2), 15, 15, true},
		{"in black cell", NewPoint(1, 5), 0, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			word := puzzle.LookupWord(tt.point, ACROSS)
			if tt.wantOK {
				assert.NotNil(t, word)
				wn := puzzle.LookupWordNumberForStartingPoint(word.point)
				assert.NotNil(t, wn)
				assert.Equal(t, tt.wantAcross, wn.seq)
			} else {
				assert.Nil(t, word)
			}

			word = puzzle.LookupWord(tt.point, DOWN)
			if tt.wantOK {
				assert.NotNil(t, word)
				wn := puzzle.LookupWordNumberForStartingPoint(word.point)
				assert.NotNil(t, wn)
				assert.Equal(t, tt.wantDown, wn.seq)
			} else {
				assert.Nil(t, word)
			}
		})
	}

}
func TestPuzzle_SetClue(t *testing.T) {
	tests := []struct {
		name   string
		seq    int
		dir    Direction
		clue   string
		wantOK bool
	}{
		{"Good across", 21, ACROSS, "21 across clue", true},
		{"Good down", 19, DOWN, "19 down clue", true},
		{"No across word", 13, ACROSS, "", false},
		{"No down word", 21, DOWN, "", false},
	}
	puzzle := getGoodPuzzle()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			word := puzzle.LookupWordByNumber(tt.seq, tt.dir)
			err := puzzle.SetClue(word, tt.clue)
			switch tt.wantOK {
			case true:
				assert.Nil(t, err)
				want, err := puzzle.GetClue(word)
				assert.Nil(t, err)
				assert.Equal(t, tt.clue, want)
			case false:
				assert.NotNil(t, err)
			}
		})
	}
}

func TestPuzzle_SetText(t *testing.T) {
	var (
		err  error
		have string
		want string
		word *Word
	)

	puzzle := getGoodPuzzle()
	type testWord struct {
		seq  int
		dir  Direction
		text string
	}
	setupWords := []testWord{
		{14, ACROSS, "OAF"},
		{13, DOWN, "TAP"},
	}
	for i, sw := range setupWords {
		word = puzzle.LookupWordByNumber(sw.seq, sw.dir)
		assert.NotNil(t, word)
		err = puzzle.SetText(word, sw.text)
		assert.Nil(t, err)
		assert.Equal(t, i+1, puzzle.undoWordStack.Len())
	}

	values := []testWord{
		{14, ACROSS, "OAF"},
		{3, DOWN, "    O    "},
		{13, DOWN, "TAP"},
		{4, DOWN, "    F    "},
		{5, DOWN, "    "},
	}
	for _, tt := range values {
		word = puzzle.LookupWordByNumber(tt.seq, tt.dir)
		assert.NotNil(t, word)
		have = puzzle.GetText(word)
		want = tt.text
		assert.Equalf(t, want, have, "%d %s", tt.seq, tt.dir)
	}
}

func TestPuzzle_SetText_Bad(t *testing.T) {
	var (
		err    error
		puzzle *Puzzle
		have   string
		text   string
		word   *Word
	)

	puzzle = getGoodPuzzle()

	// Try a non-existent word
	word = NewWord(NewPoint(6, 18), ACROSS, 6, "")
	text = "BOGUS"
	err = puzzle.SetText(word, text)
	assert.NotNil(t, err)

	// What happens if the text is shorter than the word expects?
	word = puzzle.LookupWordByNumber(21, ACROSS)
	err = puzzle.SetText(word, "X")
	assert.Nil(t, err)
	have = puzzle.GetText(word)
	assert.Equal(t, "X   ", have)

	// What happens if the text is longer than the word expects?
	word = puzzle.LookupWordByNumber(21, ACROSS)
	err = puzzle.SetText(word, "BOGUS")
	assert.NotNil(t, err)
	have = puzzle.GetText(word)
	assert.Equal(t, "X   ", have)
}

func TestPuzzle_String(t *testing.T) {

	var (
		puzzleString string
	)

	puzzle := getGoodPuzzle()
	puzzle.SetLetter(NewPoint(5, 4), "O")
	puzzle.SetLetter(NewPoint(5, 5), "A")
	puzzle.SetLetter(NewPoint(5, 6), "F")

	puzzleString = puzzle.String()
	assert.Contains(t, puzzleString, "| O | A | F |")
	assert.Equal(t, "", puzzle.GetPuzzleName())
	assert.Contains(t, puzzleString, "(Untitled)")

	// Now set the title and see that it appears in the string
	name := "MYPUZZLE"
	puzzle.SetPuzzleName(name)
	puzzleString = puzzle.String()

	assert.Contains(t, puzzleString, name)
	assert.Equal(t, name, puzzle.GetPuzzleName())

}

func TestPuzzle_LookupWordByNumber(t *testing.T) {
	var (
		puzzle *Puzzle
		word   *Word
	)

	puzzle = getGoodPuzzle()

	// Good one
	word = puzzle.LookupWordByNumber(17, DOWN)
	assert.NotNil(t, word)

	// Bad one
	word = puzzle.LookupWordByNumber(30, ACROSS)
	assert.Nil(t, word)
}

func TestPuzzle_LookupWordNumberForStartingPoint(t *testing.T) {
	puzzle := getGoodPuzzle()

	point := NewPoint(5, 4)
	want := NewWordNumber(14, point)
	have := puzzle.LookupWordNumberForStartingPoint(point)
	assert.Equal(t, want, have)

	point = Point{0, 0}
	have = puzzle.LookupWordNumberForStartingPoint(point)
	assert.Nil(t, have)
}
