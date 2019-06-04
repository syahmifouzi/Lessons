package robot

import (
	"fmt"
	"math"
)

// PI of math
const PI float64 = math.Pi

func distanceMoved(n float64) {
	r := 0.02375 // diameter == 0.0475 m
	// base := 0.1225 // base distance between 2 wheels
	nF := 8.00 // number of fractions

	d := 2 * PI * (n / nF) * r
	fmt.Println("distance =", d)
}
