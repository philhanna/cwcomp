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

	// Set the required CORS header(s)
	requester := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", requester)
	w.Header().Set("Access-Control-Allow-Headers", "Credentials")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// For a CORS preflight request, that's all we need
	if r.Method == "OPTIONS" {
		log.Println("Handled preflight OPTIONS request")
		return
	}

	// Get the username and password from the form data
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validate the user's credentials by looking up encrypted password
	// and comparing it to the one sent in the request
	// Get the row from the users table for this user name
	// If no user found, return 404 error
	// Now get the encrypted password stored in the database, hash the
	// one coming in with the request, and compare the two.
	userid, err := ValidateCredentials(w, username, password)
	if err != nil {
		log.Println(err)
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

// ValidateCredentials validates the login request according to the database
// and returns the userID (an integer) and any error
func ValidateCredentials(w http.ResponseWriter, username, password string) (int, error) {
	var (
		err            error
		hashedPassword = util.Hash256(password)
		userid         int
		dbPassword     []byte
	)

	// Get the userid and password from the database for this username
	con, _ := model.Connect()
	defer con.Close()

	rows, err := con.Query(`SELECT userid, password FROM users WHERE username=?`, username)
	if err != nil {
		errmsg := fmt.Sprintf("Could not read from users table")
		http.Error(w, errmsg, http.StatusInternalServerError)
		return 0, err
	}
	defer rows.Close()

	// Verify whether the username was found
	userFound := rows.Next()
	if !userFound {
		errmsg := fmt.Sprintf("Username %q not found in users table", username)
		http.Error(w, errmsg, http.StatusUnauthorized)
		err = fmt.Errorf(errmsg)
		return 0, err
	}

	// Verify whether the hashed passwords match
	rows.Scan(&userid, &dbPassword)
	passwordsMatch := bytes.Equal(hashedPassword, dbPassword)
	if !passwordsMatch {
		errmsg := "Passwords do not match"
		http.Error(w, errmsg, http.StatusUnauthorized)
		err = fmt.Errorf(errmsg)
		return 0, err
	}

	// Everything OK
	return userid, nil
}
