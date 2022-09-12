package login

import (
	"fmt"
	"tiktok/internal/model"
	"tiktok/internal/pkg/cache"
	"tiktok/internal/pkg/database"
	"time"

	"github.com/gomodule/redigo/redis"
)

const prefix = "login_"

var (
	expire       = int64(24 * time.Hour / time.Second)
	logoutRemain = int64(1 * time.Minute / time.Second)
)

func CheckCache(id int64) (bool, error) {
	c := cache.Cache.Get()
	defer c.Close()
	return redis.Bool(c.Do("GET", fmt.Sprintf("%s%d", prefix, id)))
}

func SetCache(id int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	_, err := c.Do("SETEX", fmt.Sprintf("%s%d", prefix, id), expire, 1)
	return err
}

func RemoveCache(id int64) error {
	c := cache.Cache.Get()
	defer c.Close()
	_, err := c.Do("SETEX", fmt.Sprintf("%s%d", prefix, id), logoutRemain, 1)
	return err
}

func ByUsername(username string) (*model.UserLogin, error) {
	db := database.DB
	users := []model.UserLogin{}
	stmt := `SELECT * FROM user_logins WHERE username=?;`
	if err := db.Select(&users, stmt, username); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}

func CountByUsername(username string) (int64, error) {
	db := database.DB
	var count int64
	stmt := `SELECT COUNT(*) FROM user_logins WHERE username=?;`
	err := db.QueryRowx(stmt, username).Scan(&count)
	return count, err
}

func Insert(ul *model.UserLogin) (int64, error) {
	db := database.DB
	stmt := `INSERT INTO user_logins(
		user_id, username, password, salt, created_at, modified_at
	) VALUES (?, ?, ?, ?, ?, ?);`
	res, err := db.Exec(stmt, ul.UserID, ul.Username, ul.Password,
		ul.Salt, ul.CreatedAt, ul.ModifiedAt)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
