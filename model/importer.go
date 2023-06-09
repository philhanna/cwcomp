package model

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Importer is an interface that must be implemented by any source of
// puzzle data that can be imported (e.g., AcrossLite)
type Importer interface {
	// Returns the number of rows or columns in this puzzle
	GetSize() int

	// Returns the puzzle name, which will be used as part of the key in
	// the database representation.
	//
	// This is not the same as the puzzle title
	GetName() string

	// Returns the puzzle title, which is a descriptive string that is
	// typically used as the heading of the page it is printed on in the
	// newspaper.
	GetTitle() string

	// Returns the letter at a given point in the grid.  These are
	// relative to 1, not 0, so
	//
	//  r = 1, 2, ..., n c = 1, 2, ..., n
	//
	// If the letter value is '\x00', it is a black cell.  Otherwise, it is
	// converted to uppercase.  If the letter is not a black cell and
	// not in the alphabet A-Z, an error is returned.
	GetCell(r, c int) (byte, error)

	// Returns a map of across word numbers to clues
	GetAcrossClues() map[int]string

	// Returns a map of down word numbers to clues
	GetDownClues() map[int]string
}
