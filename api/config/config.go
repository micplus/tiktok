package config

import (
	"log"
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
	detail := new(config)
	if _, err := toml.DecodeFile("/home/abc/workspace/tiktok/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	host := detail.Server.Host
	port := strconv.FormatInt(detail.Server.Port, 10)
	Address = host + ":" + port
}
