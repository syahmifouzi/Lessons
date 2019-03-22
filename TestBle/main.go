package main

import (
	"log"
	"math"
	"strconv"

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
	var leftM float64 = 255
	var rightM float64 = 255
	// leftM := 100
	// rightM := 100

	leftString := strconv.FormatFloat(math.Round(leftM), 'f', -1, 64)
	rightString := strconv.FormatFloat(math.Round(rightM), 'f', -1, 64)

	// Write 2 bytes to the port
	b1 := []byte{'l'}
	_, err = port.Write(b1)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}
	b11 := []byte(leftString)
	_, err = port.Write(b11)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}

	b2 := []byte{'r'}
	_, err = port.Write(b2)
	if err != nil {
		log.Fatalln("port.WriteRight:", err)
	}
	b21 := []byte(rightString)
	_, err = port.Write(b21)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}

	b3 := []byte{'q'}
	_, err = port.Write(b3)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	// nQuit := 0

	// var splice [][]byte

	// for nQuit != 8 {
	// 	buf := make([]byte, 128)
	// 	n, err := port.Read(buf)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println("port.Read:", buf[:n])
	// 	nQuit = nQuit + n

	// 	for i := 0; i < n; i++ {
	// 		splice = append(splice, buf[i:i+1])
	// 	}
	// }

	// log.Println("splice", splice[3])

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
