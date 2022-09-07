package model

// 视频作者信息
//
// User
type User struct {
	ID            int64  `json:"id" db:"id"`                                   // 用户id
	Name          string `json:"name" db:"name"`                               // 用户名称
	FollowCount   int64  `json:"follow_count,omitempty" db:"follow_count"`     // 关注总数
	FollowerCount int64  `json:"follower_count,omitempty" db:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow" db:"is_follow"`                     // true-已关注，false-未关注
	CreatedAt     int64  `json:"-" db:"created_at"`
	ModifiedAt    int64  `json:"-" db:"modified_at"`
}

type UserLogin struct {
	ID         int64  `db:"id"`
	UserID     int64  `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Salt       string `db:"salt"`
	CreatedAt  int64  `db:"created_at"`
	ModifiedAt int64  `db:"modified_at"`
}
