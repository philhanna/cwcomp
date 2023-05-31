package export

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// AcrossLite is a representation of a puzzle in a standard interchange
// format.  This is proprietary format is defined and maintained by
// https://www.litsoft.com/.  It is described in
// https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
type AcrossLite struct {
	title       string
	author      string
	copyright   string
	size        int
	grid        []string
	acrossClues []string
	downClues   []string
	notepad     []string
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewAcrossLite creates a new, empty AcrossLite structure and returns a
// pointer to it.
func NewAcrossLite() *AcrossLite {
	p := new(AcrossLite)

	p.title = ""
	p.author = ""
	p.copyright = ""
	p.size = 0
	p.grid = make([]string, 0)
	p.acrossClues = make([]string, 0)
	p.downClues = make([]string, 0)
	p.notepad = make([]string, 0)

	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Import parses the text in an AcrossLite puzzle
func Import(filename string) (*AcrossLite, error) {
	pal := NewAcrossLite()

	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	type ParsingState byte
	const (
		INIT ParsingState = iota
		LOOKING_FOR_TITLE
		READING_TITLE
		LOOKING_FOR_AUTHOR
		READING_AUTHOR
		LOOKING_FOR_COPYRIGHT
		READING_COPYRIGHT
		LOOKING_FOR_SIZE
		READING_SIZE
	)
	state := INIT

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {

		line := scanner.Text()
		line = strings.TrimSpace(line)

		switch state {

		case INIT:
			if line == "<ACROSS PUZZLE>" {
				state = LOOKING_FOR_TITLE
			} else {
				return nil, fmt.Errorf("Not a valid AcrossLite text file: %s", line)
			}

		case LOOKING_FOR_TITLE:
			if line == "<TITLE>" {
				state = READING_TITLE
			} else {
				return nil, fmt.Errorf("did not find <TITLE>")
			}

		case READING_TITLE:
			pal.title = line
			state = LOOKING_FOR_AUTHOR

		case LOOKING_FOR_AUTHOR:
			if line == "<AUTHOR>" {
				state = READING_AUTHOR
			} else {
				return nil, fmt.Errorf("did not find <AUTHOR>")
			}

		case READING_AUTHOR:
			pal.author = line
			state = LOOKING_FOR_COPYRIGHT

		case LOOKING_FOR_COPYRIGHT:
			if line == "<COPYRIGHT>" {
				state = READING_COPYRIGHT
			} else {
				return nil, fmt.Errorf("did not find <COPYRIGHT>")
			}

		case READING_COPYRIGHT:
			pal.copyright = line
			state = LOOKING_FOR_SIZE

		case LOOKING_FOR_SIZE:
			if line == "<SIZE>" {
				state = READING_SIZE
			} else {
				return nil, fmt.Errorf("did not find <SIZE>")
			}

		}
	}

	return pal, nil
}
