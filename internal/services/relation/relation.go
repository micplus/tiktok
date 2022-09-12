package relation

import (
	"fmt"
	"tiktok/internal/pkg/cache"

	"github.com/gomodule/redigo/redis"
)

const (
	followPrefix   = "follow_"
	followerPrefix = "follower_"
)

func Follow(from, to int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	c.Send("MULTI")
	c.Send("SADD", fmt.Sprintf("%s%d", followPrefix, from), to)
	c.Send("SADD", fmt.Sprintf("%s%d", followerPrefix, to), from)
	_, err := c.Do("EXEC")
	return err
}

func Unfollow(from, to int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	c.Send("MULTI")
	c.Send("SREM", fmt.Sprintf("%s%d", followPrefix, from), to)
	c.Send("SREM", fmt.Sprintf("%s%d", followerPrefix, to), from)
	_, err := c.Do("EXEC")
	return err
}

func FollowsByID(id int64) ([]int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64s(c.Do("SMEMBERS", fmt.Sprintf("%s%d", followPrefix, id)))
}

func FollowersByID(id int64) ([]int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64s(c.Do("SMEMBERS", fmt.Sprintf("%s%d", followerPrefix, id)))
}

func CountFollowsByID(id int64) (int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64(c.Do("SCARD", fmt.Sprintf("%s%d", followPrefix, id)))
}

func CountFollowersByID(id int64) (int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64(c.Do("SCARD", fmt.Sprintf("%s%d", followerPrefix, id)))

}
