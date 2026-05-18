package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 200,
			IdleConnTimeout:     90 * time.Second,
		},
		Timeout: 10 * time.Second,
	}

	url := "http://localhost:9200/perf-test/_search?q=title:OpenSearch"

	concurrentRequests := 100
	var wg sync.WaitGroup
	// wg.Add(concurrentRequests)

	start := time.Now()

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			resp, err := client.Get(url)
			if err != nil {
				panic(err)
			}

			_, err = io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
			if err != nil {
				panic(err)
			}
		}()
	}

	wg.Wait()

	total := time.Since(start)

	fmt.Printf("Total time for %d concurrent REST requests: %v\n", concurrentRequests, total)
	fmt.Printf("Average per request: %v\n", total/time.Duration(concurrentRequests))
	fmt.Printf("Requests/sec: %.2f\n", float64(concurrentRequests)/total.Seconds())
}
