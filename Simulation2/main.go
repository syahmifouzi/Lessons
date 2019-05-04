package main

import (
	"fmt"
	"math"
)

// PI of math
const PI float64 = math.Pi

// Robot is Hyper Parameter
type Robot struct {
	leftTyre, rightTyre, head                                                      Coor
	sensor                                                                         []Coor
	facing                                                                         byte
	faceDec, m1, c1, m2, c2, mpx, mpy, tRad, tCir, width, height, side, sideDegree float64
}

// Coor is Coordinate
// Value in meter (m)
type Coor struct {
	x, y float64
}

// Rect is Rect
type Rect struct {
	a, b, c, d Coor
}

// MMap is Rect
type MMap struct {
	r1, r2, r3, r4 Rect
}

// Init value distance in meter(m)
func (bot *Robot) init() {
	(*bot).leftTyre.x = -0.06125
	(*bot).leftTyre.y = 0
	(*bot).rightTyre.x = 0.06125
	(*bot).rightTyre.y = 0
	(*bot).head.x = 0
	(*bot).head.y = 0.09733
	(*bot).width = 0.1225
	(*bot).height = math.Sqrt(math.Pow(0.115, 2) - math.Pow((0.1225/2), 2)) // 0.09733
	(*bot).side = 0.115
	(*bot).sideDegree = (57.82 / 360) * 2 * PI // 1.0091493 rad
	(*bot).tRad = 0.02375
	(*bot).tCir = 2 * PI * 0.02375 // 0.14923
	(*bot).facing = 'W'
	(*bot).faceDec = 0
	(*bot).m1 = 0
	(*bot).c1 = 0
	(*bot).m2 = math.Inf(0)
	(*bot).c2 = math.NaN()
	(*bot).mpx = 0
	(*bot).mpy = 0
	(*bot).sensor = make([]Coor, 7)

	for i := range bot.sensor {
		(*bot).sensor[i].y = 0.09733
	}
	(*bot).sensor[0].x = -0.0465
	(*bot).sensor[1].x = -0.031
	(*bot).sensor[2].x = -0.0155
	(*bot).sensor[3].x = 0
	(*bot).sensor[4].x = 0.0155
	(*bot).sensor[5].x = 0.031
	(*bot).sensor[6].x = 0.0465
}

func (bot *Robot) headPos() {
	bot.getFacingDirection()
	s := bot.side
	m := bot.m1
	t := PI - bot.sideDegree - math.Abs(math.Atan(m))
	xComp := s * (math.Cos(t))
	yComp := s * (math.Sin(t))

	tSen := math.Abs(math.Atan(m))
	tXC1 := 0.0155 * math.Cos(tSen)
	tXC2 := 2 * 0.0155 * math.Cos(tSen)
	tXC3 := 3 * 0.0155 * math.Cos(tSen)
	tYC1 := 0.0155 * math.Sin(tSen)
	tYC2 := 2 * 0.0155 * math.Sin(tSen)
	tYC3 := 3 * 0.0155 * math.Sin(tSen)

	f := bot.facing
	switch f {
	case 'W':
		(*bot).head.x = bot.leftTyre.x + (bot.width / 2)
		(*bot).head.y = bot.leftTyre.y + bot.height
		(*bot).faceDec = 0
		for i := range bot.sensor {
			(*bot).sensor[i].y = bot.head.y
		}
		(*bot).sensor[0].x = bot.head.x - (3 * 0.0155)
		(*bot).sensor[1].x = bot.head.x - (2 * 0.0155)
		(*bot).sensor[2].x = bot.head.x - 0.0155
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[4].x = bot.head.x + 0.0155
		(*bot).sensor[5].x = bot.head.x + (2 * 0.0155)
		(*bot).sensor[6].x = bot.head.x + (3 * 0.0155)
	case 'A':
		(*bot).head.x = bot.leftTyre.x - bot.height
		(*bot).head.y = bot.leftTyre.y + (bot.width / 2)
		(*bot).faceDec = 6
		for i := range bot.sensor {
			(*bot).sensor[i].x = bot.head.x
		}
		(*bot).sensor[0].y = bot.head.y - (3 * 0.0155)
		(*bot).sensor[1].y = bot.head.y - (2 * 0.0155)
		(*bot).sensor[2].y = bot.head.y - 0.0155
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].y = bot.head.y + 0.0155
		(*bot).sensor[5].y = bot.head.y + (2 * 0.0155)
		(*bot).sensor[6].y = bot.head.y + (3 * 0.0155)
	case 'X':
		(*bot).head.x = bot.leftTyre.x - (bot.width / 2)
		(*bot).head.y = bot.leftTyre.y - bot.height
		(*bot).faceDec = 4
		for i := range bot.sensor {
			(*bot).sensor[i].y = bot.head.y
		}
		(*bot).sensor[0].x = bot.head.x + (3 * 0.0155)
		(*bot).sensor[1].x = bot.head.x + (2 * 0.0155)
		(*bot).sensor[2].x = bot.head.x + 0.0155
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[4].x = bot.head.x - 0.0155
		(*bot).sensor[5].x = bot.head.x - (2 * 0.0155)
		(*bot).sensor[6].x = bot.head.x - (3 * 0.0155)
	case 'D':
		(*bot).head.x = bot.leftTyre.x + bot.height
		(*bot).head.y = bot.leftTyre.y - (bot.width / 2)
		(*bot).faceDec = 2
		for i := range bot.sensor {
			(*bot).sensor[i].x = bot.head.x
		}
		(*bot).sensor[0].y = bot.head.y + (3 * 0.0155)
		(*bot).sensor[1].y = bot.head.y + (2 * 0.0155)
		(*bot).sensor[2].y = bot.head.y + 0.0155
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].y = bot.head.y - 0.0155
		(*bot).sensor[5].y = bot.head.y - (2 * 0.0155)
		(*bot).sensor[6].y = bot.head.y - (3 * 0.0155)
	case 'Q':
		(*bot).head.x = bot.leftTyre.x - xComp
		(*bot).head.y = bot.leftTyre.y + yComp
		(*bot).faceDec = 7
		(*bot).sensor[0].x = bot.head.x - tXC3
		(*bot).sensor[0].y = bot.head.y - tYC3
		(*bot).sensor[1].x = bot.head.x - tXC2
		(*bot).sensor[1].y = bot.head.y - tYC2
		(*bot).sensor[2].x = bot.head.x - tXC1
		(*bot).sensor[2].y = bot.head.y - tYC1
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].x = bot.head.x + tXC1
		(*bot).sensor[4].y = bot.head.y + tYC1
		(*bot).sensor[5].x = bot.head.x + tXC2
		(*bot).sensor[5].y = bot.head.y + tYC2
		(*bot).sensor[6].x = bot.head.x + tXC3
		(*bot).sensor[6].y = bot.head.y + tYC3
	case 'E':
		(*bot).head.x = bot.rightTyre.x + xComp
		(*bot).head.y = bot.rightTyre.y + yComp
		(*bot).faceDec = 1
		(*bot).sensor[0].x = bot.head.x - tXC3
		(*bot).sensor[0].y = bot.head.y + tYC3
		(*bot).sensor[1].x = bot.head.x - tXC2
		(*bot).sensor[1].y = bot.head.y + tYC2
		(*bot).sensor[2].x = bot.head.x - tXC1
		(*bot).sensor[2].y = bot.head.y + tYC1
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].x = bot.head.x + tXC1
		(*bot).sensor[4].y = bot.head.y - tYC1
		(*bot).sensor[5].x = bot.head.x + tXC2
		(*bot).sensor[5].y = bot.head.y - tYC2
		(*bot).sensor[6].x = bot.head.x + tXC3
		(*bot).sensor[6].y = bot.head.y - tYC3
	case 'Z':
		(*bot).head.x = bot.rightTyre.x - xComp
		(*bot).head.y = bot.rightTyre.y - yComp
		(*bot).faceDec = 5
		(*bot).sensor[0].x = bot.head.x + tXC3
		(*bot).sensor[0].y = bot.head.y - tYC3
		(*bot).sensor[1].x = bot.head.x + tXC2
		(*bot).sensor[1].y = bot.head.y - tYC2
		(*bot).sensor[2].x = bot.head.x + tXC1
		(*bot).sensor[2].y = bot.head.y - tYC1
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].x = bot.head.x - tXC1
		(*bot).sensor[4].y = bot.head.y + tYC1
		(*bot).sensor[5].x = bot.head.x - tXC2
		(*bot).sensor[5].y = bot.head.y + tYC2
		(*bot).sensor[6].x = bot.head.x - tXC3
		(*bot).sensor[6].y = bot.head.y + tYC3
	case 'C':
		(*bot).head.x = bot.leftTyre.x + xComp
		(*bot).head.y = bot.leftTyre.y - yComp
		(*bot).faceDec = 3
		(*bot).sensor[0].x = bot.head.x + tXC3
		(*bot).sensor[0].y = bot.head.y + tYC3
		(*bot).sensor[1].x = bot.head.x + tXC2
		(*bot).sensor[1].y = bot.head.y + tYC2
		(*bot).sensor[2].x = bot.head.x + tXC1
		(*bot).sensor[2].y = bot.head.y + tYC1
		(*bot).sensor[3].x = bot.head.x
		(*bot).sensor[3].y = bot.head.y
		(*bot).sensor[4].x = bot.head.x - tXC1
		(*bot).sensor[4].y = bot.head.y - tYC1
		(*bot).sensor[5].x = bot.head.x - tXC2
		(*bot).sensor[5].y = bot.head.y - tYC2
		(*bot).sensor[6].x = bot.head.x - tXC3
		(*bot).sensor[6].y = bot.head.y - tYC3
	}
}

func (bot *Robot) errDeg() float64 {
	m := bot.m2
	tComp := math.Abs(math.Atan(m))
	t := (PI / 2) - tComp

	return t
	// f := bot.facing
	// if
}

func (bot *Robot) moveBot(l, r float64) {
	if l == r {
		if l == 0 {
			return
		} else if l > 0 {
			// fmt.Println("moved W")
			bot.moveW(l)
		} else {
			// fmt.Println("moved X")
			bot.moveX(l)
		}
	} else if l >= 0 && r >= 0 {
		if l < r {
			// fmt.Println("moved Q")
			bot.moveQ(l, r)
		} else {
			// fmt.Println("moved E")
			bot.moveE(l, r)
		}
	} else if l <= 0 && r <= 0 {
		l = math.Abs(l)
		r = math.Abs(r)
		if l < r {
			// fmt.Println("moved Z")
			bot.moveZ(l, r)
		} else {
			// fmt.Println("moved C")
			bot.moveC(l, r)
		}
	} else if l < 0 && r > 0 {
		l = math.Abs(l)
		if l < r {
			// fmt.Println("moved NQ")
			bot.moveNQ(l, r)
		} else {
			// fmt.Println("moved NC")
			bot.moveNC(l, r)
		}
	} else if l > 0 && r < 0 {
		r = math.Abs(r)
		if l > r {
			// fmt.Println("moved NE")
			bot.moveNE(l, r)
		} else {
			// fmt.Println("moved NZ")
			bot.moveNZ(l, r)
		}
	} else {
		fmt.Println("Unknown rotational: at bot.moveBot()")
		fmt.Println("Left Motor:", l)
		fmt.Println("Right Motor:", r)
	}
}

func (bot *Robot) moveQ(rotL, rotR float64) {

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// fmt.Println("RmL:", RmL)
		// fmt.Println("k:", k)
		// fmt.Println("RmL is bigger than k")
		// fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// fmt.Println("new RmL:", RmL)
	}

	// fmt.Println("proceeding...")
	// fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 - p)
	rS := rB - w
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// // fmt.Println("tDegree:", tDeg)
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x - nS
		(*bot).leftTyre.y = bot.leftTyre.y + hS
		(*bot).rightTyre.x = bot.rightTyre.x - nB
		(*bot).rightTyre.y = bot.rightTyre.y + hB
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nS
		(*bot).rightTyre.x = bot.rightTyre.x - hB
		(*bot).rightTyre.y = bot.rightTyre.y - nB
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x + nS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hS
		(*bot).rightTyre.x = bot.rightTyre.x + nB
		(*bot).rightTyre.y = bot.rightTyre.y - hB
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nS
		(*bot).rightTyre.x = bot.rightTyre.x + hB
		(*bot).rightTyre.y = bot.rightTyre.y + nB
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB + yCompB
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB - yCompB
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB - yCompB
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB + yCompB
		(*bot).headPos()
	}
}

func (bot *Robot) moveE(rotL, rotR float64) {

	// switch the rotation for easier mirroring the value
	tempL := rotL
	tempR := rotR
	rotL = tempR
	rotR = tempL

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// fmt.Println("RmL:", RmL)
		// fmt.Println("k:", k)
		// fmt.Println("RmL is bigger than k")
		// fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// fmt.Println("new RmL:", RmL)
	}

	// fmt.Println("proceeding...")
	// fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 - p)
	rS := rB - w
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// // fmt.Println("tDegree:", tDeg)
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nB
		(*bot).leftTyre.y = bot.leftTyre.y + hB
		(*bot).rightTyre.x = bot.rightTyre.x + nS
		(*bot).rightTyre.y = bot.rightTyre.y + hS
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nB
		(*bot).rightTyre.x = bot.rightTyre.x - hS
		(*bot).rightTyre.y = bot.rightTyre.y + nS
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hB
		(*bot).rightTyre.x = bot.rightTyre.x - nS
		(*bot).rightTyre.y = bot.rightTyre.y - hS
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nB
		(*bot).rightTyre.x = bot.rightTyre.x + hS
		(*bot).rightTyre.y = bot.rightTyre.y - nS
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS + yCompS
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS - yCompS
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS - yCompS
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS + yCompS
		(*bot).headPos()
	}
}

func (bot *Robot) moveZ(rotL, rotR float64) {

	// Modify from moveE
	// becoz they look the same
	// only small different that need to change

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// fmt.Println("RmL:", RmL)
		// fmt.Println("k:", k)
		// fmt.Println("RmL is bigger than k")
		// fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// fmt.Println("new RmL:", RmL)
	}

	// fmt.Println("proceeding...")
	// fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 - p)
	rS := rB - w
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x - nS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hS
		(*bot).rightTyre.x = bot.rightTyre.x - nB
		(*bot).rightTyre.y = bot.rightTyre.y - hB
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x + hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nS
		(*bot).rightTyre.x = bot.rightTyre.x + hB
		(*bot).rightTyre.y = bot.rightTyre.y - nB
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x + nS
		(*bot).leftTyre.y = bot.leftTyre.y + hS
		(*bot).rightTyre.x = bot.rightTyre.x + nB
		(*bot).rightTyre.y = bot.rightTyre.y + hB
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x - hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nS
		(*bot).rightTyre.x = bot.rightTyre.x - hB
		(*bot).rightTyre.y = bot.rightTyre.y + nB
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB - yCompB
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB + yCompB
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB + yCompB
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB - yCompB
		(*bot).headPos()
	}
}

func (bot *Robot) moveC(rotL, rotR float64) {

	// switch the rotation for easier mirroring the value
	tempL := rotL
	tempR := rotR
	rotL = tempR
	rotR = tempL

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// fmt.Println("RmL:", RmL)
		// fmt.Println("k:", k)
		// fmt.Println("RmL is bigger than k")
		// fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// fmt.Println("new RmL:", RmL)
	}

	// // fmt.Println("proceeding...")
	// // fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 - p)
	rS := rB - w
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hB
		(*bot).rightTyre.x = bot.rightTyre.x + nS
		(*bot).rightTyre.y = bot.rightTyre.y - hS
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x + hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nB
		(*bot).rightTyre.x = bot.rightTyre.x + hS
		(*bot).rightTyre.y = bot.rightTyre.y + nS
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nB
		(*bot).leftTyre.y = bot.leftTyre.y + hB
		(*bot).rightTyre.x = bot.rightTyre.x - nS
		(*bot).rightTyre.y = bot.rightTyre.y + hS
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x - hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nB
		(*bot).rightTyre.x = bot.rightTyre.x - hS
		(*bot).rightTyre.y = bot.rightTyre.y - nS
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS - yCompS
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS + yCompS
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS + yCompS
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS - yCompS
		(*bot).headPos()
	}
}

func (bot *Robot) moveNQ(rotL, rotR float64) {

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// // fmt.Println("RmL:", RmL)
		// // fmt.Println("k:", k)
		// // fmt.Println("RmL is bigger than k")
		// // fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// // fmt.Println("new RmL:", RmL)
	}

	// // fmt.Println("proceeding...")
	// // fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (p + 1)
	rS := w - rB
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// // fmt.Println("tdeg", tDeg)
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nS
		(*bot).leftTyre.y = bot.leftTyre.y - hS
		(*bot).rightTyre.x = bot.rightTyre.x - nB
		(*bot).rightTyre.y = bot.rightTyre.y + hB
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x + hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nS
		(*bot).rightTyre.x = bot.rightTyre.x - hB
		(*bot).rightTyre.y = bot.rightTyre.y - nB
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + hS
		(*bot).rightTyre.x = bot.rightTyre.x + nB
		(*bot).rightTyre.y = bot.rightTyre.y - hB
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x - hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nS
		(*bot).rightTyre.x = bot.rightTyre.x + hB
		(*bot).rightTyre.y = bot.rightTyre.y + nB
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB + yCompB
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB - yCompB
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB - yCompB
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB + yCompB
		(*bot).headPos()
	}
}

func (bot *Robot) moveNE(rotL, rotR float64) {

	// switch the rotation for easier mirroring the value
	tempL := rotL
	tempR := rotR
	rotL = tempR
	rotR = tempL

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// // fmt.Println("RmL:", RmL)
		// // fmt.Println("k:", k)
		// // fmt.Println("RmL is bigger than k")
		// // fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// // fmt.Println("new RmL:", RmL)
	}

	// // fmt.Println("proceeding...")
	// // fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 + p)
	rS := w - rB
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nB
		(*bot).leftTyre.y = bot.leftTyre.y + hB
		(*bot).rightTyre.x = bot.rightTyre.x - nS
		(*bot).rightTyre.y = bot.rightTyre.y - hS
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nB
		(*bot).rightTyre.x = bot.rightTyre.x + hS
		(*bot).rightTyre.y = bot.rightTyre.y - nS
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hB
		(*bot).rightTyre.x = bot.rightTyre.x + nS
		(*bot).rightTyre.y = bot.rightTyre.y + hS
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nB
		(*bot).rightTyre.x = bot.rightTyre.x - hS
		(*bot).rightTyre.y = bot.rightTyre.y + nS
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS - yCompS
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS + yCompS
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS + yCompS
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS - yCompS
		(*bot).headPos()
	}
}

func (bot *Robot) moveNC(rotL, rotR float64) {

	// switch the rotation for easier mirroring the value
	tempL := rotL
	tempR := rotR
	rotL = tempR
	rotR = tempL

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// // fmt.Println("RmL:", RmL)
		// // fmt.Println("k:", k)
		// // fmt.Println("RmL is bigger than k")
		// // fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// // fmt.Println("new RmL:", RmL)
	}

	// // fmt.Println("proceeding...")
	// // fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (p + 1)
	rS := w - rB
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// fmt.Println("tdeg", tDeg)
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nB
		(*bot).leftTyre.y = bot.leftTyre.y - hB
		(*bot).rightTyre.x = bot.rightTyre.x - nS
		(*bot).rightTyre.y = bot.rightTyre.y + hS
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x + hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nB
		(*bot).rightTyre.x = bot.rightTyre.x - hS
		(*bot).rightTyre.y = bot.rightTyre.y - nS
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + hB
		(*bot).rightTyre.x = bot.rightTyre.x + nS
		(*bot).rightTyre.y = bot.rightTyre.y - hS
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x - hB // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nB
		(*bot).rightTyre.x = bot.rightTyre.x + hS
		(*bot).rightTyre.y = bot.rightTyre.y + nS
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS + yCompS
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB + xCompB
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS - xCompS
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNS - yCompS
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB + yCompB
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS - yCompS
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNB - xCompB
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNB - yCompB
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNS + xCompS
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNS + yCompS
		(*bot).headPos()
	}
}

func (bot *Robot) moveNZ(rotL, rotR float64) {

	w := bot.width
	k1 := (w * 2 * PI * 90) / 360 // arc of big radius
	k := k1 / bot.tCir            // max of rotation to reach 90 degree
	RmL := rotR - rotL

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	for RmL >= k {
		//scaling down the rotation
		// fmt.Println("RmL:", RmL)
		// fmt.Println("k:", k)
		// fmt.Println("RmL is bigger than k")
		// fmt.Println("scaling down")

		rotR = rotR / 2
		rotL = rotL / 2

		RmL = rotR - rotL
		// fmt.Println("new RmL:", RmL)
	}

	// fmt.Println("proceeding...")
	// fmt.Println("")

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 + p)
	rS := w - rB
	// tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// t := (tDeg/360) * (2*main.Pi) // convert to radian
	t := arcB / rB         // this in radian directly
	hS := rS * math.Sin(t) // this is y component
	hB := rB * math.Sin(t) // this is y component
	mS := rS * math.Cos(t)
	mB := rB * math.Cos(t)
	nS := rS - mS // this is x component
	nB := rB - mB // this is x component

	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.x = bot.leftTyre.x + nS
		(*bot).leftTyre.y = bot.leftTyre.y + hS
		(*bot).rightTyre.x = bot.rightTyre.x - nB
		(*bot).rightTyre.y = bot.rightTyre.y - hB
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y + nS
		(*bot).rightTyre.x = bot.rightTyre.x + hB
		(*bot).rightTyre.y = bot.rightTyre.y - nB
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.x = bot.leftTyre.x - nS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - hS
		(*bot).rightTyre.x = bot.rightTyre.x + nB
		(*bot).rightTyre.y = bot.rightTyre.y + hB
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + hS // refer diagram, the value is different
		(*bot).leftTyre.y = bot.leftTyre.y - nS
		(*bot).rightTyre.x = bot.rightTyre.x - hB
		(*bot).rightTyre.y = bot.rightTyre.y + nB
		(*bot).headPos()
	case 'Q':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is the arctan angle
		tauOp := (PI / 2) - tau       // the angle relative to hB || hS
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB - yCompB
		(*bot).headPos()
	case 'Z':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y + yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB + xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB + yCompB
		(*bot).headPos()
	case 'C':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree // refer diagram, the angle is same (Z method) as arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(tau)
		yCompNS := nS * math.Sin(tau)
		xCompNB := nB * math.Cos(tau)
		yCompNB := nB * math.Sin(tau)
		xCompB := hB * math.Cos(tauOp)
		yCompB := hB * math.Sin(tauOp)
		xCompS := hS * math.Cos(tauOp)
		yCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS - yCompS
		(*bot).rightTyre.x = bot.rightTyre.x + xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB + yCompB
		(*bot).headPos()
	case 'E':
		m := bot.m1
		tau := math.Abs(math.Atan(m)) // tilted degree
		tau = (PI / 2) - tau          // refer diagram, the angle is 90 deg minus arctan angle
		tauOp := (PI / 2) - tau
		// // fmt.Println("tauOp degree:", ((tauOp / (2 * PI)) * 360))
		yCompNS := nS * math.Cos(tau)
		xCompNS := nS * math.Sin(tau)
		yCompNB := nB * math.Cos(tau)
		xCompNB := nB * math.Sin(tau)
		yCompB := hB * math.Cos(tauOp)
		xCompB := hB * math.Sin(tauOp)
		yCompS := hS * math.Cos(tauOp)
		xCompS := hS * math.Sin(tauOp)
		(*bot).leftTyre.x = bot.leftTyre.x + xCompNS + xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y + yCompNB - yCompB
		(*bot).headPos()
	}
}

// Move base on Number of Tyres' rotation
func (bot *Robot) moveW(rot float64) {
	d := rot * bot.tCir
	// fmt.Println("Moving", d, "meters")
	// fmt.Println("")
	m := bot.m2
	t := math.Abs(math.Atan(m))
	xComp := d * (math.Cos(t))
	yComp := d * (math.Sin(t))
	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.y = bot.leftTyre.y + d
		(*bot).rightTyre.y = bot.rightTyre.y + d
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.y = bot.leftTyre.y - d
		(*bot).rightTyre.y = bot.rightTyre.y - d
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - d
		(*bot).rightTyre.x = bot.rightTyre.x - d
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + d
		(*bot).rightTyre.x = bot.rightTyre.x + d
		(*bot).headPos()
	case 'Q':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).headPos()
	case 'E':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).headPos()
	case 'Z':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).headPos()
	case 'C':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).headPos()
	}
}

func (bot *Robot) moveX(rot float64) {
	d := rot * bot.tCir
	// fmt.Println("Moving", d, "meters")
	// fmt.Println("")
	m := bot.m2
	t := math.Abs(math.Atan(m))
	xComp := d * (math.Cos(t))
	yComp := d * (math.Sin(t))
	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.y = bot.leftTyre.y - d
		(*bot).rightTyre.y = bot.rightTyre.y - d
		(*bot).headPos()
	case 'X':
		(*bot).leftTyre.y = bot.leftTyre.y + d
		(*bot).rightTyre.y = bot.rightTyre.y + d
		(*bot).headPos()
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x + d
		(*bot).rightTyre.x = bot.rightTyre.x + d
		(*bot).headPos()
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x - d
		(*bot).rightTyre.x = bot.rightTyre.x - d
		(*bot).headPos()
	case 'Q':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).headPos()
	case 'E':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).headPos()
	case 'Z':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).headPos()
	case 'C':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).headPos()
	}
}

func (bot *Robot) printBotCoor(m string) {
	y2 := bot.rightTyre.y
	y1 := bot.leftTyre.y
	x2 := bot.rightTyre.x
	x1 := bot.leftTyre.x
	grad := (y2 - y1) / (x2 - x1)
	// fmt.Println(m)
	til := math.Atan(grad) // tilted degree
	fmt.Printf("robot facing direction: %c\n", bot.facing)
	fmt.Println("robot gradient:", bot.m1)
	fmt.Println("robot tilted degree:", ((til / (2 * PI)) * 360))
	fmt.Println("leftTyre: (", bot.leftTyre.x, ",", bot.leftTyre.y, ")")
	fmt.Println("rightTyre: (", bot.rightTyre.x, ",", bot.rightTyre.y, ")")
	fmt.Println("head: (", bot.head.x, ",", bot.head.y, ")")
	fmt.Println("")
}

func (bot *Robot) getFacingDirection() {
	if bot.leftTyre.y == bot.rightTyre.y {
		if bot.leftTyre.x < bot.rightTyre.x {
			// // fmt.Println("Facing W")
			// // fmt.Println("-------------------------------------------")
			(*bot).facing = 'W'
		} else {
			// // fmt.Println("Facing X")
			// // fmt.Println("-------------------------------------------")
			(*bot).facing = 'X'
		}
	} else if bot.leftTyre.x == bot.rightTyre.x {
		if bot.leftTyre.y < bot.rightTyre.y {
			// // fmt.Println("Facing A")
			// // fmt.Println("---------------------------------------------")
			(*bot).facing = 'A'
		} else {
			// // fmt.Println("Facing D")
			// // fmt.Println("--------------------------------------------")
			(*bot).facing = 'D'
		}
	} else {
		(*bot).getGradient()
	}
}

func (bot *Robot) getGradient() {
	m1 := (bot.rightTyre.y - bot.leftTyre.y) / (bot.rightTyre.x - bot.leftTyre.x)
	c1 := bot.rightTyre.y - (m1 * bot.rightTyre.x)

	mpxv := (bot.leftTyre.x + bot.rightTyre.x) / 2
	mpyv := (bot.leftTyre.y + bot.rightTyre.y) / 2

	var mpxc, mpyc float64

	if bot.leftTyre.x < bot.rightTyre.x {
		mpxc = bot.leftTyre.x + mpxv
		if m1 > 0 {
			// // fmt.Println("Facing Q")
			// // fmt.Println("----------------------------------------------------")
			(*bot).facing = 'Q'
		} else {
			// // fmt.Println("Facing E")
			// // fmt.Println("-----------------------------------------------------")
			(*bot).facing = 'E'
		}
	} else {
		mpxc = bot.rightTyre.x + mpxv
		if m1 > 0 {
			// // fmt.Println("Facing C")
			// // fmt.Println("------------------------------------------------------")
			(*bot).facing = 'C'
		} else {
			// // fmt.Println("Facing Z")
			// // fmt.Println("-------------------------------------------------------")
			(*bot).facing = 'Z'
		}
	}

	if bot.leftTyre.y < bot.rightTyre.y {
		mpyc = bot.leftTyre.y + mpyv
	} else {
		mpyc = bot.rightTyre.y + mpyv
	}

	m2 := (-1) / m1
	c2 := mpyc - (m2 * mpxc)

	// // fmt.Println("update m1:", m1)
	(*bot).m1 = m1
	(*bot).c1 = c1
	(*bot).m2 = m2
	(*bot).c2 = c2
	(*bot).mpx = mpxc
	(*bot).mpy = mpyc
}

func isOnLine(s1 Coor, r1 Rect) bool {
	/*
			https://stackoverflow.com/questions/2752725/finding-whether-a-point-lies-inside-a-rectangle-or-not/28979027

			var AB = vector(r.A, r.B);
		    var AM = vector(r.A, m);
		    var BC = vector(r.B, r.C);
		    var BM = vector(r.B, m);
		    var dotABAM = dot(AB, AM);
		    var dotABAB = dot(AB, AB);
		    var dotBCBM = dot(BC, BM);
		    var dotBCBC = dot(BC, BC);
		    return 0 <= dotABAM && dotABAM <= dotABAB && 0 <= dotBCBM && dotBCBM <= dotBCBC;
	*/

	AB := vector(r1.a, r1.b)
	AM := vector(r1.a, s1)
	BC := vector(r1.b, r1.c)
	BM := vector(r1.b, s1)
	dotABAM := dotVec(AB, AM)
	dotABAB := dotVec(AB, AB)
	dotBCBM := dotVec(BC, BM)
	dotBCBC := dotVec(BC, BC)

	if 0 <= dotABAM && dotABAM <= dotABAB && 0 <= dotBCBM && dotBCBM <= dotBCBC {
		return true
	}

	return false

}

func vector(p1, p2 Coor) Coor {
	r := Coor{}
	r.x = p2.x - p1.x
	r.y = p2.y - p1.y

	return r
}

func dotVec(u, v Coor) float64 {
	return u.x*v.x + u.y*v.y
}

func (r *MMap) initMap() {

	r.r1.a.x = -0.0115
	r.r1.a.y = 10
	r.r1.b.x = 0.0115
	r.r1.b.y = 10
	r.r1.c.x = 0.0115
	r.r1.c.y = 0
	r.r1.d.x = -0.0115
	r.r1.d.y = 0

	r.r2.a.x = -0.0115
	r.r2.a.y = 10
	r.r2.b.x = 9.9885
	r.r2.b.y = 10
	r.r2.c.x = 9.9885
	r.r2.c.y = 9.977
	r.r2.d.x = -0.0115
	r.r2.d.y = 9.977

	r.r3.a.x = 9.9655
	r.r3.a.y = 10
	r.r3.b.x = 9.9885
	r.r3.b.y = 10
	r.r3.c.x = 9.9885
	r.r3.c.y = 0
	r.r3.d.x = 9.9655
	r.r3.d.y = 0

	r.r4.a.x = -0.0115
	r.r4.a.y = 0.023
	r.r4.b.x = 9.9885
	r.r4.b.y = 0.023
	r.r4.c.x = 9.9885
	r.r4.c.y = 0
	r.r4.d.x = -0.0115
	r.r4.d.y = 0
}

func testMAPPING() {
	bot := Robot{}
	bot.init()

	m := MMap{}
	m.initMap()

	for i := 0; i < 7; i++ {
		if isOnLine(bot.sensor[i], m.r1) || isOnLine(bot.sensor[i], m.r2) || isOnLine(bot.sensor[i], m.r3) || isOnLine(bot.sensor[i], m.r4) {
			fmt.Println("sensor", i, "is online")
		}
	}
}

// ----------------------------------------------------------
// MAIN FUNCTION
// ----------------------------------------------------------

func main() {
	fmt.Println("Start")
	fmt.Println("----------------------------------------------")
	fmt.Println("")

	// startAI()
	testMAPPING()

	// bot := Robot{}
	// bot.init()

	// for i := 0; i < 1000; i++ {
	// 	bot.moveBot(1, 1)
	// }

	// // bot.moveBot(0.407894737, 1.052631579)
	// // bot.moveBot(0.444739, -0.2)

	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)
	// bot.moveBot(1.052631579, 0.407894737)

	// bot.printBotCoor("moved bot")

	fmt.Println("")
	fmt.Println("----------------------------------------------")
	fmt.Println("End")
}
