package importer

import (
	"testing"

	al "github.com/philhanna/cwcomp/acrosslite"
	"github.com/stretchr/testify/assert"
)

func TestHandleReadingGrid(t *testing.T) {
	pal := new(al.AcrossLite)
	pal.Grid = []string{
		"AAAAA",
		"BBBBB",
		"CCCCC",
		"DDDDD",
	}
	pal.Size = 5
	_, err := HandleReadingGrid(pal, "<ACROSS>")
	assert.NotNil(t, err)

	pal.Grid = []string{
		"AAAAA",
		"BBBBB",
		"CCCCC",
	}
	// Ensure that lines are padded
	_, err = HandleReadingGrid(pal, "DDDD")
	assert.Nil(t, err)
	// But not too long
	_, err = HandleReadingGrid(pal, "DDDDDDD")
	assert.NotNil(t, err)

}
