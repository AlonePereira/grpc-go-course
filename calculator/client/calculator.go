package main

import (
	"context"
	"fmt"
	"io"
	"log"

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
