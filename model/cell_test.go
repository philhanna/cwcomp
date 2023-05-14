package model

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	PACKAGE_NAME = "model"
	PREFIX       = PACKAGE_NAME + "."
)

func TestCell_GetPoint(t *testing.T) {

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

	const n = 9

	grid := getGoodGrid()

	verify := func(t *testing.T, point Point, name string) {
		cell := grid.GetCell(point)
		fullType := fmt.Sprintf("%T", cell)
		shortType := strings.TrimPrefix(fullType, PREFIX)
		assert.Equal(t, name, shortType)
		x, y := point.ToXY()
		assert.Equal(t, abbrev[y][x], name[0])
	}

	for actualCell := range grid.CellIterator() {
		switch actualCell.(type) {
		case BlackCell:
			point := actualCell.(BlackCell).GetPoint()
			verify(t, point, "BlackCell")
		case LetterCell:
			point := actualCell.(LetterCell).GetPoint()
			verify(t, point, "LetterCell")
		default:
			t.Fatalf("Unknown type at %v\n", actualCell)
		}
	}
}