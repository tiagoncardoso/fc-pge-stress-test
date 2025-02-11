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

func Request(url string, requests int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requests; i++ {
		requestStatus := httpRequest(url)

		mu.Lock()
		updateStatusResponse(requestStatus)
		totalRequests++
		mu.Unlock()

		time.Sleep(500 * time.Millisecond)
	}

	slog.Info("Finish Report", "Total Requests", totalRequests)
	slog.Info("Finish Report", "Requests Status", requestsStatus)
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
