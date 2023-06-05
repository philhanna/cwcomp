// A web application that is used to create crossword puzzles.
//
// Contains types and functions that support the internal workings of
// the application.
//
// A puzzle consists of an nxn matrix of cells, which are of two types:
//
//   - Black cells: Blocks in the grid
//
//   - Letter cells: Ordinary cells where letters of words can be
//     placed.
//
// The puzzle also supports undo/redo for black cells and words in the
// grid.
//
// The word list is derived from git@github.com:elasticdog/yawl.git,
// with some editing by me.
package cwcomp

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Initialization function
func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

const BLACK_CELL = '\x00'

// GetPackageName returns the name of the package
func GetPackageName() string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	pkgname := filepath.Base(dir)

	return pkgname
}

// Returns true if the specified file exists
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
