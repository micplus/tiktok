package action

import (
	"tiktok/internal/pkg/database"
	"tiktok/internal/services/model"
)

var db = database.DB

func createVideo(v *model.Video) error {
	stmt := `INSERT INTO videos(
		title, play_url, cover_url, user_id, created_at, modified_at
	) VALUES(?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(stmt, v.Title, v.PlayURL, v.CoverURL,
		v.UserID, v.CreatedAt, v.ModifiedAt)
	return err
}
