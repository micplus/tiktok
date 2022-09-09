package api

import (
	"tiktok/api/config"
	"tiktok/api/remote"
	"tiktok/api/router"
)

func Run() {
	config.Load()

	remote.Init()

	r := router.New()
	r.Run(config.Address)
}
