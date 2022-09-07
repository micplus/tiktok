package register

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func countUsername(username string) (int64, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM user_logins WHERE username=?`
	if err := db.QueryRowx(stmt, username).Scan(&count); err != nil {
		return count, err
	}
	return count, nil
}

func createUser(user *model.User) (int64, error) {
	stmt := `INSERT INTO users(
		name, created_at, modified_at
	) VALUES (?, ?, ?);`
	res, err := db.Exec(stmt, user.Name, user.CreatedAt, user.ModifiedAt)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func createUserLogin(ul *model.UserLogin) error {
	stmt := `INSERT INTO user_logins(
		user_id, username, password, salt, created_at, modified_at
	) VALUES (?, ?, ?, ?, ?, ?);`
	if _, err := db.Exec(stmt, ul.UserID, ul.Username, ul.Password,
		ul.Salt, ul.CreatedAt, ul.ModifiedAt); err != nil {
		return err
	}
	return nil
}
