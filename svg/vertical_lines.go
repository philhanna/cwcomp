package svg

import (
	"fmt"
	"strconv"
	"strings"
)

// VerticalLines generates the vertical lines of the grid.
func (svg *SVG) VerticalLines() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Vertical lines -->\n")
	for x := 0; x < svg.n; x++ {
		sb.WriteString(fmt.Sprintf(
			"<line x1=%q x2=%q y1=%q y2=%q stroke=%q stroke-width=%q/>\n",
			strconv.Itoa(x*BOXSIZE),   // x1
			strconv.Itoa(x*BOXSIZE),   // x2
			"0",                       // y1
			strconv.Itoa(svg.nPixels), // y2
			"black",                   // stroke
			"0.5",                     // stroke-width
		))
	}
	return sb.String()
}
