package action

import "tiktok/internal/pkg/database"

var db = database.DB

func countUserFollows(fromID, toID int64) (int64, error) {
	var count int64
	stmt := `SELECT COUNT(*) FROM user_follows WHERE user_ID=? AND follow_id=?;`
	err := db.QueryRowx(stmt, fromID, toID).Scan(&count)
	return count, err
}

func createUserFollow(fromID, toID int64) error {
	stmt := `INSERT INTO user_follows(user_id, follow_id) VALUES(?, ?);`
	_, err := db.Exec(stmt, fromID, toID)
	return err
}

func deleteUserFollow(fromID, toID int64) error {
	stmt := `DELETE FROM user_follows WHERE user_id=? AND follow_id=?;`
	_, err := db.Exec(stmt, fromID, toID)
	return err
}
