package main

import (
	"fmt"
	"time"
)

// Wheel ...
type Wheel struct {
	radius    float64
	fractions float64
	speed     time.Duration
	rClose    bool
}

func (w *Wheel) getDistance(n []int) float64 {
	D := 2 * PI * w.radius * float64(n[0])
	d := 2 * PI * w.radius * (float64(n[1]) / w.fractions)
	return D + d
}

func (w *Wheel) rotate(c chan []int) {
	i := 0
	j := 0
	for {
		i++
		if i == int(w.fractions) {
			i = 0
			j++
		}
		if w.rClose {
			c <- []int{j, i}
			close(c)
			break
		}
		time.Sleep(w.speed)
	}
}

func setSpeed(i time.Duration) time.Duration {
	// 255 	-> 80
	// 0	-> 401
	// the graph is y = -1.259x + 401
	return i * time.Microsecond
}

func wheelMain() {
	w := Wheel{
		radius:    0.02375, // diameter == 0.0475 // radius of the wheel
		fractions: 360,     // number of fractions of wheel encoders
		speed:     setSpeed(80),
		rClose:    false,
	}

	c := make(chan []int)

	go w.rotate(c)

	go func() {
		time.Sleep(time.Second)
		w.rClose = true
	}()

	dum := <-c
	fmt.Println("Wheel has rotated", dum[0], "times plus", dum[1], "degree")
	d := w.getDistance(dum)
	fmt.Println("Distance traveled =", d, "m")
}
