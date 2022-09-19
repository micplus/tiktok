package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	Address    string
	Remote     string
	StaticAddr string
	StaticDir  string
)

type server struct {
	Host string
	Port int64
}

type remote struct {
	Host string
	Port int64
}

type static struct {
	Network string
	Host    string
	Port    int64
	Dir     string
}

type config struct {
	Server server `toml:"address"`
	Remote remote `toml:"remote"`
	Static static `toml:"static"`
}

func Load() {
	base := os.Getenv("TIKTOK_DIR")
	detail := new(config)
	if _, err := toml.DecodeFile(base+"/api/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	port := detail.Server.Port
	Address = fmt.Sprintf(":%d", port)

	host := detail.Remote.Host
	port = detail.Remote.Port
	Remote = fmt.Sprintf("%s:%d", host, port)

	network := detail.Static.Network
	host = detail.Static.Host
	port = detail.Static.Port
	StaticAddr = fmt.Sprintf("%s://%s:%d", network, host, port)
	StaticDir = detail.Static.Dir
}
