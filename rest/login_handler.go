package rest

import (
	"bytes"
	"encoding/json"
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

	// Get the username and password from the JSON form data
	username, password, err := UnmarshalCredentials(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Got username=%q\n", username)
	log.Printf("Got password=%q\n", password)

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
	log.Printf("Got userid=%d\n", userid)


	// Create a new session and store it in the session map
	session := NewSession()
	session.USERID = userid
	session.USERNAME = username
	Sessions[session.ID] = session

	w.Header().Set("Content-Type", "application/json")

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
		errmsg := fmt.Sprintf("could not read from users table")
		http.Error(w, errmsg, http.StatusInternalServerError)
		return 0, err
	}
	defer rows.Close()

	// Verify whether the username was found
	userFound := rows.Next()
	if !userFound {
		errmsg := fmt.Sprintf("username %q not found in users table", username)
		http.Error(w, errmsg, http.StatusNotFound)
		err = fmt.Errorf(errmsg)
		return 0, err
	}

	// Verify whether the hashed passwords match
	rows.Scan(&userid, &dbPassword)
	passwordsMatch := bytes.Equal(hashedPassword, dbPassword)
	if !passwordsMatch {
		errmsg := "passwords do not match"
		http.Error(w, errmsg, http.StatusUnauthorized)
		err = fmt.Errorf(errmsg)
		return 0, err
	}

	// Everything OK
	return userid, nil
}

// UnmarshalCredentials gets the user name and password from the request,
// or returns an error if they could not be found
func UnmarshalCredentials(r *http.Request, w http.ResponseWriter) (string, string, error) {

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		return "", "", err
	}
	username := cred.Username
	password := cred.Password
	return username, password, nil
}
