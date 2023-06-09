package rest

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/philhanna/cwcomp/model"
	"github.com/philhanna/cwcomp/util"
)

// LoginHandler accepts a username and password from an HTTP POST payload
// and creates a session cookie that is returned to the requester.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering LoginHandler")

	// Get the username and password from the form data
	username := r.FormValue("username")
	password := r.FormValue("password")
	hashedPassword := util.Hash256(password)

	// Validate the user's credentials by looking up encrypted password
	// and comparing it to the one sent in the request
	con, _ := model.Connect()
	defer con.Close()

	// Get the row from the users table for this user name
	rows, err := con.Query(`SELECT userid, password FROM users WHERE username=?`, username)
	if err != nil {
		http.Error(w, "could not read from users table", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// If no user found, return 404 error
	userFound := rows.Next()
	if !userFound {
		errmsg := fmt.Sprintf("username %q not found in users table", username)
		http.Error(w, errmsg, http.StatusNotFound)
		return
	}

	// Now get the encrypted password stored in the database, hash the
	// one coming in with the request, and compare the two.
	var userid int
	var dbPassword []byte
	rows.Scan(&userid, &dbPassword)
	passwordsMatch := bytes.Equal(hashedPassword, dbPassword)
	if !passwordsMatch {
		errmsg := "Passwords do not match"
		http.Error(w, errmsg, http.StatusUnauthorized)
		return
	}

	// Create a new session and store it in the session map
	session := NewSession()
	session.USERID = userid
	session.USERNAME = username
	Sessions[session.ID] = session

	// Send the session cookie back to client
	cookie := session.NewSessionCookie()
	http.SetCookie(w, cookie)

	// Say it's OK
	w.WriteHeader(http.StatusOK)

	log.Println("Leaving LoginHandler")
}
