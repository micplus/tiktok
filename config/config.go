package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host     string
	Port     int64
	Network  string
	Database string
	Username string
	Password string
	Charset  string
}

type Redis struct {
	Host string
	Port int64
}

type Server struct {
	Host string
	Port int64
}

type Static struct {
	Path string
}

type Config struct {
	Mysql  Mysql  `toml:"mysql"`
	Redis  Redis  `toml:"redis"`
	Server Server `toml:"server"`
	Static Static `toml:"static"`
}

func (c *Config) DataSourceName() string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Network,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.Database,
		c.Mysql.Charset)
}

var Detail *Config

func init() {
	Detail = new(Config)
	if _, err := toml.DecodeFile("/home/abc/workspace/tiktok/config/config.toml", Detail); err != nil {
		log.Panic(err)
	}
}
