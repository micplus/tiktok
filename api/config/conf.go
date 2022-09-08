package config

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

var (
	Address string
)

type server struct {
	Host string
	Port int64
}

type config struct {
	Server server `toml:"address"`
}

func init() {
	base := os.Getenv("TIKTOK_DIR")
	detail := new(config)
	if _, err := toml.DecodeFile(base+"/home/abc/workspace/tiktok/api/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	host := detail.Server.Host
	port := strconv.FormatInt(detail.Server.Port, 10)
	Address = host + ":" + port
}
