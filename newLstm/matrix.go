package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randArr(min, max, scale float64, r, c int) [][]float64 {
	// we are seeding the rand variable with present time
	// or use crypto/rand for more secure way
	// https://gobyexample.com/random-numbers
	// so that we would get different output each time
	// rand.Seed(time.Now().UnixNano())
	// OR WE CAN JUST CONSTANT IT TO CHECK IF THE MODEL IS WORKING!
	rand.Seed(0)

	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v[i][j] = scale*max*rand.Float64() + min
		}
	}

	return v
}

func randArr2(min, max, scale float64, r, c int) [][]float64 {
	// we are seeding the rand variable with present time
	// or use crypto/rand for more secure way
	// https://gobyexample.com/random-numbers
	// so that we would get different output each time
	// rand.Seed(time.Now().UnixNano())
	// OR WE CAN JUST CONSTANT IT TO CHECK IF THE MODEL IS WORKING!
	// Need to declare once at main function
	// or else it will give the same sequence every time
	// rand.Seed(0)

	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v[i][j] = scale*max*rand.Float64() + min
		}
	}

	return v
}

func zeros(r, c int) [][]float64 {
	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v[i][j] = 0
		}
	}

	return v
}

func tambah(m1, m2 [][]float64) [][]float64 {

	sum := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		sum[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			sum[i][i2] = v2 + m2[i][i2]
		}
	}

	return sum
}

func tolak(m1, m2 [][]float64) [][]float64 {

	difference := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		difference[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			difference[i][i2] = v2 - m2[i][i2]
		}
	}

	return difference
}

func darab(m1, m2 [][]float64) [][]float64 {

	product := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		product[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			product[i][i2] = v2 * m2[i][i2]
		}
	}

	return product
}

func darabN(m1 [][]float64, n float64) [][]float64 {

	product := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		product[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			product[i][i2] = v2 * n
		}
	}

	return product
}

func concatArr(m1, m2 [][]float64) [][]float64 {
	c := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		c[i] = make([]float64, len(m1[0])+len(m2[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			c[i][i2] = v2
		}
	}

	for i, v := range m2 {
		for i2, v2 := range v {
			c[i][len(m1[0])+i2] = v2
		}
	}

	// r = append(r, randomizeValue(outputSize, inputSize))

	return c
}

func tanH(m1 [][]float64) [][]float64 {

	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			output[i][i2] = math.Tanh(v2)
		}
	}

	return output
}

func dot(m1, m2 [][]float64) [][]float64 {

	// Ref 2d slice
	// https://gobyexample.com/slices
	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m2[0]))
	}

	// fmt.Println("l.param.wi", len(m1), len(m1[0]))
	// fmt.Println("diInput", len(m2), len(m2[0]))

	for outR, v := range m1 {
		for outC := range m2[0] {
			output[outR][outC] = 0
			for i2, v2 := range v {
				output[outR][outC] += v2 * m2[i2][outC]
			}
		}
	}

	return output
}

func sigmoid(m1 [][]float64) [][]float64 {

	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			var nX float64
			nX = 0 - v2
			output[i][i2] = 1 / (1 + math.Exp(nX))
		}
	}

	return output
}

func sigmoidDeriv(m1 [][]float64) [][]float64 {

	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			output[i][i2] = v2 * (1 - v2)
		}
	}

	return output
}

func tanHDeriv(m1 [][]float64) [][]float64 {

	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			output[i][i2] = 1 - math.Pow(math.Tanh(v2), 2)
		}
	}

	return output
}

func outer(m1, m2 [][]float64) [][]float64 {

	r := len(m1[0])
	c := len(m2[0])

	output := make([][]float64, r)
	for i := 0; i < r; i++ {
		output[i] = make([]float64, c)
	}

	// fmt.Println("output: ", len(output), len(output[0]))
	// fmt.Println("m1", m1[0])
	// fmt.Println("m2", m2[0])

	for i, v := range m1[0] {
		for i2, v2 := range m2[0] {
			// fmt.Println("status: ", i, i2)
			// fmt.Println("v: ", v)
			// fmt.Println("v2: ", v2)
			// fmt.Println("status output: ", output[i][i2])
			output[i][i2] = v * v2 // v == m1[0][i] && v2 == m2[0][i2]
			// fmt.Println("EXIT")
		}
	}

	// fmt.Println("EXIT")

	return output
}

func transpose(m1 [][]float64) [][]float64 {

	mT := make([][]float64, len(m1[0]))
	for i := 0; i < len(m1[0]); i++ {
		mT[i] = make([]float64, len(m1))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			mT[i2][i] = v2
		}
	}

	return mT
}

func cutArr(m [][]float64, f int) [][]float64 {

	// fmt.Println("m", m)
	o := make([][]float64, len(m))
	for i := 0; i < len(m); i++ {
		o[i] = make([]float64, len(m[0])-f)
	}

	// printShape("o", m)

	for i := f; i < len(m[0]); i++ {
		o[0][i-f] = m[0][i]
	}

	// printShape("o", o)
	// fmt.Println("o", o)

	// os.Exit(3)

	return o
}

func printShape(n string, m [][]float64) {
	fmt.Println(n, "[", len(m), ",", len(m[0]), "]")
}
