package grid

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// A word number in the grid
type NumberedCell struct {
	point Point // Starting point of numbered cell
	seq   int   // Word number (1, 2, ...)
	alen  int   // Length of across word starting at this point
	dlen  int   // Length of down word starting at this point
}
