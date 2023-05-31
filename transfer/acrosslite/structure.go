package export

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
