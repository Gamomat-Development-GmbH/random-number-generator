package main

import (
	"github.com/seehuhn/mt19937"
	"math/rand"
	"sync"
	"time"
)

var mux sync.Mutex

var (
	rng = newRand()
)

func newRand() *rand.Rand {
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	return rng
}

func getRandomNumbers(min, maxExclusive, count int) []int {
	mux.Lock()
	defer mux.Unlock()

	numberRange := maxExclusive - min

	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rng.Intn(numberRange) + min
	}
	return numbers
}
