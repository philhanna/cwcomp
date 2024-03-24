package importer

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   error
		wantTitle string
		wantSize  int
	}{
		{
			name: "Good",
			input: func() string {
				filename := filepath.Join("..", "testdata", "goodfile.txt")
				body, err := os.ReadFile(filename)
				assert.Nil(t, err)
				return string(body)
			}(),
			wantErr:   nil,
			wantTitle: `Politics: Who, what, where and why`,
			wantSize:  15,
		},
		{
			name:    "INIT",
			input:   "",
			wantErr: errNoLines,
		},
		{
			name: "Looking for title",
			input: `<ACROSS PUZZLE>
			<AUTHOR> 
				Created by Avalonian`,
			wantErr: errNoTitle,
		},
		{
			name:    "No title",
			input:   `<ACROSS PUZZLE>`,
			wantErr: errNoTitle,
		},
		{
			name: "No size",
			input: `<ACROSS PUZZLE>
			<TITLE> 
				Politics: Who, what, where and why
			<AUTHOR> 
				Created by Avalonian
			<COPYRIGHT> 
				Literate Software Systems
			<GRID>
				FATE.AWASH.AWOL`,
			wantErr: errNoSize,
		},
		{
			name: "No author",
			input: `<ACROSS PUZZLE>
<TITLE>
     Politics: Who, what, where and why
			`,
			wantErr: errNoAuthor,
		},
		{
			name: "Reading title",
			input: `<ACROSS PUZZLE>
<TITLE>
			`,
			wantErr: errNoAuthor,
		},
		{
			name: "No copyright",
			input: `<ACROSS PUZZLE>
			<TITLE> 
				Politics: Who, what, where and why
			<AUTHOR> 
				Created by Avalonian
			<SIZE>`,
			wantErr: errNoCopyright,
		},
		{
			name: "No grid",
			input: `<ACROSS PUZZLE>
			<TITLE> 
				Politics: Who, what, where and why
			<AUTHOR> 
				Created by Avalonian
			<COPYRIGHT> 
				Literate Software Systems
			<SIZE> 
				15x15
			<ACROSS>    `,
			wantErr: errNoGrid,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			puzzle, err := Parse(reader)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.wantTitle, puzzle.Title)
				assert.Equal(t, tt.wantSize, puzzle.Size)
			}
		})
	}

}
