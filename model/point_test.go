package model

import "testing"
import "github.com/stretchr/testify/assert"

func TestPoint_Compare(t *testing.T) {
	tests := []struct {
		name string
		this Point
		that Point
		want int
	}{
		{"same point", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			that := tt.that
			assert.Equal(t, 0, this.Compare(that))
		})
	}
}
