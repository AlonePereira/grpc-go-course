package main

import (
	"context"
	"log"

	pb "github.com/AlonePereira/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Alone",
	})

	if err != nil {
		log.Fatalf("Erro Greet %v\n", err)
	}

	log.Printf("Greet for %v", res.Result)
}
