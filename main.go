package main

import (
	"database/sql"
	"encoding/json"
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

// DBRow is the structure of the sensor table in the database.
// Each field is exported so that it can be json encoded.
// Record can be empty (null) so it uses the sql type NullString.
type DBRow struct {
	Id          int
	Date        string
	Beacon      string
	Name        string
	RecordType  int
	Record      sql.NullString
}

// Parse DBItem into string format for logging purpose
func (item DBRow) String() string {
	return fmt.Sprintf("\n\tID:\t%d\n\tDate:\t%s\n\tBeacon:\t%s\n\tName:\t%s\n\tRecord-Type:\t%d\n\tRecord:%s\n", item.Id, item.Date, item.Beacon, item.Name, item.RecordType, item.Record.String)
}

func main() {
	initDB()

	// TODO: might be better to implement a middleware to log and handle errors
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/ok", getHealth)
	http.HandleFunc("/data", getData)

	log.Println("Starting server...")
	log.Panic(http.ListenAndServe("localhost:8080", nil))
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
		log.Panic(err)
	}
	defer db.Close()

	// Create table
	query := `CREATE TABLE sensor
	(
		id INTEGER PRIMARY KEY,
		date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		beacon TEXT NOT NULL,
		name TEXT NOT NULL,
		record_type INTEGER NOT NULL,
		record BLOB
	)`
	if _, err = db.Exec(query); err != nil {
		log.Panicf("%q: %s\n", err, query)
	}
}

// Simple health check
func getHealth(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	log.Println("Health Check OK!")
	return
}

// Retrieve all data from the database
func getData(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(rw, "Only GET method is supported", http.StatusBadRequest)
		return
	}

	// TODO: request sensor values from the database
	db, err := sql.Open("sqlite3", "./sensor.db")
	if err != nil {
		http.Error(rw, "Connection to database failed", http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM sensor")
	if err != nil {
		// On query failure, report to the client but don't panic
		http.Error(rw, "Query failed", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	defer rows.Close()

	data := make([]DBRow, 0)
	for rows.Next() {
		var row DBRow
		if err := rows.Scan(&row.Id, &row.Date, &row.Beacon, &row.Name, &row.RecordType, &row.Record); err != nil {
			http.Error(rw, "Row scan failed", http.StatusInternalServerError)
			log.Print(err)
			return
		}
		data = append(data, row)
	}

	log.Println("Row successfully scanned, printing results...")
	fmt.Println(data)

	// Send content in json format
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}
