package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"
	"time"
)

func Test_basicTest(t *testing.T) {
	for i := 1; i < 5; i++ {
		fmt.Println(getRandomNumbers(10, 100, 100))
	}
}

func Test_testTime(t *testing.T) {
	before := time.Now()

	getRandomNumbers(10, 100, 1000000)

	after := time.Now()
	duration := after.Sub(before)
	fmt.Println(duration)
}

func Test_test513(t *testing.T) {
	for i := 1; i < 100; i++ {
		value, _ := rand.Int(systemSecureReader, big.NewInt(513))
		fmt.Println(value)
	}
}
