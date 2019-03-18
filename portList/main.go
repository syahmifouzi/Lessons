package main

import (
	"fmt"
	"log"

	"go.bug.st/serial.v1"
)

// https://godoc.org/go.bug.st/serial.v1

func main() {
	// Retrieve the port list
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}

	// Print the list of detected ports
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}
