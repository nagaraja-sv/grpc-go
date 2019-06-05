package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Println("Greet function was Invoked with", req)
	// Getting the request

	x := req.GetFirstNumber()
	y := req.GetSecondNumber()
	result := x + y

	// Sending those values to the response
	res := &calculatorpb.SumResponse{Result: result}

	return res, nil
}

func main() {
	fmt.Println("Hello Calculator")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
