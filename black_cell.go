package cwcomp

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
func (puzzle *Puzzle) BlackCellIterator() <-chan BlackCell {
	out := make(chan BlackCell)
	go func() {
		defer close(out)
		for point := range puzzle.PointIterator() {
			cell := puzzle.GetCell(point)
			switch cell := cell.(type) {
			case BlackCell:
				out <- cell
			}
		}
	}()
	return out
}

// CountBlackCells returns the number of black cells in the grid
func (puzzle *Puzzle) CountBlackCells() int {
	nbr := 0
	for range puzzle.BlackCellIterator() {
		nbr++
	}
	return nbr
}

// GetPoint returns the location of this cell (for the Cell interface).
func (bc BlackCell) GetPoint() Point {
	return bc.point
}

// IsBlackCell returns true if the specified point is a black cell.
func (puzzle *Puzzle) IsBlackCell(point Point) bool {
	cell := puzzle.GetCell(point)
	switch cell.(type) {
	case BlackCell:
		return true
	default:
		return false
	}
}

// RedoBlackCell pops a point from the redo stack and toggles the black
// cell at that point.
func (puzzle *Puzzle) RedoBlackCell() {
	if puzzle.redoPointStack.IsEmpty() {
		// Nothing to redo
		return
	}

	// Pop the point at the top of the redo stack
	point, _ := puzzle.redoPointStack.Pop()

	// Push that point onto the undo stack
	puzzle.undoPointStack.Push(point)

	// Toggle that point and its symmetric twin.
	puzzle.togglePoint(point)
}

// String returns a string representation of this black cell.
func (bc BlackCell) String() string {
	sb := bc.point.String()
	return sb
}

// Toggle switches a point between black cell and letter cell.
// Does so also to the symmetric point.
func (puzzle *Puzzle) Toggle(point Point) {
	if err := puzzle.ValidIndex(point); err != nil {
		panic(err)
	}
	puzzle.undoPointStack.Push(point)
	puzzle.togglePoint(point)
}

// togglePoint is an internal method that changes the specified cell
// from a letter cell to a black cell or vice versa.  Separated here
// because we need to handle the undo/redo stacks differently in
// different cases.
func (puzzle *Puzzle) togglePoint(point Point) {
	cell := puzzle.GetCell(point)

	switch cell.(type) {

	case BlackCell:
		cell = NewLetterCell(point)
		puzzle.SetCell(point, cell)
		symPoint := puzzle.SymmetricPoint(point)
		symCell := NewLetterCell(symPoint)
		puzzle.SetCell(symPoint, symCell)

	case LetterCell:
		cell = NewBlackCell(point)
		puzzle.SetCell(point, cell)
		symPoint := puzzle.SymmetricPoint(point)
		symCell := NewBlackCell(symPoint)
		puzzle.SetCell(symPoint, symCell)
	}
}

// UndoBlackCell pops a point from the undo stack and toggles the black
// cell at that point.
func (puzzle *Puzzle) UndoBlackCell() {
	if puzzle.undoPointStack.IsEmpty() {
		// Nothing to undo
		return
	}

	// Pop the point at the top of the undo stack
	point, _ := puzzle.undoPointStack.Pop()

	// Push that point onto the redo stack
	puzzle.redoPointStack.Push(point)

	// Toggle that point and its symmetric twin.  Note that this is the
	// same as the Toggle() method except that it doesn't push the
	// action onto the undo stack.
	puzzle.togglePoint(point)
}
