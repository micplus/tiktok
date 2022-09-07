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

// 并返回创建的ID
func createUser(user *model.User) (int64, error) {
	var id int64
	tx, err := db.Beginx()
	if err != nil {
		return id, err
	}
	stmt := `INSERT INTO users(
		name, created_at, modified_at
	) VALUES (?, ?, ?);`
	if _, err = tx.Exec(stmt, user.Name, user.CreatedAt, user.ModifiedAt); err != nil {
		return id, err
	}
	if err = tx.Select(&id, `SELECT LAST_INSERT_ID();`); err != nil {
		return id, err
	}
	if err = tx.Commit(); err != nil {
		return id, err
	}
	return id, nil
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
