package svg

import (
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type SVG struct {
	n       int      // Number of either rows or columns in the grid (assumed to be the same)
	nPixels int      // Width or height of grid in pixels
	cells   [][]byte // Simple matrix representation of the grid
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

const (
	XMLNS       = "http://www.w3.org/2000/svg"
	XMLNS_XLINK = "http://www.w3.org/1999/xlink"
	BOXSIZE     = 32

	LETTER_X_OFFSET    = 8
	LETTER_Y_OFFSET    = 28
	LETTER_FONT_FAMILY = "monospace"
	LETTER_FONT_SIZE   = "18pt"

	NUMBER_X_OFFSET  = 2
	NUMBER_Y_OFFSET  = 10
	NUMBER_FONT_SIZE = "8pt"
)

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewSVG will create a new SVG object from an abstract matrix of bytes.
// This for convenience of unit testing, since it does not need any
// reference to the model.
func NewSVG(cells [][]byte) *SVG {
	svg := new(SVG)
	svg.n = len(cells)
	svg.nPixels = svg.n * BOXSIZE
	svg.cells = cells
	return svg
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GenerateSVG creates a SVG image of the grid.
func (svg *SVG) GenerateSVG() string {
	sb := strings.Builder{}
	sb.WriteString(svg.Root())
	sb.WriteString(svg.BoundingRectangle())
	sb.WriteString(svg.VerticalLines())
	sb.WriteString(svg.HorizontalLines())
	sb.WriteString(svg.Cells())
	sb.WriteString(svg.WordNumbers())
	sb.WriteString(svg.EndRoot())
	return sb.String()
}

// EndRoot creates the closing </svg> element.
func (svg *SVG) EndRoot() string {
	return "\n</svg>\n"
}
