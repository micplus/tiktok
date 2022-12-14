package router

import (
	"tiktok/api/config"
	"tiktok/api/handlers/comment"
	"tiktok/api/handlers/favorite"
	"tiktok/api/handlers/feed"
	"tiktok/api/handlers/publish"
	"tiktok/api/handlers/relation"
	"tiktok/api/handlers/user"
	mw "tiktok/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	r.Static("/static/", config.StaticDir)

	g := r.Group("/douyin")
	{
		feeds := g.Group("/feed").Use(mw.AuthNoAbort())
		{
			feeds.GET("/", feed.Feed)
		}
		users := g.Group("/user")
		{
			users.GET("/", mw.Auth(), user.User)
			users.POST("/register/", user.Register)
			users.POST("/login/", user.Login)
		}
		publishes := g.Group("/publish").Use(mw.Auth())
		{
			publishes.POST("/action/", publish.Action)
			publishes.GET("/list/", publish.List)
		}
		favorites := g.Group("/favorite").Use(mw.Auth())
		{
			favorites.POST("/action/", favorite.Action)
			favorites.GET("/list/", favorite.List)
		}
		comments := g.Group("/comment").Use(mw.Auth())
		{
			comments.POST("/action/", comment.Action)
			comments.GET("/list/", comment.List)
		}
		relations := g.Group("/relation").Use(mw.Auth())
		{
			relations.POST("/action/", relation.Action)
			relations.GET("/follow/list/", relation.Follows)
			relations.GET("/follower/list/", relation.Followers)
		}
	}
	return r
}
