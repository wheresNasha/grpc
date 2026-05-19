// This file has TCP + gRPC startup

package main

import (
	"log"
	"net"
	
	pb "helloWorldGrpc/proto"
	"google.golang.org/grpc"
)

func main() {
	
	// Open TCP port 50051 and wait for connections
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// creates the actual gRPC runtime
	grpcServer := grpc.NewServer()
	
	// When SayHello request comes, route it to this struct
	pb.RegisterHelloServiceServer(
		grpcServer,
		&HelloServer{},
	)

	log.Println("gRPC server started on :50051")
	
	// grpcServer.Serve(lis): Start accepting TCP connections forever
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}