package bluetooth

import (
	"fmt"
	"log"

	"tinygo.org/x/bluetooth"
)

var adapter bluetooth.Adapter
var devices []bluetooth.ScanResult

// Initialize the bluetooth interface and scan for devices
func Init() {
	adapter := bluetooth.DefaultAdapter

	err := adapter.Enable()
	if err != nil {
		log.Panic(err)
	}

	log.Print("Scanning for bluetooth devices...")

	// This hangs unless an error has occurred or adapter.StopScan() has been called
	err = adapter.Scan(handleScanResult)
	if err != nil {
		log.Print("Failed to scan bluetooth devices...")
		return
	}

	log.Print("Stopping scan...")
}

// Stop scanning
func StopScan() (error, string) {
	log.Print("Stopping scan...")
	err := adapter.StopScan()
	if err != nil {
		return err, "Failed to stop scan..."
	}

	return nil, ""
}

// Add the result to the devices slice if it is unique
func handleScanResult(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {

	// Iterate through the device slice and return if it is not unique
	for _, device := range devices {
		// If the localname or the address matches, do not add to the slice
		if device.LocalName() == result.LocalName() || !device.Address.IsRandom() && device.Address.String() == result.Address.String() {
			return
		}
	}

	// Append to the device slice
	devices = append(devices, result)

	log.Print("Scanned new device...")
	fmt.Println("[")
	for _, device := range devices {
		fmt.Printf("\t%s\t%s\n", device.Address, device.LocalName())
	}
	fmt.Println("]")
}
