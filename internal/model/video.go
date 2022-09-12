package model

// Video
type Video struct {
	ID            int64  `json:"id" db:"id"`               // 视频唯一标识
	Title         string `json:"title" db:"title"`         // 视频标题
	PlayURL       string `json:"play_url" db:"play_url"`   // 视频播放地址
	CoverURL      string `json:"cover_url" db:"cover_url"` // 视频封面地址
	UserID        int64  `json:"-" db:"user_id"`
	CommentCount  int64  `json:"comment_count" db:"-"`  // 视频的评论总数
	FavoriteCount int64  `json:"favorite_count" db:"-"` // 视频的点赞总数
	IsFavorite    bool   `json:"is_favorite" db:"-"`    // true-已点赞，false-未点赞
	CreatedAt     int64  `json:"-" db:"created_at"`     // 视频创建时间
	ModifiedAt    int64  `json:"-" db:"modified_at"`
	User          User   `json:"author" db:"user"` // 视频作者信息
}
