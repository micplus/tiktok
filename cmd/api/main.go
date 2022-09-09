package main

import (
	"tiktok/api"
	"tiktok/api/config"
	"tiktok/api/router"
)

func main() {
	api.Run()
	r := router.New()
	r.Run(config.Address)
}
