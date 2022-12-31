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
