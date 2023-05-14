// A web application that is used to create crossword puzzles.
package cwcomp

import (
	"log"
	"path/filepath"
	"runtime"
)

// Initialization function
func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// GetPackageName returns the name of the package
func GetPackageName() string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	pkgname := filepath.Base(dir)

	return pkgname
}
