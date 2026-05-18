package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	
	pb "github.com/opensearch-project/opensearch-protobufs/go/opensearchpb"
	searchpb "github.com/opensearch-project/opensearch-protobufs/go/services"
)

func main() {
	conn, err := grpc.NewClient("localhost:9400",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := searchpb.NewSearchServiceClient(conn)

	req := &pb.SearchRequest{
		Index: []string{"perf-test"},
		Q:     proto.String("title:OpenSearch"),
	}

	var total time.Duration
	runs := 50

	for i := 0; i < runs; i++ {
		start := time.Now()

		_, err := client.Search(context.Background(), req)
		if err != nil {
			panic(err)
		}

		total += time.Since(start)
	}

	fmt.Printf("Average persistent gRPC latency: %v\n", total/time.Duration(runs))
}
