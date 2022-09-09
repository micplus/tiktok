package action

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

func createComment(c *model.Comment) error {
	db := database.DB
	stmt := `INSERT INTO comments(
		content, create_date, created_at, modified_at, video_id, user_id
	) VALUES(?, ?, ?, ?, ?, ?);`
	_, err := db.Exec(stmt, c.Content, c.CreateDate,
		c.CreatedAt, c.ModifiedAt, c.VideoID, c.UserID)
	return err
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
	return &users[0], err
}

func deleteCommentByID(id int64) error {
	db := database.DB
	stmt := `DELETE FROM comments WHERE id=?;`
	_, err := db.Exec(stmt, id)
	return err
}
