package user

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func userByID(id int64) (*model.User, error) {
	user := new(model.User)
	stmt := `SELECT * FROM users WHERE id=?;`
	if err := db.Select(user, stmt, id); err != nil {
		return nil, err
	}
	return user, nil
}

func countFollowsByID(id int64) (int64, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM user_follows WHERE user_id=?;`
	if err := db.QueryRowx(stmt, id).Scan(&count); err != nil {
		return count, err
	}
	return count, nil
}

func countFollowersByID(id int64) (int64, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM user_follows WHERE follow_id=?;`
	if err := db.QueryRowx(stmt, id).Scan(&count); err != nil {
		return count, err
	}
	return count, nil
}

func isFollowByID(from, to int64) (bool, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM user_follows WHERE user_id=? AND follow_id=?;`
	if err := db.QueryRowx(stmt, from, to).Scan(&count); err != nil {
		return false, err
	}
	return count != 0, nil
}
