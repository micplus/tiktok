package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func followerIDsByUserID(userID int64) ([]int64, error) {
	ids := []int64{}
	stmt := `SELECT DISTINCT user_id FROM user_follows WHERE follow_id=?;`
	err := db.Select(ids, stmt, userID)
	return ids, err
}

func usersByIDs(ids []int64) ([]model.User, error) {
	users := []model.User{}
	stmt := `SELECT * FROM users WHERE id IN (?);`
	err := db.Select(users, stmt, ids)
	return users, err
}

func isFollowersOfUserID(ids []int64, userID int64) (map[int64]struct{}, error) {
	f := make(map[int64]struct{})
	followers := []int64{}
	stmt := `SELECT user_id FROM user_follows 
		WHERE follow_id=? AND user_id IN (?);`
	err := db.Select(followers, stmt, userID, ids)
	if err != nil {
		return f, err
	}

	for _, id := range followers {
		f[id] = struct{}{}
	}
	return f, nil
}
