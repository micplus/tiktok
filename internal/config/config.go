package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

var (
	DSN        string
	RedisPort  string
	PublicPath string
	StaticPath string
	Address    string
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

type redis struct {
	Port int64
}

func (m *mysql) dsn() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s",
		m.Username, m.Password, m.Network, m.Host, m.Port,
		m.Database, m.Charset)
}

type server struct {
	Port int64
}

type config struct {
	Mysql  mysql  `toml:"mysql"`
	Redis  redis  `toml:"redis"`
	Server server `toml:"address"`
}

func Load() {
	base := os.Getenv("TIKTOK_DIR")
	detail := new(config)
	if _, err := toml.DecodeFile(base+"/internal/config/config.toml", detail); err != nil {
		log.Panic(err)
	}
	DSN = detail.Mysql.dsn()

	port := strconv.FormatInt(detail.Server.Port, 10)
	Address = ":" + port

	RedisPort = ":" + strconv.FormatInt(detail.Redis.Port, 10)
}
