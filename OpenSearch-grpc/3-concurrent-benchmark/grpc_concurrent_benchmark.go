package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	pb "github.com/opensearch-project/opensearch-protobufs/go/opensearchpb"
	searchpb "github.com/opensearch-project/opensearch-protobufs/go/services"
)

func main() {
	
	conn, err := grpc.NewClient(
		"localhost:9400",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := searchpb.NewSearchServiceClient(conn)

	req := &pb.SearchRequest{
		Index: []string{"perf-test"},
		Q:     proto.String("title:OpenSearch"),
	}

	concurrentRequests := 100
	var wg sync.WaitGroup
	// wg.Add(concurrentRequests)

	start := time.Now()

	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_, err := client.Search(context.Background(), req)
			if err != nil {
				panic(err)
			}
		}()
	}

	wg.Wait()

	total := time.Since(start)

	fmt.Printf("Total time for %d concurrent gRPC requests: %v\n", concurrentRequests, total)
	fmt.Printf("Average per request: %v\n", total/time.Duration(concurrentRequests))
	fmt.Printf("Requests/sec: %.2f\n", float64(concurrentRequests)/total.Seconds())
}
