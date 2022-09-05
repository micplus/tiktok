package model

// 视频作者信息
//
// User
type User struct {
	FollowCount    int64      `json:"follow_count,omitempty"`              // 关注总数
	FollowerCount  int64      `json:"follower_count,omitempty"`            // 粉丝总数
	ID             int64      `json:"id"  gorm:"primaryKey;autoIncrement"` // 用户id
	IsFollow       bool       `json:"is_follow" gorm:"-"`                  // true-已关注，false-未关注
	Name           string     `json:"name"`                                // 用户名称
	Login          *UserLogin `json:"-"`
	Published      []*Video   `json:"-"`
	FavoriteVideos []*Video   `json:"-" gorm:"many2many:user_videos"`
	Follows        []*User    `json:"-" gorm:"many2many:user_follows"`
	Comments       []*Comment `json:"-"`
}

type UserLogin struct {
	ID       int64 `gorm:"primary_key;autoIncrement"`
	UserID   int64
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Salt     string `gorm:"not null"`
}
