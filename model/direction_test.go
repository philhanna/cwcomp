package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirection_Other(t *testing.T) {
	tests := []struct {
		name string
		dir  Direction
		want Direction
	}{
		{"Across", ACROSS, DOWN},
		{"Down", DOWN, ACROSS},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dir.Other(); got != tt.want {
				t.Errorf("Direction.Other() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirectionFromString(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		want   Direction
		wantOK bool
	}{
		{"empty string", "", *new(Direction), false},
		{"bogus string", "bogus", *new(Direction), false},
		{"exact match for ACROSS", "A", ACROSS, true},
		{"exact match for DOWN", "D", DOWN, true},
		{"full word across", "across", ACROSS, true},
		{"full word down", "down", DOWN, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantOK {
				have := DirectionFromString(tt.s)
				assert.Equal(t, have, tt.want)
			} else {
				assert.Panics(t, func() {
					DirectionFromString(tt.s)
				})
			}
		})
	}
}
