package model

// Comment
type Comment struct {
	ID         int64  `json:"id" db:"id"`                   // 评论id
	Content    string `json:"content" db:"content"`         // 评论内容
	CreateDate string `json:"create_date" db:"create_date"` // 评论发布日期，格式 mm-dd
	UserID     int64  `json:"-" db:"user_id"`
	VideoID    int64  `json:"-" db:"video_id"`
	CreatedAt  int64  `json:"-" db:"create_at"`
	ModifiedAt int64  `json:"-" db:"modified_at"`
	User       User   `json:"user" db:"user"` // 评论用户信息
}
