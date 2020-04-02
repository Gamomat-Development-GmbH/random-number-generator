package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type RandomNumberHttpError struct {
	Min          string `json:"min"`
	MaxExclusive string `json:"maxExclusive"`
	Count        string `json:"count"`
	ErrorMessage string `json:"errorMessage"`
}

type RandomNumberHttpResponse struct {
	Min           string `json:"min"`
	MaxExclusive  string `json:"maxExclusive"`
	Count         string `json:"count"`
	RandomNumbers []int  `json:"randomNumbers"`
}

func RandomNumberHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	min, minConvError := strconv.Atoi(query.Get("min"))
	maxExclusive, maxExclusiveConvError := strconv.Atoi(query.Get("maxExclusive"))
	count, countConvErr := strconv.Atoi(query.Get("count"))

	errorMessage := ""

	switch {
	case minConvError != nil:
		errorMessage = "min parameter is missing or not a number"
	case maxExclusiveConvError != nil:
		errorMessage = "maxExclusive parameter is missing or not a number"
	case countConvErr != nil:
		errorMessage = "count parameter is missing or not a number"
	case min >= maxExclusive:
		errorMessage = "min >= maxExclusive"
	case count < 1:
		errorMessage = "count is smaller than 1"
	case count > 100000:
		errorMessage = "count is larger than 100 000"
	}

	if errorMessage != "" {
		bytes, _ := json.Marshal(RandomNumberHttpError{query.Get("min"), query.Get("maxExclusive"), query.Get("count"), errorMessage})
		http.Error(w, string(bytes), http.StatusBadRequest)
		return
	}

	randomNumbers := getRandomNumbers(min, maxExclusive, count)
	bytes, _ := json.Marshal(RandomNumberHttpResponse{query.Get("min"), query.Get("maxExclusive"), query.Get("count"), randomNumbers})
	_, _ = w.Write(bytes)
}
