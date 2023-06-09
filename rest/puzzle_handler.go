package rest

import (
	"fmt"
	"net/http"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// PuzzleHandler serves REST requests for a single puzzle
func PuzzleHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/puzzles/"):]
	fmt.Fprintf(w, "This is the handler for the puzzle with id %v", id)
}
