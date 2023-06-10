package svg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/philhanna/cwcomp"
)

// Cells generates the cells of the grid, including black cells and any
// letter values.
func (svg *SVG) Cells() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Cells -->\n")
	for r := 1; r <= svg.n; r++ {
		yBase := (r - 1) * BOXSIZE
		for c := 1; c <= svg.n; c++ {
			xBase := (c - 1) * BOXSIZE
			if svg.cells[r-1][c-1] == cwcomp.BLACK_CELL {
				sb.WriteString(fmt.Sprintf(
					"<rect x=%q y=%q width=%q height=%q fill=%q/>\n",
					strconv.Itoa(xBase),
					strconv.Itoa(yBase),
					strconv.Itoa(BOXSIZE),
					strconv.Itoa(BOXSIZE),
					"black",
				))
			} else {
				letter := svg.cells[r-1][c-1]
				if letter != ' ' {
					sb.WriteString(fmt.Sprintf(
						"<text x=%q y=%q font-size=%q font-family=%q>%c</text>\n",
						strconv.Itoa(xBase+LETTER_X_OFFSET),
						strconv.Itoa(yBase+LETTER_Y_OFFSET),
						LETTER_FONT_SIZE,
						LETTER_FONT_FAMILY,
						letter,
					))
				}
			}
		}
	}
	return sb.String()
}
