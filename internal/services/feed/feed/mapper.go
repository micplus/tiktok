package feed

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"

	"github.com/jmoiron/sqlx"
)

const limit = 30

// 基本功能取给定时间之前的视频
func videosBeforeTime(now int64) ([]model.Video, error) {
	db := database.DB
	videos := []model.Video{}
	stmt := `SELECT 
		videos.*, 
		users.id 'user.id',
		users.name 'user.name'
	FROM videos 
	JOIN users ON videos.user_id=users.id
	WHERE videos.created_at < ?
	ORDER BY videos.created_at DESC
	LIMIT ?;`
	err := db.Select(&videos, stmt, now, limit)
	return videos, err
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

// 查询给定IDs评论数到id: count
func commentCountsByVideoIDs(ids []int64) (map[int64]int64, error) {
	db := database.DB
	count := make(map[int64]int64)
	counts := []commentCount{}
	stmt := `SELECT video_id, COUNT(*) 'count'
	FROM comments
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
		var cc commentCount
		rows.Scan(&cc.videoID, &cc.count)
		counts = append(counts, cc)
	}

	for _, cc := range counts {
		count[cc.videoID] = cc.count
	}
	return count, nil
}

func favoritesByUserID(id int64) ([]int64, error) {
	db := database.DB
	favorites := []int64{}
	stmt := `SELECT DISTINCT video_id FROM user_favorites WHERE user_id=?`
	rows, err := db.Queryx(stmt, id)
	if err != nil {
		return favorites, err
	}
	for rows.Next() {
		var favoriteID int64
		rows.Scan(&favoriteID)
		favorites = append(favorites, favoriteID)
	}

	return favorites, nil
}

type favoriteCount struct {
	videoID int64 `db:"video_id"`
	count   int64 `db:"count"`
}

type commentCount struct {
	videoID int64 `db:"video_id"`
	count   int64 `db:"count"`
}
