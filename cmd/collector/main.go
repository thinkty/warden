// Collector package handles "collecting" data from various bluetooth devices
// and update the database. It also handles user actions sent from the server.
package main

import (
	"github.com/thinkty/warden/internal/database"
	"github.com/thinkty/warden/internal/scanner"
)

func main() {
	database.Init()

	// Add the available modules to the device list
	devices := scanner.Devices{
		List: make([]scanner.Device, 0),
	}
	devices.Scan()

	// TODO: setup go routine to query the devices at an interval

	// router.InitAndServeCollector()
}
