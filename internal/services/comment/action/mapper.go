package action

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func createComment(c *model.Comment) error {
	stmt := `INSERT INTO comments(
		content, create_date, created_at, modified_at, video_id, user_id
	) VALUES(?, ?, ?, ?, ?, ?);`
	_, err := db.Exec(stmt, c.Content, c.CreateDate,
		c.CreatedAt, c.ModifiedAt, c.VideoID, c.UserID)
	return err
}

func userByID(id int64) (*model.User, error) {
	user := new(model.User)
	stmt := `SELECT * FROM users WHERE id=?;`
	err := db.Select(user, stmt, id)
	return user, err
}

func deleteCommentByID(id int64) error {
	stmt := `DELETE FROM comments WHERE id=?;`
	_, err := db.Exec(stmt, id)
	return err
}
