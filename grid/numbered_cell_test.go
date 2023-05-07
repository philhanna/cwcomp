package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberedCell_String(t *testing.T) {
	type fields struct {
		LetterCell LetterCell
		seq        int
		aLen       int
		dLen       int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"simple",
			fields{
				LetterCell{
					ncAcross: nil,
					ncDown:   nil,
					letter:   "A",
				},
				1, 3, 4,
			},
			"letterCell:{<nil> <nil> A}, seq:1, aLen:3, dLen:4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NumberedCell{
				LetterCell: tt.fields.LetterCell,
				seq:        tt.fields.seq,
				aLen:       tt.fields.aLen,
				dLen:       tt.fields.dLen,
			}
			assert.Equal(t, tt.want, nc.String())
		})
	}
}
