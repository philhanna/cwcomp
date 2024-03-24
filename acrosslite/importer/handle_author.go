package importer

import (
	al "github.com/philhanna/cwcomp/acrosslite"
)

// HandleLookingForAuthor looks at the next line in the file and ensures
// that it is <AUTHOR>.
func HandleLookingForAuthor(pal *al.AcrossLite, line string) (ParsingState, error) {
	switch line {
	case "<AUTHOR>":
		return READING_AUTHOR, nil
	default:
		return UNKNOWN, errNoAuthor
	}
}

// HandleReadingAuthor copies the line into the Author element of the
// AcrossLite structure.
func HandleReadingAuthor(pal *al.AcrossLite, line string) (ParsingState, error) {
	pal.Author = line
	return LOOKING_FOR_COPYRIGHT, nil
}
