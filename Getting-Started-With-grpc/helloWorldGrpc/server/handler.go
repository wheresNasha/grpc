// This file acts like a Controller : Request -> Response business logic
package main

import (
	"context"
	pb "helloWorldGrpc/proto"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

// If somebody calls SayHello, return this response
// This is the business logic of SayHello() protoc only has interface method: abstract
// We need to write the actual implementation : of how do we wish to SayHello
func (s *HelloServer) SayHello(	ctx context.Context, req *pb.HelloRequest,) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}