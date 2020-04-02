package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	fmt.Println("Server started...", fmt.Sprintf("GOOS %s GOARCH %s", runtime.GOOS, runtime.GOARCH))

	http.HandleFunc("/", RandomNumberHandler)

	server := &http.Server{Addr: ":8080"}

	log.Fatal(server.ListenAndServe())
}
