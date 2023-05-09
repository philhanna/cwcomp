package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberedCell_String(t *testing.T) {
	tests := []struct {
		name string
		lc   LetterCell
		seq  int
		aLen int
		dLen int
		want string
	}{
		{
			"simple",
			LetterCell{
				point:    Point{1, 2},
				ncAcross: nil,
				ncDown:   nil,
				letter:   "O",
			},
			1,
			3,
			4,
			`LetterCell:{point:{1,2},ncAcross:<nil>,ncDown:<nil>,letter:"O"},seq:1,aLen:3,dLen:4`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NumberedCell{tt.lc, tt.seq, tt.aLen, tt.dLen}
			assert.Equal(t, tt.want, nc.String())
		})
	}
}
