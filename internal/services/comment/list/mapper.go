package list

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

func commentsByVideoID(id int64) ([]model.Comment, error) {
	db := database.DB
	comments := []model.Comment{}
	stmt := `SELECT
		comments.*,
		users.id 'user.id',
		users.name 'user.name'
	FROM comments
	JOIN users ON comments.user_id=users.id
	WHERE comments.video_id=?
	ORDER BY comments.modified_at DESC;`
	err := db.Select(&comments, stmt, id)
	return comments, err
}
