package main

import (
	"fmt"
	"os"
	"tiktok/api/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	fmt.Println(os.UserHomeDir())
	router.Run()
}
