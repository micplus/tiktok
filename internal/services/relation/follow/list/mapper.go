package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func followIDsByUserID(userID int64) ([]int64, error) {
	ids := []int64{}
	stmt := `SELECT DISTINCT follow_id FROM user_follows WHERE user_id=?;`
	err := db.Select(ids, stmt, userID)
	return ids, err
}

func usersByIDs(ids []int64) ([]model.User, error) {
	users := []model.User{}
	stmt := `SELECT * FROM users WHERE id IN (?);`
	err := db.Select(users, stmt, ids)
	return users, err
}

func isFollowsOfUserID(ids []int64, userID int64) (map[int64]struct{}, error) {
	f := make(map[int64]struct{})
	follows := []int64{}
	stmt := `SELECT follow_id FROM user_follows 
		WHERE user_id=? AND follow_id IN (?);`
	err := db.Select(follows, stmt, userID, ids)
	if err != nil {
		return f, err
	}

	for _, id := range follows {
		f[id] = struct{}{}
	}
	return f, nil
}
