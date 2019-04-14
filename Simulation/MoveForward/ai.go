package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"gonum.org/v1/gonum/stat"
)

/*
Example reward calculation
def _step(self, action):
	xposbefore = self.model.data.qpos[0, 0]
	self.do_simulation(action, self.frame_skip)
	xposafter = self.model.data.qpos[0, 0]
	ob = self._get_obs()
	reward_ctrl = - 0.1 * np.square(action).sum()
	reward_run = (xposafter - xposbefore)/self.dt
	reward = reward_ctrl + reward_run
	done = False
	return ob, reward, done, dict(reward_run=reward_run, reward_ctrl=reward_ctrl)

Another reference:
https://towardsdatascience.com/openai-gym-from-scratch-619e39af121f
set proportinality to determine a good/bad reward

*/

var memo MemoryInit
var thetaToWrite [][]float64
var nToWrite [][]float64
var meanToWrite [][]float64
var meanDiffToWrite [][]float64
var varianceToWrite [][]float64

var gInput = 7
var gOutput = 2
var gTime float64

// MemoryInit ...
type MemoryInit struct {
	Mean     [][]float64 `json:"mean"`
	MeanDiff [][]float64 `json:"meanDiff"`
	N        [][]float64 `json:"n"`
	Theta    [][]float64 `json:"theta"`
	Variance [][]float64 `json:"variance"`
}

// Hp is Hyper Parameter
type Hp struct {
	nbSteps, episodeLength, nbDirections, nbBestDirections int
	learningRate, noise                                    float64
}

// Normalizer is Normalizer
type Normalizer struct {
	n, mean, meanDiff, variance [][]float64
}

// Policy is the AI
type Policy struct {
	theta [][]float64
}

type rollouts struct {
	rPos, rNeg float64
	d          [][]float64
}

type env struct {
	nbInputs, nbOutputs int
}

func (hp *Hp) init() {
	(*hp).nbSteps = 100
	// this part (episode length) is very crucial and will determine the model is working or not (must tally with "robot goal/done" @below function) (or maybe not)
	(*hp).episodeLength = 10000000
	(*hp).nbDirections = 16
	(*hp).nbBestDirections = 16
	(*hp).learningRate = 0.02
	(*hp).noise = 0.03
}

func (nm *Normalizer) init(nbInputs int) {
	(*nm).n = zeros(1, nbInputs)
	(*nm).mean = zeros(1, nbInputs)
	(*nm).meanDiff = zeros(1, nbInputs)
	(*nm).variance = zeros(1, nbInputs)
}

func (nm *Normalizer) init2(nbInputs int) {
	(*nm).n = memo.N
	(*nm).mean = memo.Mean
	(*nm).meanDiff = memo.MeanDiff
	(*nm).variance = memo.Variance
}

func (nm *Normalizer) observe(x [][]float64) {
	(*nm).n = tambahN((*nm).n, 1)
	lastMean := (*nm).mean
	mean1 := tolak(x, (*nm).mean)
	mean1 = bahagi(mean1, (*nm).n)
	(*nm).mean = tambah((*nm).mean, mean1)
	meanDiff1 := tolak(x, lastMean)
	meanDiff1 = darab(meanDiff1, tolak(x, (*nm).mean))
	(*nm).meanDiff = tambah((*nm).meanDiff, meanDiff1)
	variance1 := bahagi((*nm).meanDiff, (*nm).n)
	variance1 = clipMin(variance1, 1e-2)
	(*nm).variance = variance1

	nToWrite = (*nm).n
	meanToWrite = (*nm).mean
	meanDiffToWrite = (*nm).meanDiff
	varianceToWrite = (*nm).variance
}

func (nm *Normalizer) normalize(inputs [][]float64) [][]float64 {
	obsMean := (*nm).mean
	obsStd := sqrt((*nm).variance)
	r1 := tolak(inputs, obsMean)
	r1 = bahagi(r1, obsStd)
	// fmt.Println("mean", (*nm).mean)
	// fmt.Println("var", (*nm).variance)
	// fmt.Println("input", inputs)
	// fmt.Println("obs_std", obsStd)
	return r1
}

func (p *Policy) init(inputSize, outputSize int) {
	(*p).theta = zeros(outputSize, inputSize)
}

func (p *Policy) init2(inputSize, outputSize int) {
	(*p).theta = memo.Theta
}

func (p *Policy) evaluate(input, delta [][]float64, direction string, hp Hp) [][]float64 {
	switch direction {
	case "none":
		// fmt.Println("NONE")
		// fmt.Println("input", input)
		// fmt.Println("theta", (*p).theta)
		return dot(input, transpose((*p).theta))
	case "positive":
		r := darabN(delta, hp.noise)
		r = tambah((*p).theta, r)
		// fmt.Println("POS")
		// fmt.Println("input", input)
		// fmt.Println("theta", (*p).theta)
		// fmt.Println("transpose(r)", transpose(r))
		return dot(input, transpose(r))
	default:
		r := darabN(delta, hp.noise)
		r = tolak((*p).theta, r)
		// fmt.Println("NEG")
		// fmt.Println("input", input)
		// fmt.Println("theta", (*p).theta)
		// fmt.Println("transpose(r)", transpose(r))
		return dot(input, transpose(r))
	}
}

func (p *Policy) sampleDeltas(inputSize, outputSize int, hp Hp) [][][]float64 {
	var r [][][]float64
	for i := 0; i < hp.nbDirections; i++ {
		r = append(r, randomizeValue(outputSize, inputSize))
	}
	return r
}

func (p *Policy) update(rollout []rollouts, sigmaR float64, hp Hp) {

	step := zeros(len(rollout[0].d), len(rollout[0].d[0]))

	for _, v := range rollout {
		s1 := v.rPos - v.rNeg
		s2 := darabN(v.d, s1)

		step = tambah(step, s2)
	}

	// fmt.Println("sigmaR", sigmaR)
	// os.Exit(3)

	ss1 := float64(hp.nbBestDirections) * sigmaR
	ss1 = float64(hp.learningRate) / ss1
	if math.IsInf(ss1, 0) {
		// this part will happen if the modal is not working (episode length & goal are not tally with each other)
		fmt.Println("err: ss1 = Inf")
		fmt.Println("sigmaR", sigmaR)
		os.Exit(3)
		// ss1 = 0
	}
	// fmt.Println("ss1", ss1)
	ss2 := darabN(step, ss1)

	// fmt.Println("theta", (*p).theta)
	// fmt.Println("ss2", ss2)
	(*p).theta = tambah((*p).theta, ss2)
	// fmt.Println("theta", (*p).theta)
	// time.Sleep(5 * time.Second)
	// os.Exit(3)

	thetaToWrite = (*p).theta
}

func readIn() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		// handle error.
		log.Fatalln(scanner.Err())
	}

	var s float64

	s, err := strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		// handle error.
		log.Fatalln(err)
	}

	return s
}

func (bot *Robot) gym(action [][]float64) ([][]float64, float64, bool) {
	// to compile the state into array later
	r := 1
	c := gInput
	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}

	leftM := action[0][0]
	rightM := action[0][1]

	bot.moveBot(leftM, rightM)

	f := bot.facing
	fI := 0.0
	switch f {
	case 'Q':
		fI = 1
	case 'W':
		fI = 0
	case 'E':
		fI = 1
	case 'D':
		fI = 2
	case 'C':
		fI = 3
	case 'X':
		fI = 4
	case 'Z':
		fI = 3
	case 'A':
		fI = 2
	}

	gTime++
	// fmt.Println("GTIME:", gTime)

	// y-axis = 0.09733, degErr = 0, x-axis = 0, motorL = 0, motorR = 0, facing = 0
	v[0][0] = bot.head.y
	v[0][1] = bot.errDeg()
	v[0][2] = bot.head.x
	v[0][3] = leftM
	v[0][4] = rightM
	v[0][5] = fI
	v[0][6] = gTime

	// fmt.Println(v[0][0], v[0][1], v[0][2], v[0][3], v[0][4], v[0][5])

	// Then decide the rewards
	reward := getReward(v[0][0], v[0][1], v[0][2], v[0][3], v[0][4], v[0][5], v[0][6])
	// it is hard to reach 900

	// Decide if it is done
	done := false

	// this part ("done" AKA. "robot goal") is very crucial and will determine the model is working or not
	if math.Abs(v[0][0]) > 5 || math.Abs(v[0][2]) > 5 {
		// fmt.Println("done1:", v[0][0], "done2:", v[0][2])
		done = true
	}
	// fmt.Println("reward:", reward)

	return v, reward, done
}

func getReward(v0, v1, v2, v3, v4, v5, v6 float64) float64 {

	var reward, e0, e1, e2, e3, e4, e5, e6 float64

	//eq. v0 || y-axis
	e0 = -v0 + 100
	e1 = v1
	e2 = math.Abs(v2)
	e3 = -v3 + 1
	e4 = -v4 + 1
	e5 = v5
	e6 = v6

	reward = 25 - (6 * e0) - (5 * e1) - (4 * e2) - (3 * e3) - (3 * e4) - (2 * e5) - (2 * e6)

	return reward
}

func (bot *Robot) envReset() [][]float64 {
	// y-axis = 0.09733, degErr = 0, x-axis = 0, motorL = 0, motorR = 0, facing = 0
	s := [][]float64{{0.09733, 0, 0, 0, 0, 0, 0}}
	gTime = 0

	return s
}

// RESET = 1x7
// STATE = 1x7 (input)
// ACTION = 1x2 (output)

func explore(hp Hp, normalizer Normalizer, policy Policy, direction string, delta [][]float64) float64 {
	// I assume reset the robot leg to default
	// state := env.reset()
	bot := Robot{}
	bot.init()
	state := bot.envReset()

	done := false
	var numPlays, sumRewards, reward float64

	// Calculate the accumulate reward on the full episode
	for !done && numPlays < float64(hp.episodeLength) {
		normalizer.observe(state)
		state = normalizer.normalize(state)
		// fmt.Println("normalize state:", state)
		action := policy.evaluate(state, delta, direction, hp)
		// fmt.Println("action:", action)

		state, reward, done = bot.gym(action)
		// fmt.Println("state:", state)
		// fmt.Println("reward:", reward)

		if reward < -1 {
			reward = -1
		} else if reward > 1 {
			reward = 1
		}
		sumRewards += reward
		numPlays++
		// fmt.Println("numPlays", numPlays)
	}
	// bot.printBotCoor("coor")

	return sumRewards
}

func train(hp Hp, p Policy, normalizer Normalizer, inputSize, outputSize int) {
	for step := 0; step < hp.nbSteps; step++ {
		// Initializing the pertubation deltas and the positive/negative rewards
		deltas := p.sampleDeltas(inputSize, outputSize, hp)
		// fmt.Println("deltas:", deltas)
		positiveRewards := zeros(1, hp.nbDirections)
		negativeRewards := zeros(1, hp.nbDirections)

		// Getting the positive rewards in the positive directions
		for k := 0; k < hp.nbDirections; k++ {
			positiveRewards[0][k] = explore(hp, normalizer, p, "positive", deltas[k])
		}

		// Getting the negative rewards in the negative/positive directions
		for k := 0; k < hp.nbDirections; k++ {
			negativeRewards[0][k] = explore(hp, normalizer, p, "negative", deltas[k])
		}

		// Gathering all the positive/negative rewards to compute the standard deviation of these rewards
		// Concat both into 1 array
		allRewards := concatArr(positiveRewards, negativeRewards)
		// go lib for std might be a bit different than python
		sigmaR := std(allRewards)
		// fmt.Println("sigmaR:", sigmaR, "allRewards", allRewards)
		// time.Sleep(10 * time.Second)

		// Sorting the rollouts by the max(r_pos, r_neg) and selecting the best directions
		order := getOrder(positiveRewards, negativeRewards)
		rollout := make([]rollouts, len(positiveRewards[0]))
		for i, v := range order {
			rollout[i].rPos = positiveRewards[0][v]
			rollout[i].rNeg = negativeRewards[0][v]
			rollout[i].d = deltas[v]
		}

		// Updating the policy
		p.update(rollout, sigmaR, hp)

		// Printing the final reward of the policy after the update
		rewardEvaluation := explore(hp, normalizer, p, "none", deltas[0])
		fmt.Println("Step:", step, "Rewards:", rewardEvaluation)
	}
}

func startAI() {
	fmt.Println("Start AI")

	hp := Hp{}
	hp.init()

	// Number of input and output depend on our own gym
	nbInputs := gInput
	nbOutputs := gOutput

	policy := Policy{}
	// policy.init(nbInputs, nbOutputs)

	normalizer := Normalizer{}
	// normalizer.init(nbInputs)

	db := "./memory/memory.json"
	exs, err := exists(db)
	if err != nil {
		log.Fatalln(err)
	}
	if !exs {
		fmt.Println("Not exist")
		// time.Sleep(2 * time.Second)
		policy.init(nbInputs, nbOutputs)
		normalizer.init(nbInputs)
	} else {
		fmt.Println("exist")
		// time.Sleep(2 * time.Second)
		content, err := ioutil.ReadFile(db)
		if err != nil {
			log.Fatalln(err)
		}
		// fmt.Println("content:", content)
		if err = json.Unmarshal(content, &memo); err != nil {
			log.Fatalln(err)
		}

		policy.init2(nbInputs, nbOutputs)
		normalizer.init2(nbInputs)
	}

	train(hp, policy, normalizer, nbInputs, nbOutputs)

	memoTheta()

	fmt.Println("End AI")
}

func exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			return false, nil
		}
		// exist but got other error
		return true, err
	}
	return true, nil
}

func memoTheta() {
	db := "./memory/memory.json"
	s := "{\"theta\":["

	// s := fmt.Sprintf("%f", (*p).theta[0][0])
	fmt.Println("theta:", thetaToWrite)
	fmt.Println("n:", nToWrite)
	fmt.Println("mean:", meanToWrite)
	fmt.Println("meanDiff:", meanDiffToWrite)
	fmt.Println("variance:", varianceToWrite)

	for i := 0; i < len(thetaToWrite); i++ {
		s = s + "["
		for j := 0; j < len(thetaToWrite[0]); j++ {
			temp := strconv.FormatFloat(thetaToWrite[i][j], 'f', -1, 64)
			s = s + temp
			if j != len(thetaToWrite[0])-1 {
				s = s + ","
			}
		}
		s = s + "]"
		if i != len(thetaToWrite)-1 {
			s = s + ","
		}
	}

	s = s + "],\"n\":["

	for i := 0; i < len(nToWrite); i++ {
		s = s + "["
		for j := 0; j < len(nToWrite[0]); j++ {
			temp := strconv.FormatFloat(nToWrite[i][j], 'f', -1, 64)
			s = s + temp
			if j != len(nToWrite[0])-1 {
				s = s + ","
			}
		}
		s = s + "]"
		if i != len(nToWrite)-1 {
			s = s + ","
		}
	}

	s = s + "],\"mean\":["

	for i := 0; i < len(meanToWrite); i++ {
		s = s + "["
		for j := 0; j < len(meanToWrite[0]); j++ {
			temp := strconv.FormatFloat(meanToWrite[i][j], 'f', -1, 64)
			s = s + temp
			if j != len(meanToWrite[0])-1 {
				s = s + ","
			}
		}
		s = s + "]"
		if i != len(meanToWrite)-1 {
			s = s + ","
		}
	}

	s = s + "],\"meanDiff\":["

	for i := 0; i < len(meanDiffToWrite); i++ {
		s = s + "["
		for j := 0; j < len(meanDiffToWrite[0]); j++ {
			temp := strconv.FormatFloat(meanDiffToWrite[i][j], 'f', -1, 64)
			s = s + temp
			if j != len(meanDiffToWrite[0])-1 {
				s = s + ","
			}
		}
		s = s + "]"
		if i != len(meanDiffToWrite)-1 {
			s = s + ","
		}
	}

	s = s + "],\"variance\":["

	for i := 0; i < len(varianceToWrite); i++ {
		s = s + "["
		for j := 0; j < len(varianceToWrite[0]); j++ {
			temp := strconv.FormatFloat(varianceToWrite[i][j], 'f', -1, 64)
			s = s + temp
			if j != len(varianceToWrite[0])-1 {
				s = s + ","
			}
		}
		s = s + "]"
		if i != len(varianceToWrite)-1 {
			s = s + ","
		}
	}

	s = s + "]}"

	var result map[string]interface{}
	var err error
	if err = json.Unmarshal([]byte(s), &result); err != nil {
		log.Fatalln(err)
	}
	var b []byte
	b, err = json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
	}
	if err = ioutil.WriteFile(db, b, 0644); err != nil {
		log.Fatalln(err)
	}
}

func randomizeValue(r int, c int) [][]float64 {
	// we are seeding the rand variable with present time
	// or use crypto/rand for more secure way
	// https://gobyexample.com/random-numbers
	// so that we would get different output each time
	rand.Seed(time.Now().UnixNano())
	// OR WE CAN JUST CONSTANT IT FOR NOW!
	// rand.Seed(0)

	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v[i][j] = 2*rand.Float64() - 1
		}
	}

	return v
}

func zeros(r int, c int) [][]float64 {
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

func std(m1 [][]float64) float64 {
	// https://www.gonum.org/post/intro-to-stats-with-gonum/
	// mean := stat.Mean(m1[0], nil)
	variance := stat.Variance(m1[0], nil)
	// fmt.Println("m1[0]", m1[0])
	if math.IsNaN(math.Sqrt(variance)) {
		fmt.Println("std() NaN encounter")
	}
	stddev := math.Sqrt(math.Abs(variance))
	// fmt.Println("stddev:", stddev)
	// os.Exit(3)

	return stddev
}

func concatArr(m1 [][]float64, m2 [][]float64) [][]float64 {
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

func getOrder(m1 [][]float64, m2 [][]float64) []int {

	m := make(map[float64]int)

	s := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		s[i] = make([]float64, len(m1[0]))
	}

	k := 0
	for i, v := range m1 {
		for i2, v2 := range v {
			if v2 > m2[i][i2] {
				s[i][i2] = v2
				m[v2] = k
				k++
			} else {
				s[i][i2] = m2[i][i2]
				m[m2[i][i2]] = k
				k++
			}
		}
	}
	sort.Float64s(s[0])

	var r []int

	for _, val := range s[0] {
		r = append(r, m[val])
		// fmt.Println(m[val], val)
	}
	// fmt.Println(r)
	return r
}

func dot(m1 [][]float64, m2 [][]float64) [][]float64 {

	// Ref 2d slice
	// https://gobyexample.com/slices
	output := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		output[i] = make([]float64, len(m2[0]))
	}

	// outR := len(m1)
	// outC := len(m2[0])

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

func darab(m1 [][]float64, m2 [][]float64) [][]float64 {

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

func bahagi(m1 [][]float64, m2 [][]float64) [][]float64 {

	division := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		division[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			division[i][i2] = v2 / m2[i][i2]
		}
	}

	return division
}

func tambah(m1 [][]float64, m2 [][]float64) [][]float64 {

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

func tambahN(m1 [][]float64, n float64) [][]float64 {

	sum := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		sum[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			sum[i][i2] = v2 + n
		}
	}

	return sum
}

func tolak(m1 [][]float64, m2 [][]float64) [][]float64 {

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

func clipMin(m1 [][]float64, x float64) [][]float64 {

	clip := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		clip[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			if v2 < x {
				clip[i][i2] = x
			} else {
				clip[i][i2] = v2
			}
		}
	}

	return clip
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

func sqrt(m1 [][]float64) [][]float64 {

	s := make([][]float64, len(m1))
	for i := 0; i < len(m1); i++ {
		s[i] = make([]float64, len(m1[0]))
	}

	for i, v := range m1 {
		for i2, v2 := range v {
			if math.IsNaN(math.Sqrt(v2)) {
				fmt.Println("sqrt() NaN encounter")
				fmt.Println("v2:", v2)
				os.Exit(3)
			}
			s[i][i2] = math.Sqrt(math.Abs(v2))
		}
	}

	return s
}
