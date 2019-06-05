package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Greet function was Invoked with", req)
	// Getting the request
	firstname := req.GetGreeting().GetFirstName()

	result := "Hello" + firstname

	// Sending those values to the response
	res := &greetpb.GreetResponse{Result: result}

	return res, nil
}

func main() {

	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetSreviceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
