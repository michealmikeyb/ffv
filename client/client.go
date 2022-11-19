package main

import (
	"context"
	"flag"
	"log"
	"time"

	tags_pb "github.com/michealmikeyb/ffv/tags"
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
	c := tags_pb.NewTagServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.LikeTag(ctx, &tags_pb.Tag{Name: *tag})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Status: %s", r.GetStatus())
}
