package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_GetGridList(t *testing.T) {
	userid := 1
	grid := getGoodGrid()
	gridNames, err := grid.GetGridList(userid)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(gridNames))
}
