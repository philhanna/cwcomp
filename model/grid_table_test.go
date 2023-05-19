package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const TEST_USERID = 1

func TestGrid_SaveGrid(t *testing.T) {

	// Create a new grid and populate it with words
	grid := getGoodGrid()
	type test struct {
		seq  int
		dir  Direction
		text string
	}
	testWords := []test{
		{1, ACROSS, "NOW"},
		{7, DOWN, "COW"},
		{8, ACROSS, "BLUE"},
		{20, DOWN, "HOW"},
	}
	for _, test := range testWords {
		word := grid.LookupWordByNumber(test.seq, test.dir)
		grid.SetText(word, test.text)
	}
	grid.SetGridName("Rhyme")
	_, err := grid.SaveGrid(TEST_USERID)
	assert.Nil(t, err)

	// Done with the grid
	grid.DeleteGrid(TEST_USERID, "Rhyme")
}

func TestGrid_GetGridList(t *testing.T) {
	grid := getGoodGrid()
	userid := TEST_USERID
	gridNames, err := grid.GetGridList(userid)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(gridNames))
}
