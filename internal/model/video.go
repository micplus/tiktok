package model

// Video
type Video struct {
	User          User       `json:"author"` // 视频作者信息
	UserID        int64      `json:"-"`
	CommentCount  int64      `json:"comment_count" gorm:"default:0"`      // 视频的评论总数
	CoverURL      string     `json:"cover_url" gorm:"not null"`           // 视频封面地址
	FavoriteCount int64      `json:"favorite_count" gorm:"default:0"`     // 视频的点赞总数
	ID            int64      `json:"id" gorm:"primaryKey;autoIncrement"`  // 视频唯一标识
	IsFavorite    bool       `json:"is_favorite"`                         // true-已点赞，false-未点赞
	PlayURL       string     `json:"play_url" gorm:"not null"`            // 视频播放地址
	Title         string     `json:"title" gorm:"not null"`               // 视频标题
	CreatedAt     int64      `json:"-" gorm:"autoCreateTime:milli;index"` // 视频创建时间
	Comments      []*Comment `json:"-"`
	FavoriteUsers []*User    `json:"-" gorm:"many2many:user_videos"`
}
