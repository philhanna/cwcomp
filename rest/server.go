package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/philhanna/cwcomp"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------
// HandleRequests registers all the routers and starts the server.
func HandleRequests() {

	// Get host and port from configuration
	config := cwcomp.Configuration
	host := config.SERVER.HOST
	port := config.SERVER.PORT
	hostAndPort := fmt.Sprintf("%s:%d", host, port)

	// Define the handler functions
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/puzzles", PuzzlesHandler)
	http.HandleFunc("/puzzles/", PuzzleHandler)

	// Start the server
	log.Printf("Starting server on %v\n", hostAndPort)
	log.Fatal(http.ListenAndServe(hostAndPort, nil))
}
