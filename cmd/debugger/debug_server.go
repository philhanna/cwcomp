package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type LoggingHandler struct {
	handler http.Handler
}

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	PORT = 5053
)

// ---------------------------------------------------------------------
// Mainline
// ---------------------------------------------------------------------

func main() {
	// Create a new handler that logs headers and wraps the existing handler
	handler := LoggingHandler{http.DefaultServeMux}

	// Start the server
	hostport := fmt.Sprintf("localhost:%d", PORT)
	log.Printf("Starting debug server on %s", hostport)
	log.Fatal(http.ListenAndServe(hostport, handler))
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// DumpHeaders creates a string representation of the values in an
// http.Header object.
func DumpHeaders(headers http.Header) string {
	sb := strings.Builder{}
	for name, values := range headers {
		sb.WriteString("\n")
		parts := []string{}
		for _, value := range values {
			part := value
			parts = append(parts, part)
		}
		value := strings.Join(parts, ", ")
		sb.WriteString(fmt.Sprintf("    %q: %q", name, value))
	}
	return sb.String() + "\n\n"
}

// init sets the logging flags at startup
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// UnmarshalCredentials extracts the username and password from a
// request.
func UnmarshalCredentials(w http.ResponseWriter, r *http.Request) (string, string) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	return username, password
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// ServeHTTP handles a request
func (h LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log incoming headers
	log.Printf("Incoming URL: %q\n", r.URL.String())
	log.Printf("Incoming method: %q\n", r.Method)
	log.Printf("Incoming headers: %v", DumpHeaders(r.Header))

	// Add required CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {

	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "credentials, content-type, access-control-allow-origin")
		w.WriteHeader(http.StatusOK)

	case "POST":
		username, password := UnmarshalCredentials(w, r)
		if username != "saspeh" {
			errmsg := "user name not found in users table"
			http.Error(w, errmsg, http.StatusNotFound)
			log.Println(errmsg)
			return
		} else if password != "waffle" {
			errmsg := "passwords do not match"
			http.Error(w, errmsg, http.StatusUnauthorized)
			log.Println(errmsg)
			return
		}
		// TODO handle login for real
		w.WriteHeader(http.StatusOK)
	}

	// Log outgoing headers
	log.Printf("Outgoing headers: %v", DumpHeaders(w.Header()))
}
