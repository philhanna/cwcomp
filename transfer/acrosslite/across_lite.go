package export

import "fmt"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// AcrossLite is a representation of a puzzle in a standard interchange
// format.  This is proprietary format is defined and maintained by
// https://www.litsoft.com/.  It is described in
// https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
type AcrossLite struct {
	Title       string
	Author      string
	Copyright   string
	Size        int
	Grid        []string
	AcrossClues []string
	DownClues   []string
	Notepad     []string
}

// ---------------------------------------------------------------------
// Implementation of Importer interface
// ---------------------------------------------------------------------

// GetSize returns the number of rows or columns in this puzzle
func (self *AcrossLite) GetSize() int {
	return self.Size
}

// GetName returns the puzzle name, which will be used as part of the
// key in the database representation.
//
// This is not the same as the puzzle title
func (self *AcrossLite) GetName() string {
	return "TODO" // TODO figure out a way to supply this
}

// GetTitle returns the puzzle title, which is a descriptive string that
// is typically used as the heading of the page it is printed on in the
// newspaper.
func (self *AcrossLite) GetTitle() string {
	return self.Title
}

// GetCell returns the letter at a given point in the grid.  These are
// relative to 1, not 0, so
//
//	r = 1, 2, ..., n c = 1, 2, ..., n
//
// If the letter value is '\x00', it is a black cell.  Otherwise, it is
// converted to uppercase.  If the letter is not a black cell and not in
// the alphabet A-Z, an error is returned.
func (self *AcrossLite) GetCell(r, c int) (byte, error) {
	n := self.GetSize()
	if r < 1 || r > n || c < 1 || c > n {
		return 0, fmt.Errorf("Invalid index: r=%d,c=%d", r, c)
	}
	i, j := r-1, c-1
	return self.Grid[i][j], nil
}

// GetAcrossClues returns the clues to the across words.  The slice
// index is one less than the word number.  Words with no clue yet still
// occupy a position in the slice containing the empty string.
func (self *AcrossLite) GetAcrossClues() []string {
	return self.AcrossClues
}

// GetDownClues returns the clues to the down words.  The slice index is
// one less than the word number.  Words with no clue yet still occupy a
// position in the slice containing the empty string.
func (self *AcrossLite) GetDownClues() []string {
	return self.DownClues
}
