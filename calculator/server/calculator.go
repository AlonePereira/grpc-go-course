package main

import (
	"context"
	"io"
	"log"

	pb "github.com/AlonePereira/grpc-go-course/calculator/proto"
)

type server struct {
	pb.CalculatorServiceServer
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum invoked with request %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *server) PrimeNumber(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeNumberServer) error {

	divisor := int32(2)
	number := in.Number

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Number: divisor,
			})

			number = number / divisor
		} else {
			divisor += 1
		}
	}

	return nil
}

func (s *server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("invoking avg")

	var numbers []int64
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			sum := 0

			for _, number := range numbers {
				sum += int(number)
			}

			result := float32(sum) / float32(len(numbers))

			return stream.SendAndClose(&pb.AvgResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req %v\n", req)

		numbers = append(numbers, req.Number)
	}
}
