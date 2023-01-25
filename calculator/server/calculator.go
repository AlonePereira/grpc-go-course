package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"

	pb "github.com/AlonePereira/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("invoking Max")

	max := int64(0)

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > max {
			max = req.Number
			err = stream.Send(&pb.MaxResponse{
				Result: max,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}

func (s *server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil

}
