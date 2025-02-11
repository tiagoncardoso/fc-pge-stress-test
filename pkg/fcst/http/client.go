package http

import (
	"log/slog"
	"net/http"
	"sync"
	"time"
)

type RequesStatusCode map[int]int

var (
	totalRequests  int
	requestsStatus RequesStatusCode
	mu             sync.Mutex
)

type RequestResult struct {
	TotalRequests  int
	RequestsStatus RequesStatusCode
	StartTest      time.Time
	EndTest        time.Time
}

func Request(url string, requests int, wg *sync.WaitGroup, results chan<- RequestResult) {
	startTest := time.Now()
	defer wg.Done()

	for i := 0; i < requests; i++ {
		requestStatus := httpRequest(url)

		mu.Lock()
		updateStatusResponse(requestStatus)
		totalRequests++
		mu.Unlock()

		time.Sleep(500 * time.Millisecond)
	}

	endTest := time.Now()

	results <- RequestResult{
		TotalRequests:  totalRequests,
		RequestsStatus: requestsStatus,
		StartTest:      startTest,
		EndTest:        endTest,
	}
}

func httpRequest(url string) int {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("Failed to create request", err)
		return http.StatusServiceUnavailable
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Request Failed", err)
		return http.StatusInternalServerError
	}

	return resp.StatusCode

}

func updateStatusResponse(statusCode int) {
	if requestsStatus == nil {
		requestsStatus = make(RequesStatusCode)
	}
	requestsStatus[statusCode]++
}
