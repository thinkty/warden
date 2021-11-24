package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// BTRequest is the data structure used for sending queries from the server to
// the bluetooth device.
// Target is the name of the bluetooth device.
// Command specifies which feature of the bluetooth device it should execute.
// Value is an optional parameter for specific command operations (ex: setting
// the status of the LED)
type BTRequest struct {
	target  string
	command byte
	value   byte
}

// BTResponse is the data structure used for receiving responses from the
// bluetooth device to the server
// Target specifies which bluetooth device it is coming from.
// Key is the name of the feature the bluetooth device is reporting (ex:
// TEMPERATURE, LED)
// Value is the actual value corresponding to the key (ex: actual temperature
// values, LED status on/off)
type BTResponse struct {
	target string
	key    string
	value  string
}

func main() {
	initDB()

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/data", getData)

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// Create the database file if it does not exist
func initDB() {
	log.Println("Initializing DB file...")

	// Return if the file exists
	if _, err := os.Stat("./sensor.db"); !errors.Is(err, os.ErrNotExist) {
		log.Println("DB file exists, skipping creation...")
		return
	}
	log.Println("DB file does not exist, proceeding initialization...")

	db, err := sql.Open("sqlite3", "./sensor.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `create table sensor (id integer not null primary key, name text);`
	if _, err = db.Exec(query); err != nil {
		log.Fatalf("%q: %s\n", err, query)
	}
}

// Retrieve all data from the database
func getData(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(rw, "Only GET method is supported", http.StatusBadRequest)
		return
	}

	// TODO: request sensor values from the database

	// rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, "OK!")
}
