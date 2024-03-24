package rest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzleHandler(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "Good",
			url:  "/puzzles/asdf",
		},
		{
			name: "no ID",
			url:  "/puzzles/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader("")
			req, err := http.NewRequest("GET", tt.url, reader)
			assert.Nil(t, err)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PuzzleHandler)
			handler.ServeHTTP(rr, req)
		})
	}
}
