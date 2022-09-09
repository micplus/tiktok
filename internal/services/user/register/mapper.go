package register

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

func countUsername(username string) (int64, error) {
	db := database.DB
	var count int64
	stmt := `SELECT COUNT(*) FROM user_logins WHERE username=?;`
	err := db.QueryRowx(stmt, username).Scan(&count)
	return count, err
}

func createUser(user *model.User) (int64, error) {
	db := database.DB
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
	db := database.DB
	stmt := `INSERT INTO user_logins(
		user_id, username, password, salt, created_at, modified_at
	) VALUES (?, ?, ?, ?, ?, ?);`
	_, err := db.Exec(stmt, ul.UserID, ul.Username, ul.Password,
		ul.Salt, ul.CreatedAt, ul.ModifiedAt)
	return err
}
