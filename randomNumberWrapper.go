package main

import (
	"crypto/rand"
	"io"
	"math"
	"math/big"
	"rng/fortuna"
	"sync"
)

const NUMBERS_TO_CYCLE_MIN = 1
const NUMBERS_TO_CYCLE_MAX = 10

const CYCLE_STEP_MIN = 100
const CYCLE_STEP_MAX = 600

var mux sync.Mutex

var (
	rngReader          = newFortuna()
	systemSecureReader = rand.Reader

	maxInt64            = big.NewInt(math.MaxInt64)
	cycleStepRange      = big.NewInt(CYCLE_STEP_MAX - CYCLE_STEP_MIN)
	numbersToCycleRange = big.NewInt(NUMBERS_TO_CYCLE_MAX - NUMBERS_TO_CYCLE_MIN)

	currentRandomNumber int64 = 0
	nextCycleStepSize         = getRandomNumber(systemSecureReader, cycleStepRange) + CYCLE_STEP_MIN
)

func newFortuna() *fortuna.Accumulator {
	rngReader, err := fortuna.NewRNG("seed.dat")
	if err != nil {
		panic("cannot initialise the RNG: " + err.Error())
	}

	return rngReader
}

func close() {
	rngReader.Close()
}

func getRandomNumbers(min, maxExclusive, count int64) []int64 {
	mux.Lock()

	numberRange := big.NewInt(maxExclusive - min)

	numbers := make([]int64, count)
	var i int64
	for i = 0; i < count; i++ {
		currentRandomNumber++
		if currentRandomNumber%nextCycleStepSize == 0 {
			cycleWithoutLocking()
			nextCycleStepSize = getRandomNumber(systemSecureReader, cycleStepRange) + CYCLE_STEP_MIN
		}

		numbers[i] = min + getRandomNumber(rngReader, numberRange)
	}
	mux.Unlock()
	return numbers
}

func getRandomNumber(reader io.Reader, maxExclusive *big.Int) int64 {
	result, err := rand.Int(reader, maxExclusive)
	if err != nil {
		panic(err)
	}
	return result.Int64()
}

func cycleWithoutLocking() {
	numbersToCycle := getRandomNumber(systemSecureReader, numbersToCycleRange) + NUMBERS_TO_CYCLE_MIN
	for i := 0; i < int(numbersToCycle); i++ {
		rngReader.Int63()
	}
}

func cycle() {
	mux.Lock()
	cycleWithoutLocking()
	mux.Unlock()
}
