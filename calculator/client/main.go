package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/AlonePereira/grpc-go-course/calculator/proto"
)

var addr string = "localhost:5002"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)
	// doPrimeNumber(client)
	// doAvg(client)
	// doMax(client)
	doSqrt(client, -2)
}
