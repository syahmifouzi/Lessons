package main

import "fmt"

// Robot ...
type Robot struct {
	pose []float64
}

func robotMain() {
	p := Robot{
		pose: []float64{0, 0, PI / 4},
	}

	wl := Wheel{
		radius:    0.02375, // diameter == 0.0475 // radius of the wheel
		fractions: 360,     // number of fractions of wheel encoders
		speed:     setSpeed(80),
		rClose:    false,
	}

	wr := Wheel{
		radius:    0.02375, // diameter == 0.0475 // radius of the wheel
		fractions: 360,     // number of fractions of wheel encoders
		speed:     setSpeed(80),
		rClose:    false,
	}

	fmt.Println(p, wl, wr)
}
