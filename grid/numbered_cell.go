package grid

import (
	"fmt"
	"log"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// NumberedCell is a letter cell that is the beginning of an across word
// and/or a down word.
type NumberedCell struct {
	LetterCell
	seq    int    // The word number
	aLen   int    // Length of the across word (if any)
	dLen   int    // Length of the down word (if any)
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// FindNumberedCells figures out the numbering of the across and down
// words in this word based on the locations of all the black cells.
//
// Algorithm:
//   - Iterate through all cells by row and column numbers (1, 2, ..., n).
//   - For each r, c:
//   - Ignore black cells.
//   - See if this is the beginning of an across word.
//     If so, find the row and column of the stopping point,
//     which is either the next black cell or the edge of the puzzle.
//   - Do the same for down words
//   - If either an across or down word is found:
//   - Create a new WordNumber with the next available number
//   - TODO set the type of the starting point to NumberedCell
func (grid *Grid) FindNumberedCells() {

	n := grid.n

	// The numbered cell counter
	seq := 0

	// Iterate through all the cells of the grid by row, column
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {

			// Get the current point
			startingPoint := Point{r, c}
			startingCell := grid.GetCell(startingPoint)

			// Ignore black cells
			if grid.IsBlackCell(startingPoint) {
				continue
			}

			// Create a structure to hold the potential new numbered
			// cell, and save its previous contents as a letter cell. We
			// do this because this grid calculation may be done when
			// some letters have already been entered in the grid
			// (EDITING_PUZZLE).
			nc := new(NumberedCell)
			switch startingCell.(type) {
			case LetterCell:
				nc.LetterCell = startingCell.(LetterCell)
			case BlackCell:
				nc.LetterCell = NewLetterCell(startingPoint)
			case NumberedCell:
				nc.LetterCell = NewLetterCell(startingPoint)
			}

			// This is the beginning point of an across word if either:
			//  - It is in the first column
			//  - The previous column on the left is a black cell
			if startingPoint.Col == 1 || grid.IsBlackCell(Point{r, c - 1}) {

				// Find the ending point of this across word
				nc.aLen = 0
				thisPoint := startingPoint
				for thisPoint.Col <= n && !grid.IsBlackCell(thisPoint) {
					nc.aLen++
					thisPoint.Col++
				}

				// Word lengths must be >= 3 to qualify. If this one
				// isn't, set its length back to zero.
				if nc.aLen < 3 {
					log.Printf("WARNING: at (%d,%d), across word is invalid\n", r, c)
					nc.aLen = 0
				}

				if nc.aLen > 0 {

					// This is a good across word. We need to set a
					// pointer back to the starting cell in each of its
					// cells (except the starting point itself).
					thisPoint = startingPoint

					for i := 1; i < nc.aLen; i++ {
						thisPoint.Col++
						cell := grid.GetCell(thisPoint).(LetterCell)
						cell.ncAcross = &startingPoint
						grid.SetCell(thisPoint, cell)
					}
				}
			}

			// This is the beginning point of a down word if either:
			//  - It is in the first row
			//  - The previous row above is a black cell
			if startingPoint.Row == 1 || grid.IsBlackCell(Point{r - 1, c}) {

				// Find the ending point of this down word
				nc.dLen = 0
				thisPoint := startingPoint
				for thisPoint.Row <= n && !grid.IsBlackCell(thisPoint) {
					nc.dLen++
					thisPoint.Row++
				}

				// Word lengths must be >= 3 to qualify. If this one
				// isn't, set its length back to zero.
				if nc.dLen < 3 {
					log.Printf("WARNING: at (%d,%d), down word is invalid\n", r, c)
					nc.dLen = 0
				}

				if nc.dLen > 0 {

					// This is a good down word. We need to set a
					// pointer back to the starting cell in each of its
					// cells (except the starting point itself).
					thisPoint = startingPoint

					for i := 1; i < nc.dLen; i++ {
						thisPoint.Row++
						cell := grid.GetCell(thisPoint).(LetterCell)
						cell.ncDown = &startingPoint
						grid.SetCell(thisPoint, cell)
					}
				}
			}

			// If either across or down words exist (length > 0) then
			// treat this starting point as a numbered cell.
			if nc.aLen > 0 || nc.dLen > 0 {

				// Assign the word number and update the cell in the
				// grid to be a numbered cell.
				seq++
				nc.seq = seq
				grid.SetCell(startingPoint, *nc)
			}
		}
	}
}

// GetPoint returns the location of this cell, for the Cell interface
func (nc NumberedCell) GetPoint() Point {
	return nc.point
}

// String returns a string representation of the structure.
func (nc *NumberedCell) String() string {
	return strings.Join([]string{
		fmt.Sprintf("point:{Row:%d,Col:%d}", nc.point.Row, nc.point.Col),
		fmt.Sprintf("seq:%d", nc.seq),
		fmt.Sprintf("aLen:%d", nc.aLen),
		fmt.Sprintf("dLen:%d", nc.dLen),
		fmt.Sprintf("letter:%q", nc.letter),
	}, ",")
}
