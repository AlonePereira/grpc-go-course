package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doLongGreet(client pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Alone"},
		{FirstName: "Giovanna"},
		{FirstName: "Laura"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", resp.Result)
}
