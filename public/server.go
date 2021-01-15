package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"grpc-init/src/blog"
	"grpc-init/src/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"
)



func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalln(err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening to port 50051: %v\n", err)
	}

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &blog.Server{})

	go func() {
		log.Println("Starting grpc server")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Error serving grpc: %v\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Stopping grpc server")
	s.Stop()
	lis.Close()
	client.Disconnect(context.TODO())

}
