package main

import (
	"context"
	"io"
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

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Alone",
	}

	stream, err := client.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
