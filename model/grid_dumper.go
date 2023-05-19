package model

import (
	"fmt"
	"log"
)

// dumpGrid is a diagnostic function that shows the exact composition of
// each cell in the grid.
func dumpGrid(grid *Grid) {
	n := grid.n
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := NewPoint(r, c)
			cell := grid.GetCell(point)
			switch typedCell := cell.(type) {
			case BlackCell:
				log.Printf("BlackCell    at %v has value %v\n", point, typedCell.String())
			case LetterCell:
				log.Printf("LetterCell   at %v has value %v\n", point, typedCell.String())
			}
		}
	}
	fmt.Println(grid.String())
}
