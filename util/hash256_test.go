package util

import (
	"fmt"
	"testing"
)

func TestHash256(t *testing.T) {
	tests := []struct {
		name     string
		password string
	}{
		{"waffle", "waffle"},
		{"123456", "123456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := Hash256(tt.password)
			fmt.Printf(`password=%s,hashed=x'%x'`+"\n", tt.password, have)
		})
	}
}
