package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/philhanna/cwcomp"
	"github.com/stretchr/testify/assert"
)

func TestCell_GetPoint(t *testing.T) {

	var (
		PACKAGE_NAME = cwcomp.GetPackageName()
		PREFIX       = PACKAGE_NAME + "."
	)

	// Map out the cell types (abbreviations).

	abbrev := []string{
		"BLLLBLLLL",
		"LLLLBLLLL",
		"LLLLBLLLL",
		"LLLLLLLLB",
		"BBBLLLBBB",
		"BLLLLLLLL",
		"LLLLBLLLL",
		"LLLLBLLLL",
		"LLLLBLLLB",
	}

	puzzle := getGoodPuzzle()

	verify := func(t *testing.T, point Point, name string) {
		cell := puzzle.GetCell(point)
		fullType := fmt.Sprintf("%T", cell)
		shortType := strings.TrimPrefix(fullType, PREFIX)
		assert.Equal(t, name, shortType)
		x, y := point.ToXY()
		assert.Equal(t, abbrev[y][x], name[0])
	}

	for actualCell := range puzzle.CellIterator() {
		switch cell := actualCell.(type) {
		case BlackCell:
			verify(t, cell.point, "BlackCell")
		case LetterCell:
			verify(t, cell.point, "LetterCell")
		default:
			t.Fatalf("Unknown type at %v\n", actualCell)
		}
	}
}
