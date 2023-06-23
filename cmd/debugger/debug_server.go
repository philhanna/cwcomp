package main

import (
	"fmt"
	"github.com/philhanna/cwcomp/rest"
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
	log.Printf("Incoming method:  %q\n", r.Method)
	log.Printf("Incoming headers: %v", dumpHeaders(r.Header))

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	if r.Method == "POST" {
		username, password, err := rest.UnmarshalCredentials(r, w)
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
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.Header().Set("Access-Control-Allow-Headers", "credentials, content-type, access-control-allow-origin")
	w.WriteHeader(http.StatusOK)

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
	sb.WriteString("\n")
	for name, values := range headers {
		parts := []string{}
		for _, value := range values {
			part := value
			parts = append(parts, part)
		}
		value := strings.Join(parts, ", ")
		sb.WriteString(fmt.Sprintf("    %q: %q\n", name, value))
	}
	return sb.String() + "\n\n"
}
