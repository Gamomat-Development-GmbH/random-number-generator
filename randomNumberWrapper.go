package main

import (
	"github.com/seehuhn/mt19937"
	"math/rand"
	"sync"
	"time"
)

const NUMBERS_TO_CYCLE_MIN = 10
const NUMBERS_TO_CYCLE_MAX = 100
const NUMBERS_TO_CYCLE_RANGE = NUMBERS_TO_CYCLE_MAX - NUMBERS_TO_CYCLE_MIN

const CYCLE_STEP_MIN = 100
const CYCLE_STEP_MAX = 600
const CYCLE_STEP_RANGE = CYCLE_STEP_MAX - CYCLE_STEP_MIN

var mux sync.Mutex

var (
	mersenneTwister     = newMersenneTwister()
	currentRandomNumber = 0
)

func reseed() {
	mux.Lock()
	mersenneTwister.Seed(rand.Int63())
	mux.Unlock()
}

func newMersenneTwister() *rand.Rand {
	rand.Seed(time.Now().UnixNano())

	mersenneTwister := rand.New(mt19937.New())
	mersenneTwister.Seed(rand.Int63())
	return mersenneTwister
}

func getRandomNumbers(min, maxExclusive, count int) []int {
	mux.Lock()
	numberRange := maxExclusive - min

	cycleStepSize := rand.Intn(CYCLE_STEP_RANGE) + CYCLE_STEP_MIN

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		currentRandomNumber++
		if currentRandomNumber%cycleStepSize == 0 {
			cycleWithoutLocking()
		}

		numbers[i] = mersenneTwister.Intn(numberRange) + min
	}
	mux.Unlock()
	return numbers
}

func cycleWithoutLocking() {
	numbersToCycle := rand.Intn(NUMBERS_TO_CYCLE_RANGE) + NUMBERS_TO_CYCLE_MIN
	for i := 0; i < numbersToCycle; i++ {
		mersenneTwister.Int()
	}
}

func cycle() {
	mux.Lock()
	cycleWithoutLocking()
	mux.Unlock()
}
