// Package for handling incoming requests to the interfacer
package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/thinkty/warden/internal/database"
)

const interfacerAddr string = "localhost:8081"
const staticPath string = "./web/dist"

// Initialize the router by specifying the handlers to each path and start the
// actual server on the specified address
func InitAndServeInterfacer() {
	http.Handle("/", http.FileServer(http.Dir(staticPath)))
	http.HandleFunc("/ok", getInterfacerRouterHealth)
	http.HandleFunc("/data", getData)
	http.HandleFunc("/data-test", getSampleData) // TODO: Temporary
	http.HandleFunc("/test", putData)            // TODO: Temporary

	log.Printf("Starting server at %s ...", interfacerAddr)
	log.Panic(http.ListenAndServe(interfacerAddr, nil))
}

// Simple health check
func getInterfacerRouterHealth(rw http.ResponseWriter, r *http.Request) {
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

	data, err, errMsg := database.ReadRecords(false)
	if err != nil {
		http.Error(rw, errMsg, http.StatusInternalServerError)
		log.Println(errMsg)
		log.Println(err)
	}

	// Send content in json format
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

// Fetch all data from the sample database
func getSampleData(rw http.ResponseWriter, r *http.Request) {
	data, err, errMsg := database.ReadRecords(true)
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
	err, errMsg := database.CreateRecord("testbeacon", "testname", "On", "Off:0,On:1")

	if err != nil {
		http.Error(rw, errMsg, http.StatusInternalServerError)
		log.Println(errMsg)
		log.Println(err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, "Success!")
	log.Println("Success!")
	return
}
