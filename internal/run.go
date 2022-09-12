package internal

import (
	"log"
	"net"
	"net/rpc"
	"tiktok/internal/config"
	"tiktok/internal/controllers/comment"
	"tiktok/internal/controllers/favorite"
	"tiktok/internal/controllers/feed"
	"tiktok/internal/controllers/publish"
	"tiktok/internal/controllers/relation"
	"tiktok/internal/controllers/user"
	"tiktok/internal/pkg/cache"
	"tiktok/internal/pkg/database"
)

func Run() {
	config.Load()

	database.Init()
	cache.Init()

	svr := rpc.NewServer()

	registerServices(svr)

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatal("main: ", err)
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
