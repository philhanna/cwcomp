package model

import "testing"

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
