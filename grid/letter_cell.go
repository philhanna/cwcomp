package grid

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Letter cell is an ordinary point in the grid. It contains:
//   - A pointer to the numbered cell for the across word (if any)
//   - A pointer to the numbered cell for the down word (if any)
//   - The character in the cell
type LetterCell struct {
	ncAcross *Point // Pointer to the numbered cell in the across direction
	ncDown   *Point // Pointer to the numbered cell in the down direction
	letter   string // Character in the cell
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// LetterCell returns a string representation of this letter cell.
func (lc *LetterCell) String() string {
	parts := []string{
		fmt.Sprintf(`ncAcross:%p`, lc.ncAcross),
		fmt.Sprintf(`ncDown:%p`, lc.ncDown),
		fmt.Sprintf(`letter:%q`, lc.letter),
	}
	s := strings.Join(parts, ", ")
	return s
}
