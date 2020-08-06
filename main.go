package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"time"
)

const RESEEDING_INTERVAL_MILLI_SECONDS_MIN = 1000
const RESEEDING_INTERVAL_MILLI_SECONDS_MAX = 10000
const RESEEDING_INTERVAL_MILLI_SECONDS_RANGE = RESEEDING_INTERVAL_MILLI_SECONDS_MAX - RESEEDING_INTERVAL_MILLI_SECONDS_MIN

const SLEEP_INTERVAL_MILLI_SECONDS_MIN = 1000
const SLEEP_INTERVAL_MILLI_SECONDS_MAX = 10000
const SLEEP_INTERVAL_MILLI_SECONDS_RANGE = SLEEP_INTERVAL_MILLI_SECONDS_MAX - SLEEP_INTERVAL_MILLI_SECONDS_MIN

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
	millisecondsToWait := rand.Intn(RESEEDING_INTERVAL_MILLI_SECONDS_RANGE) + RESEEDING_INTERVAL_MILLI_SECONDS_MIN
	return time.Millisecond * time.Duration(millisecondsToWait)
}

func getNextBackgroundCyclingInterval() time.Duration {
	millisecondsToWait := rand.Intn(SLEEP_INTERVAL_MILLI_SECONDS_RANGE) + SLEEP_INTERVAL_MILLI_SECONDS_MIN
	return time.Millisecond * time.Duration(millisecondsToWait)
}
