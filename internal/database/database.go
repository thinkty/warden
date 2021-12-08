// Package for handling DB operations
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// DBRow is the structure of the sensor table in the database. Each field is
// exported so that it can be json encoded. Record can be empty (null) so it
// uses the sql type NullString.
type DBRow struct {
	Id         int
	Date       string
	Beacon     string
	Name       string
	RecordType int
	Record     sql.NullString
}

// Parse DBItem into string format for logging purpose
func (item DBRow) String() string {
	return fmt.Sprintf("\n\tID:\t%d\n\tDate:\t%s\n\tBeacon:\t%s\n\tName:\t%s\n\tRecord-Type:\t%d\n\tRecord:%s\n", item.Id, item.Date, item.Beacon, item.Name, item.RecordType, item.Record.String)
}

// Initialize the database by creating the sql file if it does not already exist
func Init() {
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

// Create a new record in the database with the given arguments. Should probably
// validate the arguments.
func CreateRecord(beacon string, name string, recordType int, record string) (err error, errMsg string) {

	// Connect to database
	db, err := sql.Open("sqlite3", "./sensor.db")
	if err != nil {
		return err, "Connection to database failed"
	}
	defer db.Close()

	// Prepare transaction and prepare query statement
	tx, err := db.Begin()
	if err != nil {
		return err, "Failed to start transaction"
	}

	stmt, err := tx.Prepare("INSERT INTO sensor(beacon, name, record_type, record) values(?, ?, ?, ?)")
	if err != nil {
		return err, "Failed to prepare SQL query statement"
	}
	defer stmt.Close()

	_, err = stmt.Exec(beacon, name, recordType, record)
	if err != nil {
		return err, "Failed to execute SQL query statement"
	}
	tx.Commit()
	return nil, ""
}

// Retrieve all data from the database. If error has occurred, return the error
// along with the descriptive error message
func ReadRecords() ([]DBRow, error, string) {

	// Connect to database
	db, err := sql.Open("sqlite3", "./sensor.db")
	if err != nil {
		return nil, err, "Connection to database failed"
	}
	defer db.Close()

	// Request sensor values from the database
	rows, err := db.Query("SELECT * FROM sensor")
	if err != nil {
		return nil, err, "Query failed"
	}
	defer rows.Close()

	data := make([]DBRow, 0)
	for rows.Next() {
		var row DBRow
		if err := rows.Scan(&row.Id, &row.Date, &row.Beacon, &row.Name, &row.RecordType, &row.Record); err != nil {
			return nil, err, "Row scan failed"
		}
		data = append(data, row)
	}

	log.Print("Row successfully scanned...")
	if len(data) != 0 {
		log.Print("Printing scan results...")
		fmt.Println(data)
	}

	return data, nil, ""
}
