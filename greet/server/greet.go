package main

import (
	"context"
	"log"

	pb "github.com/AlonePereira/grpc-go-course/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Request %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
