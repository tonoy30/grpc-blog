package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-init/src/blog/blogpb"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	bc := blogpb.NewBlogServiceClient(conn)

	res, err := bc.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "tonoy",
			Title:    "My First Blog",
			Content:  "Content of the first blog",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)

}
