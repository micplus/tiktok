package login

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func loginByUsername(username string) (*model.UserLogin, error) {
	user := new(model.UserLogin)
	stmt := `SELECT * FROM user_logins WHERE username=?;`
	if err := db.QueryRowx(stmt, username).Scan(&user); err != nil {
		return nil, err
	}
	return user, nil
}
