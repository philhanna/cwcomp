package svg

import (
	"fmt"
	"strconv"
	"strings"
)

// Root creates the root <svg> element.
func (svg *SVG) Root() string {

	sb := strings.Builder{}

	// Write the first part of the tag (before we know whether we need
	// scale)
	sb.WriteString(fmt.Sprintf("<svg xmlns=%q xmlns:xlink=%q id=%q",
		XMLNS,
		XMLNS_XLINK,
		fmt.Sprintf("svg%dx%d", svg.n, svg.n),
	))

	// Add the width and height
	sizeAttr := strconv.Itoa(svg.nPixels)
	sb.WriteString(fmt.Sprintf(" width=%q", sizeAttr))
	sb.WriteString(fmt.Sprintf(" height=%q", sizeAttr))

	// Add the viewport
	vattr := fmt.Sprintf("%d %d %d %d", 0, 0, svg.nPixels, svg.nPixels)
	sb.WriteString(fmt.Sprintf(" viewport=%q", vattr))

	// Done
	sb.WriteString(">\n")
	return sb.String()
}
