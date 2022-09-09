package action

import "tiktok/internal/pkg/database"

func countUserFavorites(userID, videoID int64) (int64, error) {
	db := database.DB
	var count int64
	stmt := `SELECT COUNT(*) FROM user_favorites WHERE user_id=? AND video_id=?;`
	err := db.QueryRowx(stmt, userID, videoID).Scan(&count)
	return count, err
}

func createUserFavorite(userID, videoID int64) error {
	db := database.DB
	stmt := `INSERT INTO user_favorites(user_id, video_id) VALUES(?, ?);`
	_, err := db.Exec(stmt, userID, videoID)
	return err
}

func deleteUserFavorite(userID, videoID int64) error {
	db := database.DB
	stmt := `DELETE FROM user_favorites WHERE user_id=? AND video_id=?;`
	_, err := db.Exec(stmt, userID, videoID)
	return err
}
