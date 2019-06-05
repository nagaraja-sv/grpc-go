package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"

	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("hello I m a client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect %v", err)

	}
	defer conn.Close()
	c := calculatorpb.NewCalculatorServiceClient(conn)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {

	fmt.Println("Starting do Unary RPC")

	req := &calculatorpb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("Response from Sum %v", res.Result)

}
