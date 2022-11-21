package main

import (
	"context"
	"flag"
	"log"
	"time"

	users_pb "github.com/michealmikeyb/ffv/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	tag  = flag.String("tag", defaultName, "tag to like")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := users_pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPost(ctx, &users_pb.GetPostRequest{UserId: "6d9c1f1d-694f-11ed-b4a5-1c4d70a146e5"})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Status: %s", r.GetPost().Url)
}
