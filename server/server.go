package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"

	"github.com/gocql/gocql"
	models "github.com/michealmikeyb/ffv/models"
	tags_pb "github.com/michealmikeyb/ffv/tags"
	users_pb "github.com/michealmikeyb/ffv/users"
	utils "github.com/michealmikeyb/ffv/utils"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	tags_pb.UnimplementedTagServiceServer
	users_pb.UnimplementedUserServiceServer
}

func (s *server) DislikePost(ctx context.Context, post_request *tags_pb.RatePostRequest) (*tags_pb.TagBaseResponse, error) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	for _, tag := range post_request.GetPost().Tags {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)

		err := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ? AND tag_name = ? AND source = ?`, post_request.UserId, tag, post_request.Post.Source).WithContext(ctx).Scan(&tag_name, &weight, &source, &seen)
		if err == gocql.ErrNotFound {
			tag_name = tag
			weight = 0
			source = post_request.Post.Source
			seen = make([]models.Post, 0)
		} else if err != nil {
			log.Fatal(err)
		}
		if weight > 0 {
			weight = weight - 1
		}
		current_post := models.Post{
			Url:     post_request.Post.Url,
			Source:  post_request.Post.Source,
			Tags:    post_request.Post.Tags,
			Author:  post_request.Post.Author,
			Likes:   int(post_request.Post.Likes),
			Content: post_request.Post.Content,
		}
		seen = append(seen, current_post)
		err = session.Query(`INSERT INTO ffv.tag_list (tag_name, weight, source, seen, user_id) VALUES (?, ?, ?, ?, ?)`, tag_name, weight, source, seen, post_request.UserId).Consistency(gocql.One).Exec()
		if err != nil {
			log.Fatal(err)
		}

	}

	return &tags_pb.TagBaseResponse{Status: "oks", Error: ""}, nil
}

func (s *server) LikePost(ctx context.Context, post_request *tags_pb.RatePostRequest) (*tags_pb.TagBaseResponse, error) {
	session, err := utils.GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	for _, tag := range post_request.GetPost().Tags {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)

		err := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ? AND tag_name = ? AND source = ?`, post_request.UserId, tag, post_request.Post.Source).WithContext(ctx).Scan(&tag_name, &weight, &source, &seen)
		if err == gocql.ErrNotFound {
			tag_name = tag
			weight = 0
			source = post_request.Post.Source
			seen = make([]models.Post, 0)
		} else if err != nil {
			log.Fatal(err)
		}
		weight = weight + 1
		current_post := models.Post{
			Url:     post_request.Post.Url,
			Source:  post_request.Post.Source,
			Tags:    post_request.Post.Tags,
			Author:  post_request.Post.Author,
			Likes:   int(post_request.Post.Likes),
			Content: post_request.Post.Content,
		}
		seen = append(seen, current_post)
		err = session.Query(`INSERT INTO ffv.tag_list (tag_name, weight, source, seen, user_id) VALUES (?, ?, ?, ?, ?)`, tag_name, weight, source, seen, post_request.UserId).Consistency(gocql.One).Exec()
		if err != nil {
			log.Fatal(err)
		}

	}

	return &tags_pb.TagBaseResponse{Status: "oks", Error: ""}, nil
}

func (s *server) AddUser(ctx context.Context, user *users_pb.AddUserRequest) (*users_pb.AddUserResponse, error) {
	log.Printf("Received: %v", user.MastodonId)
	session, err := utils.GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	user_id := gocql.TimeUUID()
	err = session.Query(`INSERT INTO ffv.user (user_id, mastodon_id, mastodon_username) VALUES (?, ?, ?)`, user_id, user.MastodonId, user.MastodonUsername).WithContext(ctx).Exec()
	if err != nil {
		return nil, err
	}
	err = session.Query(`INSERT INTO ffv.tag_list (user_id, tag_name, weight, source) VALUES (?, ?, ?, ?)`, user_id, "popular", 20, "mastodon").WithContext(ctx).Exec()
	if err != nil {
		return nil, err
	}
	return &users_pb.AddUserResponse{UserId: user_id.String()}, nil
}

func (s *server) GetUser(ctx context.Context, user *users_pb.GetUserRequest) (*users_pb.AddUserResponse, error) {
	log.Printf("Received: %v", user.MastodonId)
	session, err := utils.GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	var user_id gocql.UUID
	err = session.Query(`SELECT  user_id FROM ffv.user WHERE mastodon_id = ?`, user.MastodonId).WithContext(ctx).Scan(&user_id)
	if err != nil {
		return nil, err
	}
	return &users_pb.AddUserResponse{UserId: user_id.String()}, nil
}
func (s *server) GetPost(ctx context.Context, user *users_pb.GetPostRequest) (*users_pb.GetPostResponse, error) {
	log.Printf("Received: %v", user.UserId)
	session, err := utils.GetCassandraSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	tag_list_scanner := session.Query(`SELECT tag_name, weight, source, seen FROM ffv.tag_list WHERE user_id = ?`, user.UserId).WithContext(ctx).Iter().Scanner()
	selection_list := []string{}
	for tag_list_scanner.Next() {
		var (
			tag_name string
			weight   int
			source   string
			seen     []models.Post
		)
		err := tag_list_scanner.Scan(&tag_name, &weight, &source, &seen)
		for i := 0; i < int(weight); i++ {
			selection_list = append(selection_list, tag_name)
		}
		if err != nil {
			return nil, err
		}
	}
	selection_index := rand.Intn(len(selection_list))
	selected_tag := selection_list[selection_index]
	var received_tag string

	var buffer []models.Post
	err = session.Query(`SELECT buffer, name FROM ffv.tag WHERE name = ? LIMIT 1`, selected_tag).WithContext(ctx).Consistency(gocql.One).Scan(&buffer, &received_tag)
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
