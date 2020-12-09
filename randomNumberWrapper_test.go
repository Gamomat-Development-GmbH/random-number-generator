package main

import (
	"fmt"
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
