package svg

import (
	"fmt"
	"strconv"
	"strings"
)

// HorizontalLines generates the horizontal lines of the grid
func (svg *SVG) HorizontalLines() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Horizontal lines -->\n")
	for x := 0; x < svg.n; x++ {
		sb.WriteString(fmt.Sprintf(
			"<line x1=%q x2=%q y1=%q y2=%q stroke=%q stroke-width=%q/>\n",
			"0",                       // x1
			strconv.Itoa(svg.nPixels), // x2
			strconv.Itoa(x*BOXSIZE),   // y1
			strconv.Itoa(x*BOXSIZE),   // y2
			"black",                   // stroke
			"0.5",                     // stroke-width
		))
	}
	return sb.String()
}
