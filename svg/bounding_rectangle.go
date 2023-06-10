package svg

import (
	"fmt"
	"strconv"
	"strings"
)

// BoundingRectangle creates an element that draws a box around the
// whole grid.
func (svg *SVG) BoundingRectangle() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Bounding rectangle -->\n")
	sb.WriteString(fmt.Sprintf(
		"<rect width=%q height=%q fill=%q stroke=%q stroke-width=%q/>\n",
		strconv.Itoa(svg.nPixels), // width
		strconv.Itoa(svg.nPixels), // height
		"white",                   // fill
		"black",                   // stroke
		"2",                       // stroke-width
	))
	return sb.String()
}
