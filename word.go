package cwcomp

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Word consists of a point and a direction
type Word struct {
	point     Point     // Point at which the word starts
	direction Direction // Across or down
	length    int       // Length of the word
	clue      string    // The text of the clue
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// NewWord is the constructor for Word
func NewWord(point Point, dir Direction, length int, clue string) *Word {
	p := new(Word)
	p.point = point
	p.direction = dir
	p.length = length
	p.clue = clue
	return p
}

// GetCrossingWords returns the words that intersect the specified word.
func (word *Word) GetCrossingWords(puzzle *Puzzle) []*Word {
	crossers := make([]*Word, 0)
	otherDirection := word.direction.Other()
	for point := range puzzle.WordIterator(word.point, word.direction) {
		otherWord := puzzle.LookupWord(point, otherDirection)
		crossers = append(crossers, otherWord)
	}
	return crossers
}

// Given a word, returns the word number for it
func (puzzle *Puzzle) GetWordNumber(word *Word) *WordNumber {
	wn := puzzle.LookupWordNumberForStartingPoint(word.point)
	return wn
}

// String returns a string representation of this object.
func (word *Word) String() string {
	parts := []string{
		fmt.Sprintf("point:%s", word.point.String()),
		fmt.Sprintf("direction:%q", word.direction.String()),
		fmt.Sprintf("length:%d", word.length),
		fmt.Sprintf("clue:%q", word.clue),
	}
	return strings.Join(parts, ",")
}

// RedoWord gets the last word change to the puzzle and re-applies it
func (puzzle *Puzzle) RedoWord() {

	// Return immediately if the redo stack is empty
	if puzzle.redoWordStack.IsEmpty() {
		return
	}

	// Pop the old Doable from the redo stack. This will give us the
	// point and direction of the word.
	oldDoable, _ := puzzle.redoWordStack.Pop()
	oldWord := &oldDoable.word
	oldText := oldDoable.text

	// Construct a new Doable from the current grid for the word in the
	// doable.
	newDoable := NewDoable(puzzle, oldWord)

	// Push the new Doable onto the undo stack.
	puzzle.undoWordStack.Push(newDoable)

	// Set the text of the old Doable in the current grid, but do not
	// use the SetText method, because it pushes the word back on the
	// undo stack.
	puzzle.SetTextWithoutPush(oldWord, oldText)
}

// UndoWord undoes the last push to the undoWordStack
func (puzzle *Puzzle) UndoWord() {

	// Return immediately if the undo stack is empty
	if puzzle.undoWordStack.IsEmpty() {
		return
	}

	// Pop the old Doable from the undo stack. This will give us the
	// point and direction of the word.
	oldDoable, _ := puzzle.undoWordStack.Pop()
	oldWord := &oldDoable.word
	oldText := oldDoable.text

	// Construct a new Doable from the current puzzle for the word in
	// the doable.
	newDoable := NewDoable(puzzle, oldWord)

	// Push the new Doable onto the redo stack.
	puzzle.redoWordStack.Push(newDoable)

	// Set the text of the old Doable in the current grid, but do not
	// use the SetText method, because it pushes the word back on the
	// undo stack.
	puzzle.SetTextWithoutPush(oldWord, oldText)
}

// WordIterator iterates through the points in a word, stopping when it
// encounters a black cell or the edge of the grid.
func (puzzle *Puzzle) WordIterator(point Point, dir Direction) <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		pp := NewPoint(point.r, point.c)
		for {
			if err := puzzle.ValidIndex(pp); err != nil {
				return
			}
			if puzzle.IsBlackCell(pp) {
				return
			}
			out <- pp
			switch dir {
			case ACROSS:
				pp.c++
			case DOWN:
				pp.r++
			}
		}
	}()
	return out
}
