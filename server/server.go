package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"github.com/gocql/gocql"
	tags_pb "github.com/michealmikeyb/ffv/tags"
	users_pb "github.com/michealmikeyb/ffv/users"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	tags_pb.UnimplementedTagServiceServer
	users_pb.UnimplementedUserServiceServer
}

type Post struct {
	Url    string   `cql:"url"`
	Tags   []string `cql:"tags"`
	Source string   `cql:"source"`
}

func GetCassandraSession() (*gocql.Session, error) {
	cluster := gocql.NewCluster(os.Getenv("CASS_HOST"))
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: os.Getenv("CASS_USER"),
		Password: os.Getenv("CASS_PASS"),
	}
	return cluster.CreateSession()

}

func (s *server) LikeTag(ctx context.Context, in *tags_pb.Tag) (*tags_pb.TagBaseResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &tags_pb.TagBaseResponse{Status: "oks", Error: ""}, nil
}

func (s *server) AddUser(ctx context.Context, user *users_pb.AddUserRequest) (*users_pb.UserBaseResponse, error) {
	log.Printf("Received: %v", user.MastadonId)
	session, err := GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	err = session.Query(`INSERT INTO ffv.user (user_id, mastadon_id, mastadon_username) VALUES (?, ?, ?)`, gocql.TimeUUID(), user.MastadonId, user.MastadonUsername).WithContext(ctx).Exec()
	if err != nil {
		return nil, err
	}
	return &users_pb.UserBaseResponse{Status: "ok", Error: ""}, nil
}

func (s *server) GetPost(ctx context.Context, user *users_pb.GetPostRequest) (*users_pb.GetPostResponse, error) {
	log.Printf("Received: %v", user.UserId)
	session, err := GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	tag_list_scanner := session.Query(`SELECT tag_name, weight, tag_id, source, seen FROM ffv.tag_list WHERE user_id = ?`, user.UserId).WithContext(ctx).Iter().Scanner()
	selection_list := []string{}
	for tag_list_scanner.Next() {
		var (
			tag_name string
			weight   float32
			tag_id   string
			source   string
			seen     []Post
		)
		err := tag_list_scanner.Scan(&tag_name, &weight, &tag_id, &source, &seen)
		for i := 0; i < int(weight); i++ {
			selection_list = append(selection_list, tag_id)
		}
		if err != nil {
			return nil, err
		}
	}
	selection_index := rand.Intn(len(selection_list))
	selected_tag := selection_list[selection_index]
	var received_tag string

	var buffer []Post
	err = session.Query(`SELECT buffer, tag_id FROM ffv.tag WHERE tag_id = ? LIMIT 1`, selected_tag).WithContext(ctx).Consistency(gocql.One).Scan(&buffer, &received_tag)
	if err != nil {
		return nil, err
	}
	selected_post := buffer[0]
	return &users_pb.GetPostResponse{Post: &users_pb.Post{
		Url:    selected_post.Url,
		Tags:   selected_post.Tags,
		Source: selected_post.Source,
	}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	tags_pb.RegisterTagServiceServer(s, &server{})
	users_pb.RegisterUserServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
