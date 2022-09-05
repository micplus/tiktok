package router

import (
	"tiktok/config"
	"tiktok/internal/controller/favorite"
	"tiktok/internal/controller/feed"
	"tiktok/internal/controller/publish"
	"tiktok/internal/controller/user"
	mw "tiktok/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.Static("/static", config.Detail.Static.Path)

	g := r.Group("/douyin")

	g.GET("/feed", feed.Feed)
	g.POST("/user/register", user.Register)
	g.POST("/user/login", user.Login)

	g.GET("/user", mw.Auth(), user.User)
	g.POST("/publish/action", mw.Auth(), publish.Action)
	g.GET("/publish/list", mw.Auth(), publish.List)

	g.POST("/favorite/action", mw.Auth(), favorite.Action)

	return r
}
