package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/AlonePereira/grpc-go-course/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Println("Invoking doSum")
	resp, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  15,
		SecondNumber: 15,
	})

	if err != nil {
		log.Fatalf("Erro Sum %v\n", err)
	}

	log.Printf("Sum for %v", resp.Result)
}

func doPrimeNumber(clinet pb.CalculatorServiceClient) {
	log.Println("Invoking doPrimeNumber")

	stream, err := clinet.PrimeNumber(context.Background(), &pb.PrimeRequest{Number: 120})

	if err != nil {
		log.Fatalf("Error while calling PrimeNumber: %v\n", err)
	}

	var result []int32

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		result = append(result, resp.Number)
	}

	fmt.Println(result)
}

func doAvg(client pb.CalculatorServiceClient) {
	log.Println("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := client.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %.2f\n", resp.Result)
}
