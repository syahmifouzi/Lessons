package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pixel() {
	r := 15
	c := 10
	s := ""
	sum := make([][]float64, r)

	for i := 0; i < r; i++ {
		sum[i] = make([]float64, c)
	}

	for i, v := range sum {
		for i2 := range v {
			sum[i][i2] = 1
		}
	}

	for _, v := range sum {
		for _, v2 := range v {
			ts := strings.Join([]string{"[", strconv.FormatFloat(v2, 'f', -1, 64), "]"}, "")
			s = strings.Join([]string{s, ts}, " ")
		}
		s = strings.Join([]string{s, "\n"}, "")
	}

	fmt.Println(s)
}
