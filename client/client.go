package main

import (
	"context"
	"flag"
	"log"
	"time"

	tags_pb "github.com/michealmikeyb/ffv/tags"
	users_pb "github.com/michealmikeyb/ffv/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:50051", "the address to connect to")
	tag  = flag.String("tag", defaultName, "tag to like")
)

func addUser(mastodon_id string) {

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := users_pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddUser(ctx, &users_pb.AddUserRequest{MastodonUsername: "test-user", MastodonId: mastodon_id})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("User id: %s", r.GetUserId())
}

func getUser(mastodon_id string) {

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := users_pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUser(ctx, &users_pb.GetUserRequest{MastodonId: "test"})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("User id: %s", r.GetUserId())
}

func getPost(user_id string) {

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := users_pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPost(ctx, &users_pb.GetPostRequest{UserId: user_id})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Content: %s", r.Post.GetContent())
}
func likeTag(user_id string, tag_name string) {

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := tags_pb.NewTagServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	rate_post_request := tags_pb.RatePostRequest{
		UserId: user_id,
		Post: &tags_pb.Post{
			Url:     "test.com",
			Likes:   0,
			Source:  "mastodon",
			Tags:    []string{tag_name},
			Content: "test content",
			Author:  "test author",
		},
	}
	r, err := c.LikePost(ctx, &rate_post_request)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Status: %s", r.GetStatus())
}
func main() {
	flag.Parse()
	getPost("9e1fc557-6fbb-11ed-a147-0242ac11000b")
	// Set up a connection to the server.
}
