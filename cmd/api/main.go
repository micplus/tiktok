package main

import (
	"fmt"
	"os"
	"tiktok/api/config"
	"tiktok/api/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	fmt.Println(os.Getenv("TIKTOK_DIR"))
	r := router.NewRouter()
	r.Run(config.Address)
}
