package cwcomp

import (
	"fmt"
	"log"
)

// DumpPuzzle is a diagnostic function that shows the exact composition
// of each cell in the grid.
func DumpPuzzle(puzzle *Puzzle) {

	dumpClues := func(direction Direction) {
		for _, wn := range puzzle.wordNumbers {
			seq := wn.seq
			word := puzzle.LookupWordByNumber(seq, direction)
			if word != nil {
				fmt.Printf("%d. %s\n", seq, word.clue)
			}
		}
	}

	fmt.Println("Across:")
	dumpClues(ACROSS)

	fmt.Println("Down:")
	dumpClues(DOWN)
}

// TracePuzzle is a diagnostic function that shows how each cell is
// constructed
func TracePuzzle(puzzle *Puzzle) {
	n := puzzle.n
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := NewPoint(r, c)
			cell := puzzle.GetCell(point)
			switch typedCell := cell.(type) {
			case BlackCell:
				log.Printf("BlackCell at  %v has value %v\n", point, typedCell.String())
			case LetterCell:
				log.Printf("LetterCell at %v has value %v\n", point, typedCell.String())
			}
		}
	}
	fmt.Println(puzzle.String())

}
