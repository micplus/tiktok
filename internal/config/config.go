package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	DSN        string
	PublicPath string
)

type mysql struct {
	Host     string
	Port     int64
	Network  string
	Database string
	Username string
	Password string
	Charset  string
}

func (m *mysql) dsn() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s",
		m.Username, m.Password, m.Network, m.Host, m.Port,
		m.Database, m.Charset)
}

type public struct {
	Path string
}

type config struct {
	Mysql  mysql  `toml:"mysql"`
	Public public `toml:"public"`
}

func init() {
	detail := new(config)
	if _, err := toml.DecodeFile("/home/abc/workspace/tiktok/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	DSN = detail.Mysql.dsn()
	PublicPath = detail.Public.Path
}
