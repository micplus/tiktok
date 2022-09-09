package config

import (
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

var (
	Address string
	Remote  string
)

type server struct {
	Port int64
}

type remote struct {
	Host string
	Port int64
}

type config struct {
	Server server `toml:"address"`
	Remote remote `toml:"remote"`
}

func Load() {
	base := os.Getenv("TIKTOK_DIR")
	detail := new(config)
	if _, err := toml.DecodeFile(base+"/api/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	port := strconv.FormatInt(detail.Server.Port, 10)
	Address = ":" + port

	host := detail.Remote.Host
	port = strconv.FormatInt(detail.Remote.Port, 10)
	Remote = host + ":" + port
}
