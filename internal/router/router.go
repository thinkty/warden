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
const staticPath string = "./web"

// Initialize the router by specifying the handlers to each path and start the
// actual server on the specified address
func InitAndServe() {
	http.Handle("/", http.FileServer(http.Dir(staticPath)))
	http.HandleFunc("/ok", getHealth)
	http.HandleFunc("/data", getData)
	http.HandleFunc("/test", putData) // TODO: Temporary

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

	data, err, errMsg := database.ReadRecords()
	if err != nil {
		http.Error(rw, errMsg, http.StatusInternalServerError)
		log.Println(errMsg)
		log.Println(err)
	}

	// Send content in json format
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

// TODO: Temporary
func putData(rw http.ResponseWriter, r *http.Request) {
	err, errMsg := database.CreateRecord("testbeacon", "testname", 1, "asdf")

	if err != nil {
		http.Error(rw, errMsg, http.StatusInternalServerError)
		log.Println(errMsg)
		log.Println(err)
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Success!")
	log.Println("Success!")
	return
}
