package video

import (
	"tiktok/internal/model"
	"tiktok/internal/pkg/database"

	"github.com/jmoiron/sqlx"
)

const limit = 30

func Before(now int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	query := `SELECT * FROM videos 
	WHERE modified_at < ?
	ORDER BY modified_at DESC
	LIMIT ?;`
	err := db.Select(&videos, query, now, limit)
	return videos, err
}

func ByIDs(ids []int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	stmt := `SELECT * FROM videos WHERE id IN (?) ORDER BY modified_at DESC;`
	query, args, err := sqlx.In(stmt, ids)
	if err != nil {
		return videos, err
	}
	err = db.Select(&videos, db.Rebind(query), args...)
	return videos, err
}

func VideosByUserID(id int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	stmt := `SELECT * FROM videos WHERE user_id=? ORDER BY modified_at DESC;`
	err := db.Select(&videos, stmt, id)
	return videos, err
}

func Insert(v *model.Video) (int64, error) {
	db := database.DB
	stmt := `INSERT INTO videos(
		title, play_url, cover_url, user_id, created_at, modified_at
	) VALUES(?, ?, ?, ?, ?, ?);`
	res, err := db.Exec(stmt, v.Title, v.PlayURL, v.CoverURL,
		v.UserID, v.CreatedAt, v.ModifiedAt)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}
