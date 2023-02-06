package main

import (
	"context"
	"log"

	pb "github.com/AlonePereira/grpc-go-course/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---createBlog was invoked---")

	blog := &pb.Blog{
		AuthorId: "Alone",
		Title:    "My first Blog",
		Content:  "Content of the first blog",
	}

	resp, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\v", resp.Id)
	return resp.Id
}
