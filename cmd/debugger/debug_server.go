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
	log.Printf("Incoming method:  %q\n", r.Method)
	log.Printf("Incoming headers: %v", dumpHeaders(r.Header))

	// Call the wrapped handler
	h.handler.ServeHTTP(w, r)

	// Log outgoing headers
	log.Printf("Outgoing headers: %v", dumpHeaders(w.Header()))
}

func main() {
	// Create a new handler that logs headers and wraps the existing handler
	handler := loggingHandler{http.DefaultServeMux}

	// Start the server
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