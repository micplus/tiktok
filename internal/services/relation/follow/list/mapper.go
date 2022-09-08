package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"

	"github.com/jmoiron/sqlx"
)

var db = database.DB

func followIDsByUserID(userID int64) ([]int64, error) {
	ids := []int64{}
	stmt := `SELECT DISTINCT follow_id FROM user_follows WHERE user_id=?;`
	err := db.Select(&ids, stmt, userID)
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

func isFollowsOfUserID(ids []int64, userID int64) (map[int64]struct{}, error) {
	f := make(map[int64]struct{})
	follows := []int64{}
	stmt := `SELECT follow_id FROM user_follows 
		WHERE user_id=? AND follow_id IN (?);`
	query, args, err := sqlx.In(stmt, userID, ids)
	if err != nil {
		return f, err
	}
	rows, err := db.Queryx(db.Rebind(query), args...)
	if err != nil {
		return f, err
	}
	for rows.Next() {
		var followID int64
		rows.Scan(&followID)
		follows = append(follows, followID)
	}

	for _, id := range follows {
		f[id] = struct{}{}
	}
	return f, nil
}
