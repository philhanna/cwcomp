package model

import (
	"testing"
)

func Test_dumpGrid(t *testing.T) {
	grid := getGoodGrid()
	dumpGrid(grid)
}
