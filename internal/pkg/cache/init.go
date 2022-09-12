package cache

import (
	"tiktok/internal/config"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Cache *redis.Pool

func Init() {
	Cache = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", config.RedisPort) },
	}
}
