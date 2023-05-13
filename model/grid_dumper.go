package model

import "log"

// dumpGrid is a utility class that logs the exact contents of a grid
func dumpGrid(grid *Grid) {
	n := grid.n
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			point := Point{r, c}
			cell := grid.GetCell(point)
			switch cellType := cell.(type) {
			case BlackCell:
				bc := cell.(BlackCell)
				log.Printf("BlackCell    at %v has value %v\n", point, bc.String())
			case LetterCell:
				lc := cell.(LetterCell)
				log.Printf("LetterCell   at %v has value %v\n", point, lc.String())
			case NumberedCell:
				nc := cell.(NumberedCell)
				log.Printf("NumberedCell at %v has value %v\n", point, nc.String())
			default:
				log.Printf("???????????  at %v is type %s, value %v\n", point, cellType, cell)
			}
		}
	}
}
