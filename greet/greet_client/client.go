package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
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
	c := greetpb.NewGreetSreviceClient(conn)

	doUnary(c)
}

func doUnary(c greetpb.GreetSreviceClient) {

	fmt.Println("Starting do Unary RPC")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Nagaraj",
			LastName:  "Suaragani Venkata",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("Response from Greet %v", res.Result)

}
