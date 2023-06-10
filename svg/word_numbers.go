package svg

import (
	"fmt"
	"strings"

	"github.com/philhanna/cwcomp/model"
)

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
