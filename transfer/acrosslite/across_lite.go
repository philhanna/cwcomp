package export

import (
	"fmt"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// AcrossLite is a representation of a puzzle in a standard interchange
// format.  This is a proprietary format is defined and maintained by
// https://www.litsoft.com/.  It is described in
// https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
type AcrossLite struct {
	Size        int
	Name        string
	Title       string
	Author      string
	Copyright   string
	Grid        []string
	AcrossClues map[int]string
	DownClues   map[int]string
	Notepad     string
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewAcrossLite creates and initialized an AcrossLite structure and
// returns a pointer to it.
func NewAcrossLite() *AcrossLite {
	pal := new(AcrossLite)
	pal.Grid = make([]string, 0)
	pal.AcrossClues = make(map[int]string)
	pal.DownClues = make(map[int]string)
	return pal
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
	return self.Name
}

func (self *AcrossLite) SetName(name string) {
	self.Name = name
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

	letter := self.Grid[i][j]
	if letter == byte('.') {
		letter = byte('\x00')
	}

	return letter, nil
}

// GetAcrossClues returns a map of across word numbers to their clues.
func (self *AcrossLite) GetAcrossClues() map[int]string {
	return self.AcrossClues
}

// GetAcrossClues returns a map of down word numbers to their clues.
func (self *AcrossLite) GetDownClues() map[int]string {
	return self.DownClues
}
