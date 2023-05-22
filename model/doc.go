// Package model contains types and functions that support the internal
// workings of the application.
//
// A grid consists of an nxn matrix of cells, which are of two types:
//
//   - Black cells: Blocks in the grid
//
//   - Letter cells: Ordinary cells where letters of words can be
//     placed.  They also contain a pair of pointers: one to the word
//     number that begins the across word that contains this cell, and
//     another to the word number that begins the down word that contains
//     this cell.
//
// The grid also supports undo/redo for black cells in this grid.
package model

import "os"

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
