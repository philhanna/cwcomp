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

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetAcrossClues returns a map of across word numbers to their clues.
func (pal *AcrossLite) GetAcrossClues() map[int]string {
	return pal.AcrossClues
}

// GetAuthor returns the author line
func (pal *AcrossLite) GetAuthor() string {
	return pal.Author
}

// GetCell returns the letter at a given point in the grid.  These are
// relative to 1, not 0, so
//
//	r = 1, 2, ..., n and c = 1, 2, ..., n
//
// If the letter value is '\x00', it is a black cell.  Otherwise, it is
// converted to uppercase.  If the letter is not a black cell and not in
// the alphabet A-Z, an error is returned.
func (pal *AcrossLite) GetCell(r, c int) (byte, error) {
	n := pal.GetSize()
	if n < 1 {
		return 0, fmt.Errorf("puzzle size has not yet been set")
	}
	if r < 1 || r > n || c < 1 || c > n {
		return 0, fmt.Errorf("invalid index: r=%d,c=%d", r, c)
	}
	i, j := r-1, c-1

	letter := pal.Grid[i][j]
	if letter == byte('.') {
		letter = byte('\x00')
	}

	return letter, nil
}

// GetCopyright returns the copyright line
func (pal *AcrossLite) GetCopyright() string {
	return pal.Copyright
}

// GetCreatedDate returns the creation date timestamp. If one has not
// been specified, uses current date/time.
func (pal *AcrossLite) GetCreatedDate() time.Time {
	if pal.CreatedDate.IsZero() {
		pal.CreatedDate = time.Now()
	}
	return pal.CreatedDate
}

// GetDownClues returns a map of down word numbers to their clues.
func (pal *AcrossLite) GetDownClues() map[int]string {
	return pal.DownClues
}

// GetGrid returns the list of strings in the grid
func (pal *AcrossLite) GetGrid() []string {
	return pal.Grid
}

// GetModifiedDate returns the modified date timestamp. If one has not
// been specified, uses current date/time.
func (pal *AcrossLite) GetModifiedDate() time.Time {
	if pal.ModifiedDate.IsZero() {
		pal.ModifiedDate = time.Now().Add(3 * 24 * time.Hour)
	}
	return pal.ModifiedDate
}

// GetName returns the puzzle name, which will be used as part of the
// key in the database representation.
//
// This is not the same as the puzzle title
func (pal *AcrossLite) GetName() string {
	return pal.Name
}

// GetNotepad returns the <NOTEPAD> entry, which may be empty
func (pal *AcrossLite) GetNotepad() string {
	parts := make([]string, 2)
	parts[0] = fmt.Sprintf("%q:%q", "created", pal.GetCreatedDate().Format(ISO8601))
	parts[1] = fmt.Sprintf("%q:%q", "modified", pal.GetModifiedDate().Format(ISO8601))
	return "{" + strings.Join(parts, ",") + "}"
}

// GetSize returns the number of rows or columns in this puzzle
func (pal *AcrossLite) GetSize() int {
	return pal.Size
}

// GetTitle returns the puzzle title, which is a descriptive string that
// is typically used as the heading of the page it is printed on in the
// newspaper.
func (pal *AcrossLite) GetTitle() string {
	return pal.Title
}

// ---------------------------------------------------------------------
// Implementation of Exporter interface
// ---------------------------------------------------------------------

// SetAcrossClues sets the across clue map
func (pal *AcrossLite) SetAcrossClues(clueMap map[int]string) {
	pal.AcrossClues = clueMap
}

// SetAuthor sets the author line
func (pal *AcrossLite) SetAuthor(author string) {
	pal.Author = author
}

// SetCell sets the letter at a given point in the grid.  These are
// relative to 1, not 0, so
//
//	r = 1, 2, ..., n and c = 1, 2, ..., n
//
// If the letter value is '\x00', it is a black cell, which must be
// represented by '.' in this struct element, according to the
// AcrossLite format.
func (pal *AcrossLite) SetCell(r, c int, letter byte) error {

	// Size must have already been parsed
	n := pal.GetSize()
	if n < 1 {
		return fmt.Errorf("puzzle size has not yet been set")
	}
	if r < 1 || r > n || c < 1 || c > n {
		return fmt.Errorf("invalid index: r=%d,c=%d", r, c)
	}

	// Convert row and column to zero-based coordinates
	i, j := r-1, c-1

	// Convert \x00 to '.' inside this struct element.
	if letter == '\x00' {
		letter = '.'
	}

	// Replace the cell in the string[i] at position j
	sb := strings.Builder{}
	for k, sLetter := range pal.Grid[i] {
		if k == j {
			// This is the one we want to replace
			sb.WriteRune(rune(letter))
		} else {
			// Copy the rest unaltered
			sb.WriteRune(sLetter)
		}
	}
	// Set the resultng string back in the struct element
	pal.Grid[i] = sb.String()

	return nil
}

// SetCopyright sets the copyright line
func (pal *AcrossLite) SetCopyright(copyright string) {
	pal.Copyright = copyright
}

// SetCreatedDate sets the creation datetime
func (pal *AcrossLite) SetCreatedDate(created time.Time) {
	pal.CreatedDate = created
}

// SetDownClues sets the down clue map
func (pal *AcrossLite) SetDownClues(clueMap map[int]string) {
	pal.DownClues = clueMap
}

// SetModifiedDate sets the modified datetime
func (pal *AcrossLite) SetModifiedDate(modified time.Time) {
	pal.ModifiedDate = modified
}

// SetName sets the puzzle nme
func (pal *AcrossLite) SetName(name string) {
	pal.Name = name
}

// SetSize sets the number of rows or columns in this puzzle
func (pal *AcrossLite) SetSize(n int) {
	pal.Size = n
	pal.Grid = make([]string, n)
	for i := 0; i < n; i++ {
		pal.Grid[i] = strings.Repeat(" ", n)
	}
}

// SetTitle sets the puzzle title
func (pal *AcrossLite) SetTitle(title string) {
	pal.Title = title
}
