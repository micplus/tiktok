package favorite

import (
	"fmt"
	"tiktok/internal/pkg/cache"

	"github.com/gomodule/redigo/redis"
)

const (
	userPrefix  = "favorite_user_"
	videoPrefix = "favorite_video_"
)

func Favorite(from, to int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	c.Send("MULTI")
	c.Send("SADD", fmt.Sprintf("%s%d", userPrefix, from), to)
	c.Send("SADD", fmt.Sprintf("%s%d", videoPrefix, to), from)
	_, err := c.Do("EXEC")
	return err
}

func Unfavorite(from, to int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	c.Send("MULTI")
	c.Send("SREM", fmt.Sprintf("%s%d", userPrefix, from), to)
	c.Send("SREM", fmt.Sprintf("%s%d", videoPrefix, to), from)
	_, err := c.Do("EXEC")
	return err
}

func FavoritesByUserID(id int64) ([]int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64s(c.Do("SMEMBERS", fmt.Sprintf("%s%d", userPrefix, id)))
}

func FavoritedsByVideoID(id int64) ([]int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64s(c.Do("SMEMBERS", fmt.Sprintf("%s%d", videoPrefix, id)))
}

func CountFavoritedsByVideoID(id int64) (int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Int64(c.Do("SCARD", fmt.Sprintf("%s%d", videoPrefix, id)))
}

func CountFavoritedsByVideoIDs(ids []int64) ([]int64, error) {
	c := cache.Cache.Get()
	defer c.Close()
	c.Send("MULTI")
	for _, id := range ids {
		c.Send("SCARD", fmt.Sprintf("%s%d", videoPrefix, id))
	}
	return redis.Int64s(c.Do("EXEC"))
}
