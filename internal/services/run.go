package services

import (
	"log"
	"net"
	"net/rpc"
	"tiktok/internal/config"
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/comment"
	"tiktok/internal/services/favorite"
	"tiktok/internal/services/feed"
	"tiktok/internal/services/publish"
	"tiktok/internal/services/relation"
	"tiktok/internal/services/user"
)

func Run() {
	config.Load()

	database.Init()

	svr := rpc.NewServer()

	registerServices(svr)

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatal("services.Run: ", err)
	}
	go svr.Accept(lis)
}

func registerServices(s *rpc.Server) {
	log.Println(s.Register(new(comment.Comment)))
	log.Println(s.Register(new(favorite.Favorite)))
	log.Println(s.Register(new(feed.Feed)))
	log.Println(s.Register(new(publish.Publish)))
	log.Println(s.Register(new(relation.Relation)))
	log.Println(s.Register(new(user.User)))
}

const (
	Comment  = "Comment"
	Favorite = "Favorite"
	Feed     = "Feed"
	Publish  = "Publish"
	Relation = "Relation"
	User     = "User"
)
