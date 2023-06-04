package acrosslite

import (
	"fmt"
	"strings"
	"time"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// AcrossLite is a representation of a puzzle in a standard interchange
// format.  This is a proprietary format is defined and maintained by
// https://www.litsoft.com/.  It is described in
// https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
type AcrossLite struct {
	Size         int
	Name         string
	Title        string
	Author       string
	Copyright    string
	Grid         []string
	AcrossClues  map[int]string
	DownClues    map[int]string
	CreatedDate  time.Time
	ModifiedDate time.Time
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------
const ISO8601 = "2006-01-02T15:04:05.999999"

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

// GetTitle returns the puzzle title, which is a descriptive string that
// is typically used as the heading of the page it is printed on in the
// newspaper.
func (self *AcrossLite) GetTitle() string {
	return self.Title
}

// GetAuthor returns the author line
func (self *AcrossLite) GetAuthor() string {
	return self.Author
}

// GetCopyright returns the copyright line
func (self *AcrossLite) GetCopyright() string {
	return self.Copyright
}

// GetGrid returns the list of strings in the grid
func (self *AcrossLite) GetGrid() []string {
	return self.Grid
}

// GetCell returns the letter at a given point in the grid.  These are
// relative to 1, not 0, so
//
//	r = 1, 2, ..., n and c = 1, 2, ..., n
//
// If the letter value is '\x00', it is a black cell.  Otherwise, it is
// converted to uppercase.  If the letter is not a black cell and not in
// the alphabet A-Z, an error is returned.
func (self *AcrossLite) GetCell(r, c int) (byte, error) {
	n := self.GetSize()
	if n < 1 {
		return 0, fmt.Errorf("Puzzle size has not yet been set")
	}
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

// GetCreatedDate returns the creation date timestamp. If one has not
// been specified, uses current date/time.
func (self *AcrossLite) GetCreatedDate() time.Time {
	if self.CreatedDate.IsZero() {
		self.CreatedDate = time.Now()
	}
	return self.CreatedDate
}

// GetModifiedDate returns the modified date timestamp. If one has not
// been specified, uses current date/time.
func (self *AcrossLite) GetModifiedDate() time.Time {
	if self.ModifiedDate.IsZero() {
		self.ModifiedDate = time.Now().Add(3 * 24 * time.Hour)
	}
	return self.ModifiedDate
}

// GetNotepad returns the <NOTEPAD> entry, which may be empty
func (self *AcrossLite) GetNotepad() string {
	parts := make([]string, 2)
	parts[0] = fmt.Sprintf("%q:%q", "created", self.GetCreatedDate().Format(ISO8601))
	parts[1] = fmt.Sprintf("%q:%q", "modified", self.GetModifiedDate().Format(ISO8601))
	return "{" + strings.Join(parts, ",") + "}"
}

// ---------------------------------------------------------------------
// Implementation of Exporter interface
// ---------------------------------------------------------------------

// SetSize sets the number of rows or columns in this puzzle
func (self *AcrossLite) SetSize(n int) {
	self.Size = n
	self.Grid = make([]string, n)
	for i := 0; i < n; i++ {
		self.Grid[i] = strings.Repeat(" ", n)
	}
}

// SetName sets the puzzle nme
func (self *AcrossLite) SetName(name string) {
	self.Name = name
}

// SetTitle sets the puzzle title
func (self *AcrossLite) SetTitle(title string) {
	self.Title = title
}

// SetAuthor sets the author line
func (self *AcrossLite) SetAuthor(author string) {
	self.Author = author
}

// SetCopyright sets the copyright line
func (self *AcrossLite) SetCopyright(copyright string) {
	self.Copyright = copyright
}

// SetCell sets the letter at a given point in the grid.  These are
// relative to 1, not 0, so
//
//	r = 1, 2, ..., n and c = 1, 2, ..., n
//
// If the letter value is '\x00', it is a black cell, which must be
// represented by '.' in this struct element, according to the
// AcrossLite format.
func (self *AcrossLite) SetCell(r, c int, letter byte) error {

	// Size must have already been parsed
	n := self.GetSize()
	if n < 1 {
		return fmt.Errorf("Puzzle size has not yet been set")
	}
	if r < 1 || r > n || c < 1 || c > n {
		return fmt.Errorf("Invalid index: r=%d,c=%d", r, c)
	}

	// Convert row and column to zero-based coordinates
	i, j := r-1, c-1

	// Convert \x00 to '.' inside this struct element.
	if letter == '\x00' {
		letter = '.'
	}

	// Replace the cell in the string[i] at position j
	sb := strings.Builder{}
	for k, sLetter := range self.Grid[i] {
		if k == j {
			// This is the one we want to replace
			sb.WriteRune(rune(letter))
		} else {
			// Copy the rest unaltered
			sb.WriteRune(sLetter)
		}
	}
	// Set the resultng string back in the struct element
	self.Grid[i] = sb.String()

	return nil
}

// SetAcrossClues sets the across clue map
func (self *AcrossLite) SetAcrossClues(clueMap map[int]string) {
	self.AcrossClues = clueMap
}

// SetDownClues sets the down clue map
func (self *AcrossLite) SetDownClues(clueMap map[int]string) {
	self.DownClues = clueMap
}

// SetCreatedDate sets the creation datetime
func (self *AcrossLite) SetCreatedDate(created time.Time) {
	self.CreatedDate = created
}

// SetModifiedDate sets the modified datetime
func (self *AcrossLite) SetModifiedDate(modified time.Time) {
	self.ModifiedDate = modified
}