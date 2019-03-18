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
		PortName:        "COM13", // in my PC, it is outgoing hc-05 'dev B'
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 8,
	}

	// Open the port
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalln("serial.Open:", err)
	}

	// err = port.Flush()
	// if err != nil {
	// 	log.Fatalln("port.Flush:", err)
	// }

	// Write 2 bytes to the port
	b := []byte{'l', 0}
	_, err = port.Write(b)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	b2 := []byte{'r', 0}
	_, err = port.Write(b2)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	b3 := []byte{'q', 0}
	_, err = port.Write(b3)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	nQuit := 0

	for nQuit != 8 {
		buf := make([]byte, 128)
		n, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("port.Read:", buf[:n])
		nQuit = nQuit + n
	}

	// log.Println("port.Read:", buf[:n], n, buf[0:1])

	// s1 := string(buf[7:8])

	// // log.Println(int(s1[0]))

	// if int(s1[0]) == 0x71 { // 0x71 == 133 == 'q'
	// 	log.Println("i quit")
	// }

	// done := false

	// for !done {
	// 	log.Println("not done", s1)
	// 	time.Sleep(2 * time.Second)
	// 	if s1 == "q" {
	// 		log.Println("i quit")
	// 		done = true
	// 	}
	// }

	// Make sure to close it later
	defer port.Close()
}
