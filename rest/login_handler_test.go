package rest

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionsHandler(t *testing.T) {
	cred := url.Values{}
	cred.Set("username", "saspeh")
	cred.Set("password", "waffle")
	postdata := cred.Encode()

	reader := strings.NewReader(postdata)
	req, err := http.NewRequest("OPTIONS", "/login", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)
}

func TestLoginHandler(t *testing.T) {
	type Test struct {
		name       string // Test name
		username   string
		password   string
		wantStatus int
		wantBody   string
	}
	tests := []Test{
		{
			"valid user", "saspeh", "waffle",
			http.StatusOK,
			``,
		},
		{
			"invalid user", "bogus", "waffle",
			http.StatusUnauthorized,
			`username "bogus" not found in users table`,
		},
		{
			"invalid password", "saspeh", "wffl",
			http.StatusUnauthorized,
			`passwords do not match`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cred := url.Values{}
			cred.Set("username", tt.username)
			cred.Set("password", tt.password)
			postdata := cred.Encode()

			reader := strings.NewReader(postdata)
			req, err := http.NewRequest("POST", "/login", reader)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LoginHandler)
			handler.ServeHTTP(rr, req)

			haveStatus := rr.Code
			wantStatus := tt.wantStatus
			assert.Equal(t, wantStatus, haveStatus)

			body, err := io.ReadAll(rr.Body)
			assert.Nil(t, err)

			wantBody := tt.wantBody
			haveBody := strings.Trim(string(body), "\n")
			assert.Equal(t, wantBody, haveBody)
		})
	}
}
