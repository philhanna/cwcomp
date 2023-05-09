package grid

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

// AddBlack cell sets the cells at the specifed point and at the
// symmetric point to be black cells.
func (grid *Grid) AddBlackCell(point Point) {
	cell := NewBlackCell(point)
	grid.SetCell(point, cell)

	symPoint := grid.SymmetricPoint(point)
	symCell := BlackCell{point: symPoint}
	grid.SetCell(grid.SymmetricPoint(point), symCell)

	grid.RenumberCells()
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

// RemoveBlackCell sets the cells at the specified point and at the
// symmetric point to be letter cells.
func (grid *Grid) RemoveBlackCell(point Point) {
	cell := NewLetterCell(point)
	grid.SetCell(point, cell)

	symPoint := grid.SymmetricPoint(point)
	symCell := NewLetterCell(symPoint)
	grid.SetCell(grid.SymmetricPoint(point), symCell)

	grid.RenumberCells()

}

// String returns a string representation of this black cell.
func (bc *BlackCell) String() string {
	sb := bc.point.String()
	return sb
}

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.Row
	c := point.Col
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}
