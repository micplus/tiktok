package feed

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

const limit = 30

var db = database.DB

// 基本功能取给定时间之前的视频
func videosBeforeTime(now int64) ([]model.Video, error) {
	videos := []model.Video{}
	stmt := `SELECT 
		videos.*, 
		users.id 'user.id',
		users.name 'user.name',
	FROM videos 
	JOIN users ON videos.user_id=users.id
	WHERE videos.create_at < ?
	ORDER BY videos.created_at DESC
	LIMIT ?;`
	err := db.Select(&videos, stmt, now, limit)
	return videos, err
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

// 查询给定IDs评论数到id: count
func commentCountsByVideoIDs(ids []int64) (map[int64]int64, error) {
	count := make(map[int64]int64)

	counts := []commentCount{}
	stmt := `SELECT video_id, COUNT(*) 'count'
	FROM comments
	WHERE video_id IN (?)
	GROUP BY video_id;`
	if err := db.Select(&counts, stmt, ids); err != nil {
		return count, err
	}

	for _, cc := range counts {
		count[cc.videoID] = cc.count
	}
	return count, nil
}

func favoritesByUserID(id int64) ([]int64, error) {
	favorites := []int64{}
	stmt := `SELECT DISTINCT video_id FROM user_favorites WHERE user_id=?`
	if err := db.Select(&favorites, stmt, id); err != nil {
		return favorites, err
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
