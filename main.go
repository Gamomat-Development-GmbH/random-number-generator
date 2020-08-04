package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"
)

const RESEEDING_INTERVAL_MILLI_SECONDS_MIN = 1000
const RESEEDING_INTERVAL_MILLI_SECONDS_MAX = 10000

const SLEEP_INTERVAL_MILLI_SECONDS_MIN = 1000
const SLEEP_INTERVAL_MILLI_SECONDS_MAX = 10000

const NUMBERS_TO_CYCLE_MIN = 10
const NUMBERS_TO_CYCLE_MAX = 100

func main() {
	fmt.Println("Server started...", fmt.Sprintf("GOOS %s GOARCH %s", runtime.GOOS, runtime.GOARCH))

	http.HandleFunc("/", RandomNumberHandler)

	server := &http.Server{Addr: ":8080"}

	go func() {
		time.Sleep(getNextBackgroundReseedingInterval())
		reseed()
	}()

	go func() {
		time.Sleep(getNextBackgroundCyclingInterval())
		cycle()
	}()

	log.Fatal(server.ListenAndServe())
}

func getNextBackgroundReseedingInterval() time.Duration {
	return time.Millisecond * time.Duration(getRandomNumbers(RESEEDING_INTERVAL_MILLI_SECONDS_MIN, RESEEDING_INTERVAL_MILLI_SECONDS_MAX, 1)[0])
}

func getNextBackgroundCyclingInterval() time.Duration {
	return time.Millisecond * time.Duration(getRandomNumbers(SLEEP_INTERVAL_MILLI_SECONDS_MIN, SLEEP_INTERVAL_MILLI_SECONDS_MAX, 1)[0])
}

func cycle() {
	getRandomNumbers(0, 10, getRandomNumbers(NUMBERS_TO_CYCLE_MIN, NUMBERS_TO_CYCLE_MAX, 1)[0])
}
