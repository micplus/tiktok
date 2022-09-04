package main

import (
	"strconv"
	"tiktok/config"
	"tiktok/internal/router"
)

func main() {
	port := strconv.FormatInt(config.Detail.Server.Port, 10)

	r := router.Init()
	r.Run(":" + port)
}
