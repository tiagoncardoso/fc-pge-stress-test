package fcst

import (
	"github.com/tiagoncardoso/fc-pge-stress-test/pkg/fcst/http"
	"log/slog"
	"sync"
	"time"
)

type StressTestParams struct {
	Requests    int
	Concurrency int
	Url         string
}

type StressTestReport struct {
	StartTest      time.Time
	EndTest        time.Time
	TotalRequests  int
	RequestsStatus http.RequesStatusCode
}

func FcStress(params StressTestParams) {
	initRoutines(params)
}

func initRoutines(params StressTestParams) {
	var wg sync.WaitGroup
	results := make(chan http.RequestResult, params.Concurrency)

	for i := 0; i < params.Concurrency; i++ {
		wg.Add(1)
		go http.Request(params.Url, params.Requests, &wg, results)
	}

	wg.Wait()
	close(results)

	var report StressTestReport
	for result := range results {
		report.TotalRequests += result.TotalRequests
		report.StartTest = result.StartTest
		report.EndTest = result.EndTest
		for status, count := range result.RequestsStatus {
			report.RequestsStatus[status] += count
		}
	}

	slog.Info("Finish Report", "Total Requests", report.TotalRequests)
	slog.Info("Finish Report", "Requests Status", report.RequestsStatus)
	slog.Info("Finish Report", "Start Test", report.StartTest)
	slog.Info("Finish Report", "End Test", report.EndTest)
}
