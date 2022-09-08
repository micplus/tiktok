package login

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func loginByUsername(username string) (*model.UserLogin, error) {
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
