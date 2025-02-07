package http

import (
	"fmt"
	"sync"
)

func Request(url string, requests int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < requests; i++ {
		fmt.Printf("%d - Requesting: %s\n", i+1, url)
	}
}
