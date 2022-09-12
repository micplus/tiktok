package user

import (
	"tiktok/internal/model"
	"tiktok/internal/pkg/database"

	"github.com/jmoiron/sqlx"
)

func ByID(id int64) (*model.User, error) {
	db := database.DB
	user := []model.User{}
	stmt := `SELECT * FROM users WHERE id=?;`
	if err := db.Select(&user, stmt, id); err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, nil
	}
	return &user[0], nil
}

func ByIDs(ids []int64) ([]model.User, error) {
	db := database.DB
	users := []model.User{}
	stmt := `SELECT * FROM users WHERE id IN (?);`
	query, args, err := sqlx.In(stmt, ids)
	if err != nil {
		return users, err
	}
	err = db.Select(&users, db.Rebind(query), args...)
	return users, err
}

func Insert(user *model.User) (int64, error) {
	db := database.DB
	stmt := `INSERT INTO users(
		name, created_at, modified_at
	) VALUES (?, ?, ?);`
	res, err := db.Exec(stmt, user.Name, user.CreatedAt, user.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
