// Model is the internal component of the application, used by the
// controller and the view.
package model

import "github.com/philhanna/cwcomp"

const (
	BLACK_CELL = cwcomp.BLACK_CELL
)

func init() {
	LoadDictionary()
}
