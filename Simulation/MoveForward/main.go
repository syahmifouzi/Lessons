package main

import (
	"fmt"
	"math"
)

// PI of math
const PI float64 = math.Pi

// Robot is Hyper Parameter
type Robot struct {
	leftTyre, rightTyre, head                                 Coor
	facing                                                    byte
	m1, c1, m2, c2, mpx, mpy, tRad, tCir, width, height, side float64
}

// Coor is Coordinate
// Value in meter (m)
type Coor struct {
	x, y float64
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
	(*bot).tRad = 0.02375
	(*bot).tCir = 2 * PI * 0.02375 // 0.14923

	bot.printBotCoor("Initial")
	bot.getFacingDirection()
}

func (bot *Robot) headPos() {
	bot.getFacingDirection()
	h := bot.height
	m := bot.m1
	// fmt.Println("calling m1:", m)
	t := (PI / 2) - math.Abs(math.Atan(m))
	xComp := h * (math.Cos(t))
	yComp := h * (math.Sin(t))
	f := bot.facing
	switch f {
	case 'Q':
		(*bot).head.x = bot.mpx - xComp
		(*bot).head.y = bot.mpy + yComp
	case 'E':
		(*bot).head.x = bot.mpx + xComp
		(*bot).head.y = bot.mpy + yComp
	case 'Z':
		(*bot).head.x = bot.mpx - xComp
		(*bot).head.y = bot.mpy - yComp
	case 'C':
		(*bot).head.x = bot.mpx + xComp
		(*bot).head.y = bot.mpy - yComp
	}

}

func (bot *Robot) moveQ(rotL, rotR float64) {

	w := bot.width
	// k1 := (w * 2 * PI * 90) / 360
	// k := k1 / bot.tCir
	// RmL := rotR - rotL
	// fmt.Println("rml:", RmL)
	// fmt.Println("k:", k)
	// fmt.Println("")

	// RmL MUST BE LOWER THAN k
	// else below eq. is not valid

	arcB := rotR * bot.tCir
	// arcS := rotL * bot.tCir

	// err: p cannot equal to 1
	p := rotL / rotR
	rB := w / (1 - p)
	rS := rB - w
	tDeg := (arcB * 360) / (2 * PI * rB) // this in degree
	// fmt.Println("arcB:", arcB)
	// fmt.Println("arcS:", arcS)
	// fmt.Println("rB:", rB)
	// fmt.Println("rS:", rS)
	fmt.Println("tDeg", tDeg)
	fmt.Println("")
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
		fmt.Println("eq. W")
		fmt.Println("")
		(*bot).leftTyre.x = bot.leftTyre.x - nS
		(*bot).leftTyre.y = bot.leftTyre.y + hS
		(*bot).rightTyre.x = bot.rightTyre.x - nB
		(*bot).rightTyre.y = bot.rightTyre.y + hB
		(*bot).headPos()
	case 'Q':
		fmt.Println("eq. Q")
		fmt.Println("")
		m := bot.m1
		til := math.Abs(math.Atan(m)) // tilted degree
		fmt.Println("tilted degree:", ((til / (2 * PI)) * 360))
		t1 := (PI / 2) - til
		// fmt.Println("t1 degree:", ((t1 / (2 * PI)) * 360))
		xCompNS := nS * math.Cos(til)
		yCompNS := nS * math.Sin(til)
		xCompNB := nB * math.Cos(til)
		yCompNB := nB * math.Sin(til)
		xCompB := hB * math.Cos(t1)
		yCompB := hB * math.Sin(t1)
		xCompS := hS * math.Cos(t1)
		yCompS := hS * math.Sin(t1)
		(*bot).leftTyre.x = bot.leftTyre.x - xCompNS - xCompS
		(*bot).leftTyre.y = bot.leftTyre.y - yCompNS + yCompS
		(*bot).rightTyre.x = bot.rightTyre.x - xCompNB - xCompB
		(*bot).rightTyre.y = bot.rightTyre.y - yCompNB + yCompB
		(*bot).headPos()
	}
}

// Move base on Number of Tyres' rotation
func (bot *Robot) moveForward(rot float64) {
	d := rot * bot.tCir
	fmt.Println("Moving", d, "meters")
	fmt.Println("")
	m := bot.m2
	t := math.Abs(math.Atan(m))
	xComp := d * (math.Cos(t))
	yComp := d * (math.Sin(t))
	f := bot.facing
	switch f {
	case 'W':
		(*bot).leftTyre.y = bot.leftTyre.y + d
		(*bot).rightTyre.y = bot.rightTyre.y + d
		(*bot).head.y = bot.head.y + d
	case 'X':
		(*bot).leftTyre.y = bot.leftTyre.y - d
		(*bot).rightTyre.y = bot.rightTyre.y - d
		(*bot).head.y = bot.head.y - d
	case 'A':
		(*bot).leftTyre.x = bot.leftTyre.x - d
		(*bot).rightTyre.x = bot.rightTyre.x - d
		(*bot).head.x = bot.head.x - d
	case 'D':
		(*bot).leftTyre.x = bot.leftTyre.x + d
		(*bot).rightTyre.x = bot.rightTyre.x + d
		(*bot).head.x = bot.head.x + d
	case 'Q':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).head.x = bot.head.x - xComp
		(*bot).head.y = bot.head.y + yComp
	case 'E':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y + yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y + yComp
		(*bot).head.x = bot.head.x + xComp
		(*bot).head.y = bot.head.y + yComp
	case 'Z':
		(*bot).leftTyre.x = bot.leftTyre.x - xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x - xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).head.x = bot.head.x - xComp
		(*bot).head.y = bot.head.y - yComp
	case 'C':
		(*bot).leftTyre.x = bot.leftTyre.x + xComp
		(*bot).leftTyre.y = bot.leftTyre.y - yComp
		(*bot).rightTyre.x = bot.rightTyre.x + xComp
		(*bot).rightTyre.y = bot.rightTyre.y - yComp
		(*bot).head.x = bot.head.x + xComp
		(*bot).head.y = bot.head.y - yComp
	}
}

func (bot *Robot) printBotCoor(m string) {
	y2 := bot.rightTyre.y
	y1 := bot.leftTyre.y
	x2 := bot.rightTyre.x
	x1 := bot.leftTyre.x
	grad := (y2 - y1) / (x2 - x1)
	til := math.Abs(math.Atan(grad)) // tilted degree
	fmt.Println("coor degree:", ((til / (2 * PI)) * 360))
	// fmt.Println("gradient is:", grad)
	fmt.Println(m)
	fmt.Println("leftTyre: (", bot.leftTyre.x, ",", bot.leftTyre.y, ")")
	fmt.Println("rightTyre: (", bot.rightTyre.x, ",", bot.rightTyre.y, ")")
	fmt.Println("head: (", bot.head.x, ",", bot.head.y, ")")
	fmt.Println("")
}

func (bot *Robot) getFacingDirection() {
	if bot.leftTyre.y == bot.rightTyre.y {
		if bot.leftTyre.x < bot.rightTyre.x {
			fmt.Println("Facing North")
			fmt.Println("")
			(*bot).facing = 'W'
		} else {
			fmt.Println("Facing South")
			fmt.Println("")
			(*bot).facing = 'X'
		}
	} else if bot.leftTyre.x == bot.rightTyre.x {
		if bot.leftTyre.y < bot.rightTyre.y {
			fmt.Println("Facing West")
			fmt.Println("")
			(*bot).facing = 'A'
		} else {
			fmt.Println("Facing East")
			fmt.Println("")
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
			fmt.Println("Facing North West")
			fmt.Println("")
			(*bot).facing = 'Q'
		} else {
			fmt.Println("Facing North East")
			fmt.Println("")
			(*bot).facing = 'E'
		}
	} else {
		mpxc = bot.rightTyre.x + mpxv
		if m1 > 0 {
			fmt.Println("Facing South East")
			fmt.Println("")
			(*bot).facing = 'C'
		} else {
			fmt.Println("Facing South West")
			fmt.Println("")
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

	// fmt.Println("update m1:", m1)
	(*bot).m1 = m1
	(*bot).c1 = c1
	(*bot).m2 = m2
	(*bot).c2 = c2
	(*bot).mpx = mpxc
	(*bot).mpy = mpyc
}

// ----------------------------------------------------------
// MAIN FUNCTION
// ----------------------------------------------------------

func main() {
	fmt.Println("Start")
	fmt.Println("----------------------------------------------")
	fmt.Println("")

	bot := Robot{}
	bot.init()

	// bot.moveForward(13)
	// bot.moveQ(0.5, 0.8)
	bot.moveQ(0.25, 0.4)

	bot.printBotCoor("moved Q")

	bot.moveQ(0.25, 0.4)

	bot.printBotCoor("moved Q")

	til := math.Abs(math.Atan(bot.m1)) // tilted degree
	fmt.Println("ended degree:", ((til / (2 * PI)) * 360))

	fmt.Println("")
	fmt.Println("----------------------------------------------")
	fmt.Println("End")
}
