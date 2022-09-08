package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"

	"github.com/jmoiron/sqlx"
)

var db = database.DB

func followerIDsByUserID(userID int64) ([]int64, error) {
	ids := []int64{}
	stmt := `SELECT DISTINCT user_id FROM user_follows WHERE follow_id=?;`
	rows, err := db.Queryx(stmt, userID)
	if err != nil {
		return ids, err
	}
	for rows.Next() {
		var userID int64
		rows.Scan(&userID)
		ids = append(ids, userID)
	}
	return ids, err
}

func usersByIDs(ids []int64) ([]model.User, error) {
	users := []model.User{}
	stmt := `SELECT * FROM users WHERE id IN (?);`
	query, args, err := sqlx.In(stmt, ids)
	if err != nil {
		return users, err
	}
	err = db.Select(&users, db.Rebind(query), args...)
	return users, err
}

func isFollowersOfUserID(ids []int64, userID int64) (map[int64]struct{}, error) {
	f := make(map[int64]struct{})
	followers := []int64{}
	stmt := `SELECT user_id FROM user_follows 
		WHERE follow_id=? AND user_id IN (?);`
	query, args, err := sqlx.In(stmt, userID, ids)
	if err != nil {
		return f, err
	}
	rows, err := db.Queryx(db.Rebind(query), args...)
	if err != nil {
		return f, err
	}
	for rows.Next() {
		var followerID int64
		rows.Scan(&followerID)
		followers = append(followers, followerID)
	}

	for _, id := range followers {
		f[id] = struct{}{}
	}
	return f, nil
}
