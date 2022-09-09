package main

import (
	"tiktok/internal/services"
)

func main() {
	services.Run()
	select {}
}
