package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Persistent HTTP client with keep-alive
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     90 * time.Second,
		},
		Timeout: 10 * time.Second,
	}

	url := "http://localhost:9200/perf-test/_search?q=title:OpenSearch"

	var total time.Duration
	runs := 50

	for i := 0; i < runs; i++ {
		start := time.Now()

		resp, err := client.Get(url)
		if err != nil {
			panic(err)
		}

		// Fully read response body so connection can be reused
		_, err = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		total += time.Since(start)
	}

	fmt.Printf("Average persistent REST latency: %v\n", total/time.Duration(runs))
}
