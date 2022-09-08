package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func favoriteIDsByUserID(id int64) ([]int64, error) {
	fids := []int64{}
	stmt := `SELECT DISTINCT video_id FROM user_favorites WHERE user_id=?;`
	err := db.Select(&fids, stmt, id)
	return fids, err
}

func videosByIDs(ids []int64) ([]model.Video, error) {
	videos := []model.Video{}
	stmt := `SELECT 
		videos.*, 
		users.id 'user.id',
		users.name 'user.name',
	FROM
		videos, users
	WHERE 
		videos.id IN (?) AND
		videos.user_id=users.id 
	ORDER BY videos.created_at DESC;`
	err := db.Select(&videos, stmt, ids)
	return videos, err
}

func isFavoritesOfUserID(videoIDs []int64, userID int64) (map[int64]struct{}, error) {
	f := make(map[int64]struct{})
	ids := []int64{}
	stmt := `SELECT video_id FROM user_favorites
		WHERE user_id=? AND video_id IN (?);`
	if err := db.Select(&ids, stmt, videoIDs); err != nil {
		return f, err
	}
	for _, id := range ids {
		f[id] = struct{}{}
	}
	return f, nil
}
