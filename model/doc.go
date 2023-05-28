// Package model contains types and functions that support the internal
// workings of the application.
//
// A puzzle consists of an nxn matrix of cells, which are of two types:
//
//   - Black cells: Blocks in the grid
//
//   - Letter cells: Ordinary cells where letters of words can be
//     placed.
//
// The puzzle also supports undo/redo for black cells in the grid.
package model

import "os"

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
