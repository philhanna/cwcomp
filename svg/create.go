package svg

import (
	"fmt"
	"strconv"
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

const (
	XMLNS       = "http://www.w3.org/2000/svg"
	XMLNS_XLINK = "http://www.w3.org/1999/xlink"
	BLACK_CELL  = '\x00'
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

// NewSVGFromPuzzle will create a new SVG object from a grid, delegating
// that to NewSVG after creating the simple cell matrix it needs.
func NewSVGFromPuzzle(puzzle *model.Puzzle) *SVG {
	cells := model.PuzzleToSimpleMatrix(puzzle)
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

// Cells generates the cells of the grid, including black cells and any
// letter values.
func (svg *SVG) Cells() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Cells -->\n")
	for r := 1; r <= svg.n; r++ {
		yBase := (r - 1) * BOXSIZE
		for c := 1; c <= svg.n; c++ {
			xBase := (c - 1) * BOXSIZE
			if svg.cells[r-1][c-1] == BLACK_CELL {
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

// WordNumbers generates the word numbers of the grid.
func (svg *SVG) WordNumbers() string {
	sb := strings.Builder{}
	sb.WriteString("\n<!-- Word numbers -->\n")
	for _, nc := range model.GetNumberedCells(svg.cells) {
		seq := nc.Seq
		xbase := (nc.Col-1)*BOXSIZE + NUMBER_X_OFFSET
		ybase := (nc.Row-1)*BOXSIZE + NUMBER_Y_OFFSET
		fmtstr := `<text x="%d" y="%d" font-size="%s">%d</text>`
		line := fmt.Sprintf(fmtstr, xbase, ybase, NUMBER_FONT_SIZE, seq)
		sb.WriteString(line + "\n")
	}
	return sb.String()
}

// EndRoot creates the closing </svg> element.
func (svg *SVG) EndRoot() string {
	return "\n</svg>\n"
}
