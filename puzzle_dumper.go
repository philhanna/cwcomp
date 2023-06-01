package cwcomp

import (
	"fmt"
	"log"
)

// DumpPuzzle is a diagnostic function that shows the exact composition
// of each cell in the grid.
func DumpPuzzle(puzzle *Puzzle) {

	fmt.Printf(puzzle.String())

	dumpClues := func(direction Direction) {
		nPrinted := 0
		for _, wn := range puzzle.wordNumbers {
			seq := wn.seq
			word := puzzle.LookupWordByNumber(seq, direction)
			if word != nil && word.clue != "" {
				nPrinted++
				fmt.Printf("%d. %s\n", seq, word.clue)
			}
		}
		
		switch nPrinted {
		case 0: fmt.Printf("\tNo non-blank clues\n")
		case 1: fmt.Printf("\t1 non-blank clue\n")
		default: fmt.Printf("\t%d non-blank clues\n", nPrinted)
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
				log.Printf("BlackCell at  %v\n", point)
			case LetterCell:
				log.Printf("LetterCell at %v has value %v\n", point, typedCell.String())
			}
		}
	}

}
