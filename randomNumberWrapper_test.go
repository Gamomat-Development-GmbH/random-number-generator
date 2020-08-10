package main

import (
	"fmt"
	"testing"
)

func Test_basicTest(t *testing.T) {
	for i := 1; i < 5; i++ {
		fmt.Println(getRandomNumbers(10, 100, 100))
	}
}
