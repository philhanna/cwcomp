package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestDumpHeaders(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
	}{
		{
			name: "good",
			headers: func() http.Header {
				rr := httptest.NewRecorder()
				rr.Header().Add("Tom", "Older")
				return rr.Header()
			}(),
			want: "\n    \"Tom\": \"Older\"\n\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.want
			have := DumpHeaders(tt.headers)
			assert.Equal(t, want, have)
		})
	}
}

func TestUnmarshalCredentials(t *testing.T) {
	tests := []struct {
		name         string
		w            http.ResponseWriter
		r            *http.Request
		wantUsername string
		wantPassword string
	}{
		{
			name: "good",
			w:    nil,
			r: func() *http.Request {
				r, err := http.NewRequest(
					"POST",
					"https://www.example.com",
					bytes.NewReader(
						[]byte(``),
					),
				)
				assert.Nil(t, err)
				r.Form = map[string][]string{
					"username": {"seville"},
					"password": {"dardego"},
				}
				return r
			}(),
			wantUsername: "seville",
			wantPassword: "dardego",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			username, password := UnmarshalCredentials(tt.w, tt.r)
			assert.Equal(t, tt.wantUsername, username)
			assert.Equal(t, tt.wantPassword, password)
		})
	}
}
