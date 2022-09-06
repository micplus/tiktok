package request

type FavoriteAction struct {
	LoginID int64
	VideoID int64
	Type    int64 // 1点赞，2取消
}

type FavoriteList struct {
	LoginID int64
	UserID  int64
}
