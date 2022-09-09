package remote

import (
	"log"
	"net/rpc"
	"tiktok/api/config"
)

var Client *rpc.Client

func Init() {
	var err error
	Client, err = rpc.Dial("tcp", config.Remote)
	if err != nil {
		log.Fatal("api.Run: ", err)
	}
}

const (
	Comment  = "Comment"
	Favorite = "Favorite"
	Feed     = "Feed"
	Publish  = "Publish"
	Relation = "Relation"
	User     = "User"
)
