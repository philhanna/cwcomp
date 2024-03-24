package importer

import (
	"testing"

	al "github.com/philhanna/cwcomp/acrosslite"
	"github.com/stretchr/testify/assert"
)

func TestHandleInit(t *testing.T) {
	tests := []struct {
		name      string
		pal       *al.AcrossLite
		line      string
		wantState ParsingState
		wantErr   bool
	}{
		{
			name:      "Good",
			pal:       nil,
			line:      "<ACROSS PUZZLE>",
			wantState: LOOKING_FOR_TITLE,
			wantErr:   false,
		},
		{
			name:      "Bogus",
			pal:       nil,
			line:      "<BOGUS>",
			wantState: UNKNOWN,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			haveState, err := HandleInit(tt.pal, tt.line)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, tt.wantState, haveState)
		})
	}
}
