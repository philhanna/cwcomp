package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			`Username "bogus" not found in users table`,
		},
		{
			"invalid password", "saspeh", "wffl",
			http.StatusUnauthorized,
			`Passwords do not match`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postdata := fmt.Sprintf("{%q:%q,%q:%q}",
				"username", tt.username,
				"password", tt.password)
			reader := strings.NewReader(postdata)
			req, err := http.NewRequest("POST", "/login", reader)
			assert.Nil(t, err)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LoginHandler)
			handler.ServeHTTP(rr, req)

			haveStatus := rr.Code
			wantStatus := tt.wantStatus
			assert.Equal(t, wantStatus, haveStatus)

			body, err := ioutil.ReadAll(rr.Body)
			assert.Nil(t, err)

			wantBody := tt.wantBody
			haveBody := strings.Trim(string(body), "\n")
			assert.Equal(t, wantBody, haveBody)
		})
	}
}
