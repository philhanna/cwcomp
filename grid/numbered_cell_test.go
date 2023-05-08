package grid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberedCell_String(t *testing.T) {
	tests := []struct {
		name   string
		seq    int
		aLen   int
		dLen   int
		letter string
		want   string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nc := &NumberedCell{
				seq:    tt.seq,
				aLen:   tt.aLen,
				dLen:   tt.dLen,
				letter: tt.letter,
			}
			assert.Equal(t, tt.want, nc.String())
		})
	}
}
