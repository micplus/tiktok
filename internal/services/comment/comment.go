package comment

import (
	"tiktok/internal/model"
	"tiktok/internal/pkg/database"

	"github.com/jmoiron/sqlx"
)

func Insert(c *model.Comment) (int64, error) {
	db := database.DB
	stmt := `INSERT INTO comments(
		content, create_date, created_at, modified_at, video_id, user_id
	) VALUES(?, ?, ?, ?, ?, ?);`
	res, err := db.Exec(stmt, c.Content, c.CreateDate,
		c.CreatedAt, c.ModifiedAt, c.VideoID, c.UserID)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

func CountByVideoID(id int64) (int64, error) {
	db := database.DB
	var count int64
	stmt := `SELECT COUNT(*) FROM comments WHERE video_id=?;`
	err := db.QueryRowx(stmt, id).Scan(&count)
	return count, err
}

func CountsByVideoIDs(ids []int64) (map[int64]int64, error) {
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

func ByVideoID(id int64) ([]model.Comment, error) {
	db := database.DB
	comments := []model.Comment{}
	stmt := `SELECT * FROM comments WHERE video_id=? ORDER BY modified_at DESC;`
	err := db.Select(&comments, stmt, id)
	return comments, err
}

func DeleteByID(id int64) error {
	db := database.DB
	stmt := `DELETE FROM comments WHERE id=?;`
	_, err := db.Exec(stmt, id)
	return err
}

type commentCount struct {
	videoID int64 `db:"video_id"`
	count   int64 `db:"count"`
}
