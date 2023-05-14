package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_LetterCellIterator(t *testing.T) {
	grid := getGoodGrid()
	nlc := 0
	for range grid.LetterCellIterator() {
		nlc++
	}
	assert.Equal(t, 9*9-16, nlc)
}
