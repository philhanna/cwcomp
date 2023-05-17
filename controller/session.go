package controller

import "github.com/philhanna/cwcomp/model"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Session contains HTTP session objects for the current user
type Session struct {
	userid int
	grid   model.Grid
}
