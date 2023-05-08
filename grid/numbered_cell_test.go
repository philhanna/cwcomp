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
	}{
		{"simple", 1, 3, 4, "O", `seq:1,aLen:3,dLen:4,letter:"O"`},
	}
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
