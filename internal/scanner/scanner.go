// Package for scanning and managing devices (Wards)
package scanner

import (
	"log"
	"net"
	"reflect"
)

// Port to send the UDP broadcast to
const ScannerPort = ":48636"
const BroadcastAddr = "192.168.255.255:48637"

// Struct to store IP address and valid bit. If valid bit is set to false, it
// should not be used to request sensor values
type Device struct {
	IP    net.IP
	Valid bool
}

type Devices struct {
	list []Device
	// TODO: add other properties necessary
}

// Add a new IP address to the list
func (devices *Devices) Add(ip net.IP) {
	// TODO: validate that it is a local IP address

	devices.list = append(devices.list, Device{IP: ip, Valid: true})
	log.Printf("Added %s to device list", ip.String())
}

// Remove the address from the list if it exists and is valid
func (devices *Devices) Remove(ip net.IP) {
	length := len(devices.list)

	if length == 0 {
		log.Printf("Cannot remove %s from device list. Length 0", ip.String())
		return
	}

	// Since the list slice won't be big, just traverse and search
	for index, device := range devices.list {
		if device.Valid && reflect.DeepEqual(device.IP, ip) {
			devices.list[index].Valid = false
			log.Printf("Removed %s from device list", ip.String())
			return
		}
	}

	log.Printf("Cannot remove %s from device list. Not found", ip.String())
}

// Scan the local area network using UDP broadcast and add the IP addresses of
// the available wards that respond
func (devices *Devices) Scan() {
	// Setup socket on local address at port 48636
	listenAddr, err := net.ResolveUDPAddr("udp4", ScannerPort)
	if err != nil {
		log.Panic(err)
	}

	conn, err := net.ListenUDP("udp4", listenAddr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	log.Printf("Scanner listening at %s", listenAddr.String())

	broadcastAddr, err := net.ResolveUDPAddr("udp4", BroadcastAddr)
	if err != nil {
		log.Panic(err)
	}

	// Send a message to the broadcast address
	message := []byte("ward")
	_, err = conn.WriteToUDP(message, broadcastAddr)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Dialing to %s", broadcastAddr.String())

	// TODO: is this correct? what if there are multiple responses?
	// Read server response
	buf := make([]byte, 1000)
	n, resAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Read %d bytes from %s : %s", n, resAddr.String(), string(buf[:]))

	// TODO: Parse ip address from response (resAddr.IP) and save to list
}
