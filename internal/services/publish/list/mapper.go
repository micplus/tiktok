package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"

	"github.com/jmoiron/sqlx"
)

func videosByUserID(id int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	stmt := `SELECT * FROM videos WHERE user_id=? ORDER BY modified_at DESC;`
	err := db.Select(&videos, stmt, id)
	return videos, err
}

func userByID(id int64) (*model.User, error) {
	db := database.DB
	users := []model.User{}
	stmt := `SELECT * FROM users WHERE id=?;`
	err := db.Select(&users, stmt, id)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}

// 查询给定IDs点赞数量到id: count
func favoriteCountsByVideoIDs(ids []int64) (map[int64]int64, error) {
	db := database.DB
	count := make(map[int64]int64)
	counts := []favoriteCount{}
	stmt := `SELECT video_id, COUNT(*) 'count' 
	FROM user_favorites
	WHERE video_id IN (?)
	GROUP BY video_id;`
	query, args, err := sqlx.In(stmt, ids)
	if err != nil {
		return count, err
	}
	rows, err := db.Queryx(db.Rebind(query), args...)
	if err != nil {
		return count, err
	}
	for rows.Next() {
		var fc favoriteCount
		rows.Scan(&fc.videoID, &fc.count)
		counts = append(counts, fc)
	}

	for _, fc := range counts {
		count[fc.videoID] = fc.count
	}
	return count, nil
}

type favoriteCount struct {
	videoID int64 `db:"video_id"`
	count   int64 `db:"count"`
}
