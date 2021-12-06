// Package for handling incoming requests to the server and checking the beacons
// at a fixed interval
package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/thinkty/warden/internal/database"
)

const addr string = "localhost:8080"
const static_path string = "./web"

// Initialize the router by specifying the handlers to each path and start the
// actual server on the specified address
func InitAndServe() {
	http.Handle("/", http.FileServer(http.Dir(static_path)))
	http.HandleFunc("/ok", getHealth)
	http.HandleFunc("/data", getData)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe(addr, nil))
}

// Simple health check
func getHealth(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Health Check OK!")
	log.Println("Health Check OK!")
	return
}

// Fetch all data from the database
func getData(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(rw, "Only GET method is supported", http.StatusBadRequest)
		return
	}

	data, err, err_msg := database.GetData()
	if err != nil {
		http.Error(rw, err_msg, http.StatusInternalServerError)
		log.Println(err_msg)
		log.Println(err)
	}

	// Send content in json format
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
