package rest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Session struct {
	ID       string
	EXPIRES  time.Time
	USERID   int
	USERNAME string
}

// ---------------------------------------------------------------------
// Variables and constants
// ---------------------------------------------------------------------

const (
	SESSION_COOKIE = "session_id"
)

// Sessions is a map of session IDs to session pointers
var Sessions = make(map[string]*Session)

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewSession creates a new session and returns a pointer to it.
// The session expires in 30 minutes.
func NewSession() *Session {
	ps := new(Session)
	ps.ID = uuid.NewString()
	ps.EXPIRES = time.Now().Add(30 * time.Minute)
	return ps
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Renew extends the session expiration time to another 30 minutes.
func (ps *Session) Renew() {
	ps.EXPIRES = time.Now().Add(30 * time.Minute)
}

// NewSessionCookie creates a session cookie and returns a pointer to it.
func (ps *Session) NewSessionCookie() *http.Cookie {
	cookie := http.Cookie{
		Name:     "session",
		Value:    ps.ID,
		Path:     "/",                     // Set the cookie path to the root
		HttpOnly: true,                    // Ensure the cookie is only accessible via HTTP(S)
		Secure:   true,                    // Send the cookie only over HTTPS
		SameSite: http.SameSiteStrictMode, // Enforce strict same-site policy
		Expires:  ps.EXPIRES,              // Set an expiration time for the cookie
	}
	return &cookie
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// GetSession gets the session ID from the request cookie and looks up
// the correct session in the session map. If there is no cookie, or if
// the session does not exist or has expired, a 401 HTTP error is
// written to the response
func GetSession(w http.ResponseWriter, r *http.Request) (*Session, error) {

	// Check for the existence of a session cookie, which will contain
	// the session ID and an expiration time. If no session ID is found,
	// or if the session is expired, go to the login screen.
	cookie, err := r.Cookie("session")
	if err != nil {
		err := fmt.Errorf("No session cookie found: %v", err)
		log.Println(err)
		return nil, err
	}

	session_id := cookie.Value

	// Get the session from the map
	session, ok := Sessions[session_id]
	if !ok {
		err := fmt.Errorf("Session id %q not found in session map", session_id)
		log.Println(err)
		return nil, err
	}

	// Check for expired session
	if session.EXPIRES.Before(time.Now()) {
		err := fmt.Errorf("Session id %q has expired\n", session_id)
		log.Println(err)
		return nil, err
	}

	return session, nil
}
