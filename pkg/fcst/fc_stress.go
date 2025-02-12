package fcst

import (
	"fmt"
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
	TestDuration   string
	TotalRequests  int
	RequestsStatus http.RequesStatusCode
}

func FcStress(params StressTestParams) {
	report := StressTestReport{
		StartTest: time.Now(),
	}

	report = initRoutines(params)

	report.EndTest = time.Now()

	slog.Info("Finish Report", "Total Requests", report.TotalRequests)
	slog.Info("Finish Report", "Requests Status", report.RequestsStatus)
	slog.Info("Finish Report", "Start Test", report.StartTest)
	slog.Info("Finish Report", "End Test", report.EndTest)
	slog.Info("Finish Report", "Test Duration", report.TestDuration)
}

func initRoutines(params StressTestParams) StressTestReport {
	var wg sync.WaitGroup
	results := make(chan http.RequestResult, params.Concurrency)

	for i := 0; i < params.Concurrency; i++ {
		wg.Add(1)
		go http.Request(params.Url, params.Requests, &wg, results)
	}

	wg.Wait()
	close(results)

	var report StressTestReport
	report.RequestsStatus = make(http.RequesStatusCode)

	for result := range results {
		report.TotalRequests = result.TotalRequests
		report.StartTest = result.StartTest
		report.EndTest = result.EndTest
		report.TestDuration = getDuration(result.StartTest, result.EndTest)
		report.RequestsStatus = result.RequestsStatus
	}

	return report
}

func getDuration(start, end time.Time) string {
	duration := end.Sub(start)

	return fmt.Sprintf("%02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60, int(duration.Seconds())%60)
}
