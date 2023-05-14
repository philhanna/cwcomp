package model

import (
	"fmt"
	"testing"
)

func Test_dumpGrid(t *testing.T) {
	grid := getGoodGrid()
	dumpGrid(grid)
	fmt.Println(grid.String())
}
