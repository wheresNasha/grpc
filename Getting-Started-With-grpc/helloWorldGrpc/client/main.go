package main

// Client only cares about API contract. Even if tomorrow the handler.go changes and 
// SayHello in a different style still client wont change, 
// as proto file where rpc of SayHello is that is the API contract
// Clients depend on contracts/interfaces, not server internals.

import (
    "context" // used to control timeout / cancellation for RPC calls
    "log"      // for logging errors and output
    "time"     // used to set request timeout

    pb "helloWorldGrpc/proto" // generated protobuf + gRPC code (client + messages)
    "google.golang.org/grpc"  // core gRPC library
    "google.golang.org/grpc/credentials/insecure" // disables TLS for local dev
)

func main() {

    // STEP 1: Create a connection to the gRPC server
    // This opens a TCP connection to localhost:50051
    conn, err := grpc.Dial(
        "localhost:50051",
        grpc.WithTransportCredentials(insecure.NewCredentials()), // no TLS (dev only)
    )
    if err != nil {
        log.Fatalf("did not connect: %v", err) // stop program if connection fails
    }
    defer conn.Close() // ensures connection is closed when main exits

    // STEP 2: Create a gRPC client from generated code
    // This is a "stub" that knows how to call SayHello over the network
    client := pb.NewHelloServiceClient(conn)

    // STEP 3: Create a context with timeout
    // If server doesn't respond in 1 second, request is cancelled
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel() // releases resources tied to context

    // STEP 4: Call remote function (RPC call)
    // This looks like a normal function call but is actually:
    // serialize -> send over TCP -> wait -> deserialize response
    res, err := client.SayHello(
        ctx,
        &pb.HelloRequest{Name: "Sakshi"}, // request payload
    )
    if err != nil {
        log.Fatalf("could not greet: %v", err) // handle RPC failure
    }

    // STEP 5: Print server response
    log.Printf("Response: %s", res.Message)
}