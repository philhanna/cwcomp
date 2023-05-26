package svg

import (
	"strings"

	"github.com/philhanna/cwcomp/model"
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

const BOXSIZE = 32
const XMLNS = "http://www.w3.org/2000/svg"
const XMLNS_XLINK = "http://www.w3.org/1999/xlink"

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

// NewSVGFromGrid will create a new SVG object from a grid, delegating
// that to NewSVG after creating the simple cell matrix it needs.
func NewSVGFromGrid(grid *model.Grid) *SVG {
	cells := model.GridToCells(grid)
	return NewSVG(cells)
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

// Root creates the root <svg> element.
func (svg *SVG) Root() string {
	return "" // TODO implement me
}

// BoundingRectangle creates an element that draws a box around the
// whole grid.
func (svg *SVG) BoundingRectangle() string {
	return "" // TODO implement me
}

// VerticalLines generates the vertical lines of the grid.
func (svg *SVG) VerticalLines() string {
	return "" // TODO implement me
}

// HorizontalLines generates the horizontal lines of the grid
func (svg *SVG) HorizontalLines() string {
	return "" // TODO implement me
}

// Cells generates the cells of the grid, including black cells and any
// letter values.
func (svg *SVG) Cells() string {
	return "" // TODO implement me
}

// WordNumbers generates the word numbers of the grid.
func (svg *SVG) WordNumbers() string {
	return "" // TODO implement me
}

// EndRoot creates the closing </svg> element.
func (svg *SVG) EndRoot() string {
	return "</svg>\n"
}