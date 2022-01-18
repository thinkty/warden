// Ward package implements a mock version of ward that
package main

import (
	"log"
	"net"
)

const WardAddr = "127.0.0.1:48637"

func main() {
	conn, err := net.ListenPacket("udp4", WardAddr)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	log.Printf("Ward listening at %s", conn.LocalAddr().String())

	buf := make([]byte, 1024)
	n, addr, err := conn.ReadFrom(buf)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%s sent this: %s\n", addr, buf[:n])
}
