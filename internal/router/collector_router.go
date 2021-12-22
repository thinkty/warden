// Package for handling beacon fetch and user actions
package router

import (
	"fmt"
	"log"
	"net/http"
)

const collectorAddr string = "localhost:8082"

// Initialize the router by specifying the handlers to each path and start the
// actual server on the specified address
func InitAndServeCollector() {
	http.HandleFunc("/ok", getCollectorRouterHealth)
	http.HandleFunc("/action", handleUserAction)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(collectorAddr, nil))
}

// Simple health check
func getCollectorRouterHealth(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Health Check OK!")
	log.Println("Health Check OK!")
	return
}

// Check the user action and handle it
func handleUserAction(rw http.ResponseWriter, r *http.Request) {
	// TODO: the data structure used for passing user action is not clear yet

	// temporary code
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "OK!")
	log.Println("Handle user action")
}
