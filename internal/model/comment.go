package model

// 评论成功返回评论内容，不需要重新拉取整个列表
//
// Comment
type Comment struct {
	Content    string `json:"content"`                            // 评论内容
	CreateDate string `json:"create_date" gorm:"-"`               // 评论发布日期，格式 mm-dd
	ID         int64  `json:"id" gorm:"primaryKey;autoIncrement"` // 评论id
	User       User   `json:"user"`                               // 评论用户信息
	UserID     int64  `json:"-"`
	VideoID    int64  `json:"-"`
	CreatedAt  int64  `json:"-" gorm:"autoCreateTime:milli"`
}
