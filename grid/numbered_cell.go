package grid

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// NumberedCell is a letter cell that is the beginning of an across word
// and/or a down word.
type NumberedCell struct {
	LetterCell
	wordNumber int // The word number
	aLen       int // Length of the across word (if any)
	dLen       int // Length of the down word (if any)
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// assignNumberedCells is an internal method that looks at each cell in
// the grid and creates a NumberedCell for each one that starts an
// across and/or a down word.
func (grid *Grid) assignNumberedCells() {

	wordNumber := 0 // Next available word number
	var nc NumberedCell
	var ncAcross, ncDown *Point

	for point := range grid.PointIterator() {

		// Skip black cells
		if grid.IsBlackCell(point) {
			continue
		}

		// Check for word starts: across, down, or both.  We don't know
		// the lengths at this point, only that the word starts here.
		r, c := point.Row, point.Col
		startAcross := point.Col == 1 || grid.IsBlackCell(Point{r, c - 1})
		startDown := point.Row == 1 || grid.IsBlackCell(Point{r - 1, c})

		switch {

		case startAcross && startDown:
			ncAcross = &point
			ncDown = &point
			wordNumber++
			nc = createNumberedCell(wordNumber, point, ncAcross, ncDown)

		case startAcross:
			ncAcross = &point
			ncDown = nil
			wordNumber++
			nc = createNumberedCell(wordNumber, point, ncAcross, ncDown)

		case startDown:
			ncAcross = nil
			ncDown = &point
			wordNumber++
			nc = createNumberedCell(wordNumber, point, ncAcross, ncDown)

		default:
			// Not a word starting point. Just assign a letter cell for
			// this point.
			grid.SetCell(point, LetterCell{point: point})
			continue
		}

		// Only if this was a new NumberedCell
		grid.SetCell(point, nc)
	}
}

// createNumberedCell is an internal function to create a numbered cell
// and find the lengths of its across and down words. Used only in
// assignNumberedCells.
func createNumberedCell(wordNumber int, point Point, ncAcross *Point, ncDown *Point) NumberedCell {
	lc := LetterCell{point: point}
	if ncAcross != nil {
		lc.ncAcross = ncAcross
	}
	if ncDown != nil {
		lc.ncDown = ncDown
	}
	nc := NumberedCell{LetterCell: lc, wordNumber: wordNumber}
	return nc
}

// GetAcrossWordLength returns the length of the across word for this
// numbered cell.
func (grid *Grid) GetAcrossWordLength(nc NumberedCell) int {
	if nc.ncAcross == nil {
		return 0
	}
	n := grid.n
	length := 0
	point := nc.point
	point.Col++
	for point.Col <= n && !grid.IsBlackCell(point) {
		length++
		point.Col++
	}
	return length
}

// GetDownWordLength returns the length of the down word for this
// numbered cell.
func (grid *Grid) GetDownWordLength(nc NumberedCell) int {
	if nc.ncDown == nil {
		return 0
	}
	n := grid.n
	length := 0
	point := nc.point
	point.Row++
	for point.Row <= n && !grid.IsBlackCell(point) {
		length++
		point.Row++
	}
	return length

}

// GetPoint returns the location of this cell (for the Cell interface)
func (nc NumberedCell) GetPoint() Point {
	return nc.point
}

// RenumberCells figures out the numbering of the across and down
// words in this word based on the locations of all the black cells.
func (grid *Grid) RenumberCells() {

	// First, find all the numbered cells.
	grid.assignNumberedCells()

	// Now fill in the lengths for each word
	for cell := range grid.CellIterator() {
		switch cell.(type) {

		case NumberedCell:
			nc := cell.(NumberedCell)

			// Set the across length value in the starting numbered cell
			nc.aLen = grid.GetAcrossWordLength(nc)

			// Copy the pointer into the rest of the across word cells
			if nc.aLen > 0 {
				nc.LetterCell.ncAcross = &nc.point
				for i := 1; i < nc.aLen; i++ {
					cellPoint := cell.GetPoint()
					cellPoint.Col++
					grid.SetNumberedCellAcross(cellPoint, nc)
				}
			}

			// Set the down length value in the starting numbered cell
			nc.dLen = grid.GetDownWordLength(nc)
			if nc.dLen > 0 {
				nc.LetterCell.ncDown = &nc.point
				for i := 1; i < nc.dLen; i++ {
					cellPoint := cell.GetPoint()
					cellPoint.Row++
					grid.SetNumberedCellDown(cellPoint, nc)
				}
			}
		}
	}
}

// SetNumberedCellAcross sets the "ncAcross" pointer in this cell to the
// specified starting word number cell.
func (grid *Grid) SetNumberedCellAcross(point Point, nc NumberedCell) {
	cell := grid.GetCell(point)
	switch cell.(type) {
	case LetterCell:
		lc := cell.(LetterCell)
		lc.ncAcross = &nc.point
		grid.SetCell(point, lc)
	case NumberedCell:
		nc := cell.(NumberedCell)
		nc.ncAcross = &nc.point
		grid.SetCell(point, nc)
	}
}

// SetNumberedCellDown sets the "ncDown" pointer in this cell to the
// specified starting word number cell.
func (grid *Grid) SetNumberedCellDown(point Point, nc NumberedCell) {
	cell := grid.GetCell(point)
	switch cell.(type) {
	case LetterCell:
		lc := cell.(LetterCell)
		lc.ncAcross = &nc.point
		grid.SetCell(point, lc)
	case NumberedCell:
		nc := cell.(NumberedCell)
		nc.ncAcross = &nc.point
		grid.SetCell(point, nc)
	}
}

// String returns a string representation of the structure.
func (nc *NumberedCell) String() string {
	return strings.Join([]string{
		fmt.Sprintf("LetterCell:{%v}", nc.LetterCell.String()),
		fmt.Sprintf("seq:%d", nc.wordNumber),
		fmt.Sprintf("aLen:%d", nc.aLen),
		fmt.Sprintf("dLen:%d", nc.dLen),
	}, ",")
}
