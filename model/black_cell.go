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
			switch cell := cell.(type) {
			case BlackCell:
				out <- cell
			}
		}
	}()
	return out
}

// CountBlackCells returns the number of black cells in the grid
func (grid *Grid) CountBlackCells() int {
	nbr := 0
	for range grid.BlackCellIterator() {
		nbr++
	}
	return nbr
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

// RedoBlackCell pops a point from the redo stack and toggles the black
// cell at that point.
func (grid *Grid) RedoBlackCell() {
	if grid.redoPointStack.IsEmpty() {
		// Nothing to redo
		return
	}

	// Pop the point at the top of the redo stack
	point, _ := grid.redoPointStack.Pop()

	// Push that point onto the undo stack
	grid.undoPointStack.Push(point)

	// Toggle that point and its symmetric twin.
	grid.togglePoint(point)
}

// String returns a string representation of this black cell.
func (bc BlackCell) String() string {
	sb := bc.point.String()
	return sb
}

// Toggle switches a point between black cell and letter cell.
// Does so also to the symmetric point.
func (grid *Grid) Toggle(point Point) {
	if err := grid.ValidIndex(point); err != nil {
		panic(err)
	}
	grid.undoPointStack.Push(point)
	grid.togglePoint(point)
}

// togglePoint is an internal method that changes the specified cell
// from a letter cell to a black cell or vice versa.  Separated here
// because we need to handle the undo/redo stacks differently in
// different cases.
func (grid *Grid) togglePoint(point Point) {
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
}

// UndoBlackCell pops a point from the undo stack and toggles the black
// cell at that point.
func (grid *Grid) UndoBlackCell() {
	if grid.undoPointStack.IsEmpty() {
		// Nothing to undo
		return
	}

	// Pop the point at the top of the undo stack
	point, _ := grid.undoPointStack.Pop()

	// Push that point onto the redo stack
	grid.redoPointStack.Push(point)

	// Toggle that point and its symmetric twin.  Note that this is the
	// same as the Toggle() method except that it doesn't push the
	// action onto the undo stack.
	grid.togglePoint(point)
}
