// Package for scanning, connecting, and communicating between bluetooth modules
package bluetooth

import (
	"fmt"
	"log"
	"regexp"

	"tinygo.org/x/bluetooth"
)

type Ward struct {
	result bluetooth.ScanResult
	device bluetooth.Device
}

var adapter bluetooth.Adapter
var wards []Ward

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
		log.Panic("Failed to scan bluetooth devices. Make sure bluetooth is enabled...")
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

	// If no local name is given, ignore
	if result.LocalName() == "" {
		return
	}

	// If the local name of the device does not follow the pattern, ignore
	if match, _ := regexp.MatchString("Ward[0-9]+", result.LocalName()); !match {
		return
	}

	// Iterate through the device slice and return if the local name is not unique
	for _, ward := range wards {
		if ward.result.LocalName() == result.LocalName() {
			return
		}

		// If the scanned device has the same name but a different address, it could
		// possibly be that the device has updated its name, replace the old one
		if ward.result.Address.String() == result.Address.String() {
			// TODO: disconnect the old device and connect the new device
			// TODO: append to the slice and return
		}
	}

	// Connect to the device
	device, err := adapter.Connect(result.Address, bluetooth.ConnectionParams{})
	if err != nil {
		log.Panic("Failed to connect to bluetooth device: ", result.LocalName())
	}

	// Append to the device slice
	wards = append(wards, Ward{result: result, device: *device})

	log.Print("Scanned new device...")
	printDevices()
}

func printDevices() {
	fmt.Println("[")
	for _, ward := range wards {
		fmt.Printf("\t%s\t%d\t%s\n", ward.result.Address, ward.result.RSSI, ward.result.LocalName())
	}
	fmt.Println("]")
}
