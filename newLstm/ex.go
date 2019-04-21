package main

import (
	"fmt"
	"math"
	"math/rand"
)

// ToyLossLayer is struct
type ToyLossLayer struct{}

func ex0() {
	// learns to repeat simple sequence from random inputs
	rand.Seed(0)

	// parameters for input data dimension and lstm cell count
	memCellCt := 100
	xDim := 50
	lp := LstmParam{}
	// fmt.Println("lp.xDim B4:", lp.xDim)
	lp.init(memCellCt, xDim)
	// fmt.Println("lp.xDim After:", lp.xDim)
	ln := LstmNetwork{}
	ln.init(lp)
	yList := [][]float64{{-0.5, 0.2, 0.1, -0.5}}
	var inputValArr [][][]float64
	// for i := 0; i < len(yList[0]); i++ {
	// 	inputValArr = append(inputValArr, randArr2(0, 1, 1, 1, xDim))
	// }
	inputValArr = setRand3()
	// fmt.Println("inputValArr", inputValArr)

	for i := 0; i < 100; i++ {
		fmt.Println("iter:", i)
		fmt.Print("y_pred: [")
		for i2 := range yList[0] {
			ln.xListAdd(inputValArr[i2])
			fmt.Print(ln.lstmNodeList[i2].state.h[0][0], ",")
		}
		fmt.Println("]")

		t := ToyLossLayer{}

		loss := ln.yListIs(yList, t)
		println("loss:", loss)
		lr := 0.1
		ln.lstmParam.applyDiff(lr)
		// lp.applyDiff(lr)
		ln.xListClear()
	}
}

func (t *ToyLossLayer) loss(pred [][]float64, label float64) float64 {
	return math.Pow(pred[0][0]-label, 2)
}

func (t *ToyLossLayer) btmDiff(pred [][]float64, label float64) [][]float64 {
	diff := darabN(pred, 0)
	diff[0][0] = 2 * (pred[0][0] - label)
	return diff
}
