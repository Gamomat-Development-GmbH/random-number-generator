package main

import (
	"crypto/rand"
	"github.com/seehuhn/mt19937"
	"io"
	"math"
	"math/big"
	"sync"
)

const NUMBERS_TO_CYCLE_MIN = 10
const NUMBERS_TO_CYCLE_MAX = 100

const CYCLE_STEP_MIN = 100
const CYCLE_STEP_MAX = 600

var mux sync.Mutex

var (
	mersenneTwisterReader = mt19937.New()
	systemSecureReader    = rand.Reader

	maxInt64            = big.NewInt(math.MaxInt64)
	cycleStepRange      = big.NewInt(CYCLE_STEP_MAX - CYCLE_STEP_MIN)
	numbersToCycleRange = big.NewInt(NUMBERS_TO_CYCLE_MAX - NUMBERS_TO_CYCLE_MIN)

	currentRandomNumber int64 = 0
	nextCycleStepSize         = getRandomNumber(systemSecureReader, cycleStepRange) + CYCLE_STEP_MIN
)

func reseed() {
	mux.Lock()
	mersenneTwisterReader.Seed(getRandomNumber(systemSecureReader, maxInt64))
	mux.Unlock()
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

		numbers[i] = getRandomNumber(mersenneTwisterReader, numberRange)
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
		mersenneTwisterReader.Int63()
	}
}

func cycle() {
	mux.Lock()
	cycleWithoutLocking()
	mux.Unlock()
}
