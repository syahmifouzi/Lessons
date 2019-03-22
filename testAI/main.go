package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/jacobsa/go-serial/serial"
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
	(*hp).nbSteps = 1
	(*hp).episodeLength = 10
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
		// fmt.Println("none")
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

	ss1 := float64(hp.nbBestDirections) * sigmaR
	ss1 = float64(hp.learningRate) / ss1
	ss2 := darabN(step, ss1)

	(*p).theta = tambah((*p).theta, ss2)

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

func gym(action [][]float64, port io.ReadWriteCloser) ([][]float64, float64, bool) {
	// to compile the state into array later
	r := 1
	c := 5
	v := make([][]float64, r)
	for i := 0; i < r; i++ {
		v[i] = make([]float64, c)
	}
	// v[i][j] = 0

	// Slot buat ACTION begin (gerakkan motor)
	fmt.Println("Sending motor tru ble")

	// leftM := action[0][0] * 255
	// rightM := action[0][1] * 255

	leftM := action[0][0] * 205
	rightM := action[0][1] * 205

	if leftM < 0 {
		leftM = leftM - 50
	} else {
		leftM = leftM + 50
	}

	if rightM < 0 {
		rightM = rightM - 50
	} else {
		rightM = rightM + 50
	}

	if math.IsNaN(math.Round(leftM)) {
		fmt.Println("Left motor NaN triggered")
		leftM = 0
	}

	if math.IsNaN(math.Round(rightM)) {
		fmt.Println("Right motor NaN triggered")
		rightM = 0
	}

	leftString := strconv.FormatFloat(math.Round(leftM), 'f', -1, 64)
	rightString := strconv.FormatFloat(math.Round(rightM), 'f', -1, 64)

	// Write 2 bytes to the port
	b1 := []byte{'l'}
	_, err := port.Write(b1)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}
	b11 := []byte(leftString)
	_, err = port.Write(b11)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}

	b2 := []byte{'r'}
	_, err = port.Write(b2)
	if err != nil {
		log.Fatalln("port.WriteRight:", err)
	}
	b21 := []byte(rightString)
	_, err = port.Write(b21)
	if err != nil {
		log.Fatalln("port.WriteLeft:", err)
	}

	b3 := []byte{'q'}
	_, err = port.Write(b3)
	if err != nil {
		log.Fatalln("port.Write:", err)
	}

	fmt.Println("motor left:", leftString)
	fmt.Println("motor right:", rightString)

	fmt.Println("Done sending")
	fmt.Println("Waiting for sensor")
	// Slot buat ACTION end

	// Sekarang read sensor condition aka STATE

	splice := make([]byte, 128)
	nQuit := make([]byte, 1)
	total := 0

	for nQuit[0] != 'q' {
		buf := make([]byte, 128)
		n, err := port.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println("port.Read:", buf[:n])

		for i := 0; i < n; i++ {
			splice[i+total] = buf[i]
			nQuit[0] = buf[i]
		}
		// log.Println("is quit?", nQuit, nQuit[0])
		total = total + n
	}

	// log.Println("port.Read:", buf[:n], n, buf[0:1], buf[1:2])
	fmt.Println("done sensor.... now converting...")

	var command byte
	var ls0 []byte
	var ls1 []byte
	var ls2 []byte
	var ls3 []byte
	var ls4 []byte
	var ls5 []byte
	var ls6 []byte
	for i := 0; i < total; i++ {
		if command == 'a' && splice[i] > 47 && splice[i] < 58 {
			ls0 = append(ls0, splice[i])
		} else if command == 'b' && splice[i] > 47 && splice[i] < 58 {
			ls1 = append(ls1, splice[i])
		} else if command == 'c' && splice[i] > 47 && splice[i] < 58 {
			ls2 = append(ls2, splice[i])
		} else if command == 'd' && splice[i] > 47 && splice[i] < 58 {
			ls3 = append(ls3, splice[i])
		} else if command == 'e' && splice[i] > 47 && splice[i] < 58 {
			ls4 = append(ls4, splice[i])
		} else if command == 'f' && splice[i] > 47 && splice[i] < 58 {
			ls5 = append(ls5, splice[i])
		} else if command == 'g' && splice[i] > 47 && splice[i] < 58 {
			ls6 = append(ls6, splice[i])
		} else {
			command = splice[i]
		}
	}

	// s0 := string(buf[0:1])
	s1 := string(ls1)
	s2 := string(ls2)
	s3 := string(ls3)
	s4 := string(ls4)
	s5 := string(ls5)
	// s6 := string(buf[6:7])

	// f0, _ := strconv.ParseFloat(s0, 8)
	f1, _ := strconv.ParseFloat(s1, 8)
	f2, _ := strconv.ParseFloat(s2, 8)
	f3, _ := strconv.ParseFloat(s3, 8)
	f4, _ := strconv.ParseFloat(s4, 8)
	f5, _ := strconv.ParseFloat(s5, 8)
	// f6, _ := strconv.ParseFloat(s6, 8)

	// fmt.Print("s1-s5:", s1, s2, s3, s4, s5)
	v[0][0] = f1
	// fmt.Print("Insert S1: ")
	v[0][1] = f2
	// fmt.Print("Insert S2: ")
	v[0][2] = f3
	// fmt.Print("Insert S3: ")
	v[0][3] = f4
	// fmt.Print("Insert S4: ")
	v[0][4] = f5
	// fmt.Print("Insert S5: ")
	// v[0][5] = float64(s5[0])
	// // fmt.Print("Insert S6: ")
	// v[0][6] = float64(s6[0])

	log.Println(v[0][0], v[0][1], v[0][2], v[0][3], v[0][4])

	// Then decide the rewards
	reward := getReward(leftM, rightM, v[0][0], v[0][1], v[0][2], v[0][3], v[0][4])
	// it is hard to reach 900
	// var reward float64
	// // fmt.Print("Insert Reward: ")
	// if leftM < 0 && rightM < 0 {
	// 	reward = -10
	// 	fmt.Println("Case 1: reward =", reward)
	// } else if leftM == 0 && rightM == 0 {
	// 	reward = -10
	// 	fmt.Println("Case 2: reward =", reward)
	// } else if v[0][2] < 600 {
	// 	reward = -20
	// 	fmt.Println("Case 3: reward =", reward)
	// } else if v[0][2] > 600 && leftM > 50 && rightM > 50 {
	// 	reward = 20
	// 	fmt.Println("Case 4: reward =", reward)
	// } else {
	// 	reward = -1.1
	// 	fmt.Println("Case 5: reward =", reward)
	// }

	// Decide if it is done
	done := false
	fmt.Println("reward:", reward)

	return v, reward, done
}

func getReward(lm, rm, s1, s2, s3, s4, s5 float64) float64 {

	var reward, k1m, k2m, k3s, k4s, k5s, k6s, k7s float64

	k1m = lm / 255
	k1m = 1 - k1m

	k2m = rm / 255
	k2m = 1 - k2m

	k3s = 0.0
	if s1 > 600 && s1 < 800 {
		k3s = 0
	} else if s1 < 600 {
		k3s = s1 / 600
		k3s = 1 - k3s
	} else if s1 > 800 {
		k3s = s1 - 800
		k3s = k3s / 600
	}

	k4s = 0.0
	if s2 > 500 && s2 < 700 {
		k4s = 0
	} else if s2 < 500 {
		k4s = s2 / 500
		k4s = 1 - k4s
	} else if s2 > 700 {
		k4s = s2 - 700
		k4s = k4s / 700
	}

	k5s = 0.0
	if s3 > 900 {
		k5s = 0
	} else if s3 < 900 {
		k5s = s3 / 900
		k5s = 1 - k5s
	}

	k6s = 0.0
	if s4 > 500 && s4 < 700 {
		k6s = 0
	} else if s4 < 500 {
		k6s = s4 / 500
		k6s = 1 - k6s
	} else if s4 > 700 {
		k6s = s4 - 700
		k6s = k6s / 700
	}

	k7s = 0.0
	if s5 > 600 && s5 < 800 {
		k7s = 0
	} else if s5 < 600 {
		k7s = s5 / 600
		k7s = 1 - k7s
	} else if s5 > 800 {
		k7s = s5 - 800
		k7s = k7s / 600
	}

	reward = 5 - (2 * k1m) - (2 * k2m) - (3 * k3s) - (4 * k4s) - (5 * k5s) - (4 * k6s) - (3 * k7s)

	return reward
}

// RESET = 1x5
// STATE = 1x5 (input)
// ACTION = 1x2 (output)

func explore(hp Hp, normalizer Normalizer, policy Policy, direction string, delta [][]float64, port io.ReadWriteCloser) float64 {
	// I assume reset the robot leg to default
	// state := env.reset()
	state := zeros(1, 5)

	done := false
	var numPlays, sumRewards, reward float64

	// Calculate the accumulate reward on the full episode
	for !done && numPlays < float64(hp.episodeLength) {
		normalizer.observe(state)
		state = normalizer.normalize(state)
		action := policy.evaluate(state, delta, direction, hp)

		// fmt.Println("action:", action)

		state, reward, done = gym(action, port)

		if reward < -1 {
			reward = -1
		} else if reward > 1 {
			reward = 1
		}
		sumRewards += reward
		numPlays++
		// fmt.Println("numPlays", numPlays)
	}
	// fmt.Println("DONE")

	return sumRewards
}

func train(hp Hp, p Policy, normalizer Normalizer, inputSize, outputSize int, port io.ReadWriteCloser) {
	for step := 0; step < hp.nbSteps; step++ {
		// Initializing the pertubation deltas and the positive/negative rewards
		deltas := p.sampleDeltas(inputSize, outputSize, hp)
		positiveRewards := zeros(1, hp.nbDirections)
		negativeRewards := zeros(1, hp.nbDirections)

		// Getting the positive rewards in the positive directions
		for k := 0; k < hp.nbDirections; k++ {
			positiveRewards[0][k] = explore(hp, normalizer, p, "positive", deltas[k], port)
		}

		// Getting the negative rewards in the negative/positive directions
		for k := 0; k < hp.nbDirections; k++ {
			negativeRewards[0][k] = explore(hp, normalizer, p, "negative", deltas[k], port)
		}

		// Gathering all the positive/negative rewards to compute the standard deviation of these rewards
		// Concat both into 1 array
		allRewards := concatArr(positiveRewards, negativeRewards)
		// go lib for std might be a bit different than python
		sigmaR := std(allRewards)

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
		rewardEvaluation := explore(hp, normalizer, p, "none", deltas[0], port)
		fmt.Println("Step:", step, "Rewards:", rewardEvaluation)
	}
}

func main() {
	fmt.Println("Start")

	options := serial.OpenOptions{
		PortName:        "COM9", // in my PC, it is outgoing hc-05 'dev B'
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 2,
	}

	// Open the port
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalln("serial.Open:", err)
	}

	hp := Hp{}
	hp.init()

	// Number of input and output depend on our own gym
	nbInputs := 5
	nbOutputs := 2

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
		policy.init(nbInputs, nbOutputs)
		normalizer.init(nbInputs)
	} else {
		fmt.Println("exist")
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

	train(hp, policy, normalizer, nbInputs, nbOutputs, port)

	defer port.Close()

	memoTheta()

	fmt.Println("End")
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
	if math.IsNaN(math.Sqrt(variance)) {
		fmt.Println("std() NaN encounter")
	}
	stddev := math.Sqrt(math.Abs(variance))

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
			}
			s[i][i2] = math.Sqrt(math.Abs(v2))
		}
	}

	return s
}
