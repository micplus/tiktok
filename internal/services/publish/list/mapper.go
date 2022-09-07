package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func videosByUserID(id int64) ([]model.Video, error) {
	videos := []model.Video{}
	stmt := `SELECT * FROM videos WHERE user_id=?`
	if err := db.Select(videos, stmt, id); err != nil {
		return videos, err
	}
	return videos, nil
}

func userByID(id int64) (*model.User, error) {
	user := new(model.User)
	stmt := `SELECT * FROM users WHERE id=?`
	if err := db.Select(user, stmt, id); err != nil {
		return nil, err
	}
	return user, nil
}

// 查询给定IDs点赞数量到id: count
func favoriteCountsByVideoIDs(ids []int64) (map[int64]int64, error) {
	count := make(map[int64]int64)

	counts := []favoriteCount{}
	stmt := `SELECT video_id, COUNT(*) 'count' 
	FROM user_favorites
	WHERE video_id IN (?)
	GROUP BY video_id;`
	if err := db.Select(&counts, stmt, ids); err != nil {
		return count, err
	}

	for _, fc := range counts {
		count[fc.videoID] = fc.count
	}
	return count, nil
}

type favoriteCount struct {
	videoID int64 `db:"user_id"`
	count   int64 `db:"count"`
}
