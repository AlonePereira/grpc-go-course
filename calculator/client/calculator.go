package main

import (
	"context"
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
