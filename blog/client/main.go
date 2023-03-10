package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/AlonePereira/grpc-go-course/blog/proto"
)

var addr string = "localhost:5003"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	id := createBlog(client)

	readBlog(client, id) //valid
	// readBlog(client, "dajkhs") //invalid
	// updateBlog(client, id)
	listBlog(client)
	deleteBlog(client, id)
	deleteBlog(client, id)
}
