package grid

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// BlackCell is a point in the grid that can have no letters. It marks
// the boundaries for the starting and stopping point of words.
type BlackCell struct {
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// AddBlack cell sets the cell at the specifed point and at the
// symmetric point to be black cells.
// TODO: push this cell onto the undo stack
func (grid *Grid) AddBlackCell(point Point) {

	cell := new(BlackCell)
	symCell := new(BlackCell)

	grid.SetCell(point, *cell)
	grid.SetCell(grid.SymmetricPoint(point), *symCell)
}

// IsBlackCell returns true if the specified point is a black cell.
func (g *Grid) IsBlackCell(point Point) bool {
	cell := g.GetCell(point)
	switch cell.(type) {
	case BlackCell:
		return true
	default:
		return false
	}
}

// String returns a string representation of this black cell
func (bc *BlackCell) String() string {
	sb := "bc"
	return sb
}

// SymmetricPoint returns the point of the cell at 180 degrees rotation.
func (grid *Grid) SymmetricPoint(point Point) Point {
	r := point.Row
	c := point.Col
	return Point{grid.n + 1 - r, grid.n + 1 - c}
}
