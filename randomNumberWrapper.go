package main

import (
	"github.com/seehuhn/mt19937"
	"math/rand"
	"sync"
	"time"
)

const NUMBERS_TO_CYCLE_MIN = 10
const NUMBERS_TO_CYCLE_MAX = 100

const CYCLE_STEP_MIN = 100
const CYCLE_STEP_MAX = 600
const CYCLE_STEP_RANGE = CYCLE_STEP_MAX - CYCLE_STEP_MIN

var mux sync.Mutex

var (
	rng                 = newRand()
	currentRandomNumber = 0
)

func reseed() {
	mux.Lock()
	rng.Seed(time.Now().UnixNano())
	mux.Unlock()
}

func newRand() *rand.Rand {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	return rng
}

func getRandomNumbers(min, maxExclusive, count int) []int {
	mux.Lock()
	numberRange := maxExclusive - min

	cycleStepSize := rng.Intn(CYCLE_STEP_RANGE) + CYCLE_STEP_MIN

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		currentRandomNumber++
		if currentRandomNumber%cycleStepSize == 0 {
			cycleWithoutLocking()
		}

		numbers[i] = rng.Intn(numberRange) + min
	}
	mux.Unlock()
	return numbers
}

func cycleWithoutLocking() {
	numbersToCycle := rng.Intn(NUMBERS_TO_CYCLE_MAX) + NUMBERS_TO_CYCLE_MIN
	for i := 0; i < numbersToCycle; i++ {
		rng.Int()
	}
}

func cycle() {
	mux.Lock()
	cycleWithoutLocking()
	mux.Unlock()
}
