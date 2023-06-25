package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type loggingHandler struct {
	handler http.Handler
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Log incoming headers
	log.Printf("Incoming URL: %q\n", r.URL.String())
	log.Printf("Incoming method: %q\n", r.Method)
	log.Printf("Incoming headers: %v", dumpHeaders(r.Header))

	// Add required CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {

	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "credentials, content-type, access-control-allow-origin")
		w.WriteHeader(http.StatusOK)
	
	case "POST":
		username, password, err := UnmarshalCredentials(w, r)
		if err != nil {
			log.Fatal(err)
		}
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
	log.Printf("Outgoing headers: %v", dumpHeaders(w.Header()))
}

func main() {
	// Create a new handler that logs headers and wraps the existing handler
	handler := loggingHandler{http.DefaultServeMux}

	// Start the server
	log.Printf("Starting debug server on localhost:5053")
	log.Fatal(http.ListenAndServe(":5053", handler))
}

func dumpHeaders(headers http.Header) string {
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

func UnmarshalCredentials(w http.ResponseWriter, r *http.Request) (string, string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", "", err
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	return username, password, nil
}