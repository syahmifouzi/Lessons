package main

import (
	"log"

	"github.com/jacobsa/go-serial/serial"
)

// https://github.com/jacobsa/go-serial
// https://github.com/tarm/serial
// https://medium.com/@rodzzlessa/serial-communication-with-arduino-c75e48443f11

func main() {
	// Set up options
	options := serial.OpenOptions{
		PortName:        "COM9", // in my PC, it is outgoing hc-05 'dev B'
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 2,
	}

	// Open the port
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalln("serial.Open:", err)
	}

	// Write 2 bytes to the port
	b := []byte{'m', 0x64}
	_, err = port.Write(b)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	buf := make([]byte, 128)
	n, err := port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("port.Read:", buf[:n], n, buf[0:1], buf[1:2])

	s1 := string(buf[1:2])

	// log.Println(int(s1[0]))

	if int(s1[0]) == 0x64 {
		log.Println("i got 100")
	}

	// Make sure to close it later
	defer port.Close()
}
