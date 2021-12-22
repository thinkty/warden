// Collector package handles "collecting" data from various bluetooth devices
// and update the database. It also handles user actions sent from the server.
package main

import (
	"github.com/thinkty/warden/internal/bluetooth"
	"github.com/thinkty/warden/internal/database"
)

func main() {
	database.Init()
	bluetooth.Init()

	// TODO: handle data creation, and add router to handle user action
}
