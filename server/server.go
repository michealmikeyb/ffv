package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	tags_pb "github.com/michealmikeyb/ffv/tags"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	tags_pb.UnimplementedTagServiceServer
}

func (s *server) LikeTag(ctx context.Context, in *tags_pb.Tag) (*tags_pb.BaseResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &tags_pb.BaseResponse{Status: "oks", Error: ""}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tags_pb.RegisterTagServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
