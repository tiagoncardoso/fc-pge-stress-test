package http

import (
	"crypto/tls"
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
		slog.Info("STRESS TEST", "Request count", totalRequests, "Requested URL", url, "Status Code", requestStatus)

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
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("Service Error", "Failed to create request", err)
		return http.StatusServiceUnavailable
	}

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Response Error", "Request Failed", err)
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
