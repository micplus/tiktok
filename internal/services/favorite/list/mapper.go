package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"

	"github.com/jmoiron/sqlx"
)

func favoriteIDsByUserID(id int64) ([]int64, error) {
	db := database.DB
	stmt := `SELECT DISTINCT video_id FROM user_favorites WHERE user_id=?;`
	rows, err := db.Queryx(stmt, id)
	if err != nil {
		return []int64{}, err
	}
	fids := []int64{}
	for rows.Next() {
		var fid int64
		rows.Scan(&fid)
		fids = append(fids, fid)
	}
	return fids, err
}

func videosByIDs(ids []int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	stmt := `SELECT 
		videos.*, 
		users.id 'user.id',
		users.name 'user.name'
	FROM
		videos, users
	WHERE 
		videos.id IN (?) AND
		videos.user_id=users.id 
	ORDER BY videos.created_at DESC;`
	query, args, err := sqlx.In(stmt, ids)
	if err != nil {
		return videos, err
	}
	err = db.Select(&videos, db.Rebind(query), args...)
	return videos, err
}

func isFavoritesOfUserID(videoIDs []int64, userID int64) (map[int64]struct{}, error) {
	db := database.DB
	f := make(map[int64]struct{})
	stmt := `SELECT video_id FROM user_favorites
		WHERE user_id=? AND video_id IN (?);`
	query, args, err := sqlx.In(stmt, userID, videoIDs)
	if err != nil {
		return f, err
	}
	rows, err := db.Queryx(db.Rebind(query), args...)
	if err != nil {
		return f, err
	}
	ids := []int64{}
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		ids = append(ids, id)
	}
	for _, id := range ids {
		f[id] = struct{}{}
	}
	return f, nil
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
