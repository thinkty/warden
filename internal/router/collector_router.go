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
func InitAndServe() {
	http.Handle("/", http.FileServer(http.Dir(staticPath)))
	http.HandleFunc("/ok", getCollectorRouterHealth)

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
