package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const SLEEP_INTERVAL_MILLI_SECONDS_MIN = 1000
const SLEEP_INTERVAL_MILLI_SECONDS_MAX = 10000

var (
	sleepIntervalRange = big.NewInt(SLEEP_INTERVAL_MILLI_SECONDS_MAX - SLEEP_INTERVAL_MILLI_SECONDS_MIN)
)

func main() {
	fmt.Println("Server started...", fmt.Sprintf("GOOS %s GOARCH %s", runtime.GOOS, runtime.GOARCH))

	http.HandleFunc("/", RandomNumberHandler)

	server := &http.Server{Addr: ":8080"}

	go func() {
		for {
			time.Sleep(getNextBackgroundInterval())
			reseed()
		}
	}()

	go func() {
		for {
			time.Sleep(getNextBackgroundInterval())
			cycle()
		}
	}()

	go listenForGracefulShutdown(server, time.Second)
	log.Fatal(server.ListenAndServe())
}

func getNextBackgroundInterval() time.Duration {
	millisecondsToWait := getRandomNumber(systemSecureReader, sleepIntervalRange) + SLEEP_INTERVAL_MILLI_SECONDS_MIN
	return time.Millisecond * time.Duration(millisecondsToWait)
}

func listenForGracefulShutdown(server *http.Server, gracefulPeriod time.Duration) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM)
	<-sigint

	//close rng to save the seed file
	close()

	fmt.Println("Waiting for graceful shutdown")
	time.Sleep(gracefulPeriod)
	fmt.Println("Shutdown now")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}
}
