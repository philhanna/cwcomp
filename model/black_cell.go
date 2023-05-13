package model

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// BlackCell is a point in the grid that can have no letters. It marks
// the boundaries for the starting and stopping point of words.
type BlackCell struct {
	point Point // Location of this cell
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewBlackCell creates a new BlackCell at the specified location.
func NewBlackCell(point Point) BlackCell {
	return BlackCell{point}
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// BlackCellIterator is a generator for all the black cells in the grid.
func (grid *Grid) BlackCellIterator() <-chan BlackCell {
	out := make(chan BlackCell)
	go func() {
		defer close(out)
		for point := range grid.PointIterator() {
			cell := grid.GetCell(point)
			switch cell.(type) {
			case BlackCell:
				bc := cell.(BlackCell)
				out <- bc
			}
		}
	}()
	return out
}

// GetPoint returns the location of this cell (for the Cell interface).
func (bc BlackCell) GetPoint() Point {
	return bc.point
}

// IsBlackCell returns true if the specified point is a black cell.
func (grid *Grid) IsBlackCell(point Point) bool {
	cell := grid.GetCell(point)
	switch cell.(type) {
	case BlackCell:
		return true
	default:
		return false
	}
}

// String returns a string representation of this black cell.
func (bc BlackCell) String() string {
	sb := bc.point.String()
	return sb
}

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.Row
	c := point.Col
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}

// Toggle switches a point between black cell and letter cell.
// Does so also to the symmetric point.
func (grid *Grid) Toggle(point Point) {

	if err := grid.ValidIndex(point); err != nil {
		panic(err)
	}
	grid.undoStack.Push(point)

	cell := grid.GetCell(point)

	switch cell.(type) {

	case BlackCell:
		cell = NewLetterCell(point)
		grid.SetCell(point, cell)
		symPoint := grid.SymmetricPoint(point)
		symCell := NewLetterCell(symPoint)
		grid.SetCell(symPoint, symCell)

	case LetterCell:
		cell = NewBlackCell(point)
		grid.SetCell(point, cell)
		symPoint := grid.SymmetricPoint(point)
		symCell := NewBlackCell(symPoint)
		grid.SetCell(symPoint, symCell)
	}

	// Renumber the cells
	grid.RenumberCells()
}
