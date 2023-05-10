package grid

import (
	"fmt"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

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
		case startAcross, startDown:
			wordNumber++
			nc := grid.createNumberedCell(wordNumber, point, startAcross, startDown)
			grid.SetCell(point, *nc)

		default:
			// Not a word starting point. Just assign a letter cell for
			// this point.
			lc := new(LetterCell)
			lc.point = point
			lc.ncAcross = nil
			lc.ncDown = nil
			grid.SetCell(point, *lc)
			continue
		}
	}

	// Now go through the numbered cells.  For each one, go through its
	// across and/or down word cells and set their back pointers
	// (ncAcross and ncDown) to the numbered cell.

	for nc := range grid.NumberedCellIterator() {
		pnc := new(Point)
		pnc.Row = nc.point.Row
		pnc.Col = nc.point.Col
		if nc.aLen > 0 {
			// There is an across word. Skip the initial (numbered) cell.
			for i := 1; i < nc.aLen; i++ {
				point := Point{nc.point.Row, nc.point.Col + i}
				cell := grid.GetCell(point)
				switch cell.(type) {
				case LetterCell:
					cellLc := cell.(LetterCell)
					cellLc.ncAcross = pnc
					grid.SetCell(cellLc.GetPoint(), cellLc)
				case NumberedCell:
					cellNc := cell.(NumberedCell)
					cellNc.ncAcross = pnc
					grid.SetCell(cellNc.GetPoint(), cellNc)
				}
			}
		}
		if nc.dLen > 0 {
			// There is a down word. Skip the initial (numbered) cell.
			for i := 1; i < nc.dLen; i++ {
				point := Point{nc.point.Row + i, nc.point.Col}
				cell := grid.GetCell(point)
				switch cell.(type) {
				case LetterCell:
					cellLc := cell.(LetterCell)
					cellLc.ncDown = pnc
					grid.SetCell(cellLc.GetPoint(), cellLc)
				case NumberedCell:
					cellNc := cell.(NumberedCell)
					cellNc.ncDown = pnc
					grid.SetCell(cellNc.GetPoint(), cellNc)
				}
			}
		}
	}
}

// createNumberedCell is an internal function to create a numbered cell
// and find the lengths of its across and down words. Used only in
// assignNumberedCells.
func (grid *Grid) createNumberedCell(wordNumber int, point Point, startAcross bool, startDown bool) *NumberedCell {

	var ncAcross *Point
	var ncDown *Point

	if startAcross {
		ncAcross = new(Point)
		ncAcross.Row = point.Row
		ncAcross.Col = point.Col
	}

	if startDown {
		ncDown = new(Point)
		ncDown.Row = point.Row
		ncDown.Col = point.Col
	}

	nc := new(NumberedCell)
	nc.LetterCell.point = point
	nc.LetterCell.ncAcross = ncAcross
	nc.LetterCell.ncDown = ncDown
	nc.wordNumber = wordNumber
	if startAcross {
		nc.aLen = grid.GetAcrossWordLength(ncAcross)
	}
	if startDown {
		nc.dLen = grid.GetDownWordLength(ncDown)
	}
	return nc
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
	for nc := range grid.NumberedCellIterator() {
		pnc := &nc

		// Set the across length value in the starting numbered cell
		if pnc.aLen > 0 {
			pnc.LetterCell.ncAcross = &nc.point
			for i := 1; i < pnc.aLen; i++ {
				cellPoint := pnc.GetPoint()
				cellPoint.Col++
				grid.SetNumberedCellAcross(cellPoint, pnc)
			}
		}

		// Set the down length value in the starting numbered cell
		if pnc.dLen > 0 {
			pnc.LetterCell.ncDown = &nc.point
			for i := 1; i < pnc.dLen; i++ {
				cellPoint := nc.GetPoint()
				cellPoint.Row++
				grid.SetNumberedCellDown(cellPoint, pnc)
			}
		}
	}
}

// SetNumberedCellAcross sets the "ncAcross" pointer in this cell to the
// specified starting word number cell.
func (grid *Grid) SetNumberedCellAcross(point Point, nc *NumberedCell) {
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
func (grid *Grid) SetNumberedCellDown(point Point, nc *NumberedCell) {
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
func (nc NumberedCell) String() string {
	return strings.Join([]string{
		fmt.Sprintf("LetterCell:{%v}", nc.LetterCell.String()),
		fmt.Sprintf("seq:%d", nc.wordNumber),
		fmt.Sprintf("aLen:%d", nc.aLen),
		fmt.Sprintf("dLen:%d", nc.dLen),
	}, ",")
}
