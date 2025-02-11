package fcst

import (
	"github.com/tiagoncardoso/fc-pge-stress-test/pkg/fcst/http"
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

	for i := 0; i < params.Concurrency; i++ {
		wg.Add(1)
		go http.Request(params.Url, params.Requests, &wg)
	}

	wg.Wait()
}
