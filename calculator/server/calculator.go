package main

import (
	"context"
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
