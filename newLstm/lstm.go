package main

import (
	"fmt"
)

// LstmParam is Parameter
type LstmParam struct {
	memCellCt, xDim                                                                                int
	wg, wi, wf, wo, bg, bi, bf, bo, wgDiff, wiDiff, wfDiff, woDiff, bgDiff, biDiff, bfDiff, boDiff [][]float64
}

// LstmState is Gates
type LstmState struct {
	g, i, f, o, s, h, btmDiffH, btmDiffS [][]float64
}

// LstmNode is Node
type LstmNode struct {
	state            LstmState
	xc, sPrev, hPrev [][]float64
}

// LstmNetwork is Network
type LstmNetwork struct {
	lstmParam    LstmParam
	lstmNodeList []LstmNode
	xList        [][][]float64
}

// var fp int

func (l *LstmParam) init(memCellCt, xDim int) {
	(*l).memCellCt = memCellCt
	(*l).xDim = xDim
	concatLen := xDim + memCellCt
	// weight matrices
	// (*l).wg = randArr(-0.1, 0.1, 2, memCellCt, concatLen)
	(*l).wg = setRand()
	// fmt.Println("(*l).wg", (*l).wg)
	// fmt.Println("")
	// fmt.Println("")
	// (*l).wi = randArr(-0.1, 0.1, 2, memCellCt, concatLen)
	// fmt.Println("(*l).wi", (*l).wi)
	// (*l).wf = randArr(-0.1, 0.1, 2, memCellCt, concatLen)
	// (*l).wo = randArr(-0.1, 0.1, 2, memCellCt, concatLen)
	(*l).wi = setRand()
	(*l).wf = setRand()
	(*l).wo = setRand()
	// bias terms
	// (*l).bg = randArr(-0.1, 0.1, 2, 1, memCellCt)
	// (*l).bi = randArr(-0.1, 0.1, 2, 1, memCellCt)
	// (*l).bf = randArr(-0.1, 0.1, 2, 1, memCellCt)
	// (*l).bo = randArr(-0.1, 0.1, 2, 1, memCellCt)
	(*l).bg = setRand2()
	(*l).bi = setRand2()
	(*l).bf = setRand2()
	(*l).bo = setRand2()
	// fmt.Println("(*l).bg", (*l).bg)
	// diffs (derivative of loss function w.r.t. all parameters)
	(*l).wgDiff = zeros(memCellCt, concatLen)
	(*l).wiDiff = zeros(memCellCt, concatLen)
	(*l).wfDiff = zeros(memCellCt, concatLen)
	(*l).woDiff = zeros(memCellCt, concatLen)
	(*l).bgDiff = zeros(1, memCellCt)
	(*l).biDiff = zeros(1, memCellCt)
	(*l).bfDiff = zeros(1, memCellCt)
	(*l).boDiff = zeros(1, memCellCt)
}

func (l *LstmParam) applyDiff(lr float64) {
	// if l.wgDiff[0][0] != 0 {
	// 	fmt.Println("l.wgDiff", l.wgDiff)
	// }

	(*l).wg = tolak(l.wg, darabN(l.wgDiff, lr))
	(*l).wi = tolak(l.wi, darabN(l.wiDiff, lr))
	(*l).wf = tolak(l.wf, darabN(l.wfDiff, lr))
	(*l).wo = tolak(l.wo, darabN(l.woDiff, lr))
	(*l).bg = tolak(l.bg, darabN(l.bgDiff, lr))
	(*l).bi = tolak(l.bi, darabN(l.biDiff, lr))
	(*l).bf = tolak(l.bf, darabN(l.bfDiff, lr))
	(*l).bo = tolak(l.bo, darabN(l.boDiff, lr))
	// reset diffs to zero
	(*l).wgDiff = darabN(l.wgDiff, 0)
	(*l).wiDiff = darabN(l.wiDiff, 0)
	(*l).wfDiff = darabN(l.wfDiff, 0)
	(*l).woDiff = darabN(l.woDiff, 0)
	(*l).bgDiff = darabN(l.bgDiff, 0)
	(*l).biDiff = darabN(l.biDiff, 0)
	(*l).bfDiff = darabN(l.bfDiff, 0)
	(*l).boDiff = darabN(l.boDiff, 0)
}

func (l *LstmState) init(memCellCt int) {
	(*l).g = zeros(1, memCellCt)
	(*l).i = zeros(1, memCellCt)
	(*l).f = zeros(1, memCellCt)
	(*l).o = zeros(1, memCellCt)
	(*l).s = zeros(1, memCellCt)
	(*l).h = zeros(1, memCellCt)
	(*l).btmDiffH = darabN(l.h, 0)
	(*l).btmDiffS = darabN(l.s, 0)
}

func (l *LstmNode) init(lp LstmParam, ls LstmState) {
	// store reference to parameters and to activations
	(*l).state = ls
	// non-recurrent input concatenated with recurrent input
	// (*l).xc = zeros(1, 1)
}

func (l *LstmNetwork) btmDataIs(x, sPrev, hPrev [][]float64, idx int) {
	// if this is the first lstm node in the network
	if sPrev == nil {
		// fmt.Println("no recurrent inputs yet")
		sPrev = darabN(l.lstmNodeList[idx].state.s, 0)
	}
	if hPrev == nil {
		// fmt.Println("no recurrent inputs yet")
		hPrev = darabN(l.lstmNodeList[idx].state.h, 0)
	}
	// save data for use in backdrop
	(*l).lstmNodeList[idx].sPrev = sPrev
	(*l).lstmNodeList[idx].hPrev = hPrev
	// fmt.Println("(*l).sPrev", (*l).sPrev)
	// fmt.Println("(*l).hPrev", (*l).hPrev)

	// concatenate x(t) and h(t-1)
	xc := concatArr(x, hPrev)
	// fmt.Println("xc", xc)
	// fmt.Println("dot(l.param.wg, transpose(xc))", dot(l.param.wg, transpose(xc)))
	// fmt.Println("l.param.bg", len(l.param.bg), ":", len(l.param.bg[0]))
	(*l).lstmNodeList[idx].state.g = tanH(tambah(transpose(dot(l.lstmParam.wg, transpose(xc))), l.lstmParam.bg))
	(*l).lstmNodeList[idx].state.i = sigmoid(tambah(transpose(dot(l.lstmParam.wi, transpose(xc))), l.lstmParam.bi))
	(*l).lstmNodeList[idx].state.f = sigmoid(tambah(transpose(dot(l.lstmParam.wf, transpose(xc))), l.lstmParam.bf))
	(*l).lstmNodeList[idx].state.o = sigmoid(tambah(transpose(dot(l.lstmParam.wo, transpose(xc))), l.lstmParam.bo))
	(*l).lstmNodeList[idx].state.s = tambah(darab(l.lstmNodeList[idx].state.g, l.lstmNodeList[idx].state.i), darab(sPrev, l.lstmNodeList[idx].state.f))
	(*l).lstmNodeList[idx].state.h = darab(l.lstmNodeList[idx].state.s, l.lstmNodeList[idx].state.o)

	(*l).lstmNodeList[idx].xc = xc
}

func (l *LstmNetwork) topDiffIs(topDiffH, topDiffS [][]float64, idx int) {
	// notice that top_diff_s is carried along the constant error carousel
	ds := tambah(darab(l.lstmNodeList[idx].state.o, topDiffH), topDiffS)
	do := darab(l.lstmNodeList[idx].state.s, topDiffH)
	di := darab(l.lstmNodeList[idx].state.g, ds)
	dg := darab(l.lstmNodeList[idx].state.i, ds)
	df := darab(l.lstmNodeList[idx].sPrev, ds)

	// diffs w.r.t. vector inside sigma / tanh function
	diInput := darab(sigmoidDeriv(l.lstmNodeList[idx].state.i), di)
	dfInput := darab(sigmoidDeriv(l.lstmNodeList[idx].state.f), df)
	doInput := darab(sigmoidDeriv(l.lstmNodeList[idx].state.o), do)
	dgInput := darab(tanHDeriv(l.lstmNodeList[idx].state.g), dg)

	// diffs w.r.t. inputs
	// fmt.Println("diInput", len(diInput[0]))
	// if fp == 0 {
	// 	// fmt.Println("(*l).lstmParam.wgDiff B4", (*l).lstmParam.wgDiff)
	// 	// fmt.Println("diInput", dgInput)
	// 	// fmt.Println("l.xc", l.xc)
	// 	// fmt.Println("OUTER", outer(dgInput, l.xc))
	// 	printShape("OUTER SHAPE", outer(dgInput, l.lstmNodeList[idx].xc))
	// }
	// fmt.Println("(*l).lstmParam.wgDiff B4", (*l).lstmParam.wgDiff[0][0])
	(*l).lstmParam.wiDiff = tambah(l.lstmParam.wiDiff, outer(diInput, l.lstmNodeList[idx].xc))
	(*l).lstmParam.wfDiff = tambah(l.lstmParam.wfDiff, outer(dfInput, l.lstmNodeList[idx].xc))
	(*l).lstmParam.woDiff = tambah(l.lstmParam.woDiff, outer(doInput, l.lstmNodeList[idx].xc))
	(*l).lstmParam.wgDiff = tambah(l.lstmParam.wgDiff, outer(dgInput, l.lstmNodeList[idx].xc))
	(*l).lstmParam.biDiff = tambah(l.lstmParam.biDiff, diInput)
	(*l).lstmParam.bfDiff = tambah(l.lstmParam.bfDiff, dfInput)
	(*l).lstmParam.boDiff = tambah(l.lstmParam.boDiff, doInput)
	(*l).lstmParam.bgDiff = tambah(l.lstmParam.bgDiff, dgInput)
	// if fp == 0 {
	// 	// fmt.Println("(*l).lstmParam.wgDiff AFTER", (*l).lstmParam.wgDiff)
	// 	fp++
	// }
	// if fp == 0 {
	// 	fp++
	// }

	// compute bottom diff
	// fmt.Println("diInput", len(diInput), len(diInput[0]))
	// fmt.Println("l.lstmParam.wi", len(l.lstmParam.wi), len(l.lstmParam.wi[0]))
	// fmt.Println("l.lstmParam.wi", diInput)
	// fmt.Println("l.lstmParam.wi", len(transpose(l.lstmParam.wi)))
	dxc := darabN(l.lstmNodeList[idx].xc, 0)
	dxc = tambah(dxc, dot(diInput, l.lstmParam.wi))
	dxc = tambah(dxc, dot(dfInput, l.lstmParam.wf))
	dxc = tambah(dxc, dot(doInput, l.lstmParam.wo))
	dxc = tambah(dxc, dot(dgInput, l.lstmParam.wg))

	// save bottom diffs
	// printShape("dxc", dxc)
	// fmt.Println("xDim", l.lstmParam.xDim)
	(*l).lstmNodeList[idx].state.btmDiffS = darab(ds, l.lstmNodeList[idx].state.f)
	(*l).lstmNodeList[idx].state.btmDiffH = cutArr(dxc, l.lstmParam.xDim)
}

func (l *LstmNetwork) init(lp LstmParam) {
	(*l).lstmParam = lp
	// (*l).lstmNodeList = zeros(1, 1)
	// input sequence
	// d := [][][]float64{{{0}}, {{0}}, {{0}}, {{0}}}
	// (*l).xList = [][][]float64{{{0}}, {{0}}, {{0}}, {{0}}}
}

func (l *LstmNetwork) yListIs(yList [][]float64, lossLayer ToyLossLayer) float64 {
	/*
			Updates diffs by setting target sequence
		    with corresponding loss layer.
		    Will *NOT* update parameters.  To update parameters,
		    call self.lstm_param.apply_diff()
	*/

	// assert len(y_list) == len(self.x_list)
	if len(yList[0]) != len(l.xList) {
		// log.Fatalln("len(yList) != len(l.xList)")
		fmt.Println("yList", len(yList), ":", len(yList[0]))
		fmt.Println("xList", len(l.xList))
	}

	idx := len(l.xList) - 1

	// fmt.Println("idx:", idx)
	// fmt.Println("l.lstmNodeList[idx].state.h:", l.lstmNodeList[idx].state.h[0][0])
	// fmt.Println("yList[0][idx]:", yList[0][idx])

	//first node only gets diffs from label ...
	loss := lossLayer.loss(l.lstmNodeList[idx].state.h, yList[0][idx])
	diffH := lossLayer.btmDiff(l.lstmNodeList[idx].state.h, yList[0][idx])

	// fmt.Println("diffH:", diffH[0][0])

	// here s is not affecting loss due to h(t+1), hence we set equal to zero
	diffS := zeros(1, l.lstmParam.memCellCt)
	// l.lstmNodeList[idx].topDiffIs(diffH, diffS, idx)
	l.topDiffIs(diffH, diffS, idx)

	idx = idx - 1

	// ... following nodes also get diffs from next nodes, hence we add diffs to diff_h
	// we also propagate error along constant error carousel using diff_s
	for idx >= 0 {
		loss = loss + lossLayer.loss(l.lstmNodeList[idx].state.h, yList[0][idx])
		diffH = lossLayer.btmDiff(l.lstmNodeList[idx].state.h, yList[0][idx])
		diffH = tambah(diffH, l.lstmNodeList[idx+1].state.btmDiffH)
		diffS = l.lstmNodeList[idx+1].state.btmDiffS
		l.topDiffIs(diffH, diffS, idx)
		idx = idx - 1
	}

	return loss
}

func (l *LstmNetwork) xListClear() {
	// d := [][][]float64{{{0}}}
	(*l).xList = nil
}

func (l *LstmNetwork) xListAdd(x [][]float64) {
	// fmt.Println("len(l.xList)", l.xList)
	// fmt.Println("x", x)
	(*l).xList = append(l.xList, x)
	// fmt.Println("len(l.xList)", l.xList)
	// fmt.Println("len(l.xList)", len(l.xList))
	// fmt.Println("len(l.lstmNodeList)", len(l.lstmNodeList))
	if len(l.xList) > len(l.lstmNodeList) {
		// need to add new lstm node, create new state mem
		ls := LstmState{}
		ls.init(l.lstmParam.memCellCt)
		ln := LstmNode{}
		ln.init(l.lstmParam, ls)
		(*l).lstmNodeList = append(l.lstmNodeList, ln)
	}

	// get index of most recent x input
	idx := len(l.xList) - 1
	if idx == 0 {
		// no recurrent inputs yet
		var sPrev [][]float64
		var hPrev [][]float64
		l.btmDataIs(x, sPrev, hPrev, idx)
	} else {
		// fmt.Println("idx:", idx)
		// fmt.Println("l.lstmNodeList[idx-1].state.s", len(l.lstmNodeList))
		sPrev := l.lstmNodeList[idx-1].state.s
		hPrev := l.lstmNodeList[idx-1].state.h
		(*l).btmDataIs(x, sPrev, hPrev, idx)
	}

	// fmt.Println("exit this func")
}
