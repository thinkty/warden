// Package for keeping the main package slim, initiating the router, database,
// bluetooth modules and also the necessary channels to communicate between
// go routines
package app

import (
	"github.com/thinkty/warden/internal/bluetooth"
	"github.com/thinkty/warden/internal/database"
	"github.com/thinkty/warden/internal/router"
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

// Initialize the blueooth modules and database and serve the static website
// and APIs through the router
func Start() {
	database.Init()
	go bluetooth.Init() // Running the bluetooth operations on a goroutine
	router.InitAndServe()
}
