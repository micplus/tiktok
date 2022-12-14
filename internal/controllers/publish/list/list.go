package list

import (
	"log"
	"tiktok/internal/model"
	"tiktok/internal/services/favorite"
	"tiktok/internal/services/user"
	"tiktok/internal/services/video"
)

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// ok, err := login.CheckCache(args.LoginID)
	// if err != nil || !ok {
	// 	log.Println("Publish.List: ", err)
	// 	reply.StatusCode = int32(StatusTokenExpired)
	// 	reply.StatusMsg = StatusTokenExpired.msg()
	// 	return reply
	// }

	u, err := user.ByID(args.UserID)
	if err != nil {
		log.Println("publish.list.List: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	videos, err := video.VideosByUserID(args.UserID)
	if err != nil {
		log.Println("publish.list.List: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	// 设置User，并取出所有ID
	ids := make([]int64, len(videos))
	for i := range videos {
		videos[i].User = *u
		ids[i] = videos[i].ID
	}

	// 设置点赞数
	favoriteCount, err := favorite.CountFavoritedsByVideoIDs(ids)
	if err != nil {
		log.Println("publish.list.List: ", err)
	}
	if err == nil {
		for i := range videos {
			videos[i].FavoriteCount = favoriteCount[i]
		}
	}

	reply.VideoList = videos

	return reply
}

type Request struct {
	LoginID int64
	UserID  int64
}

type Response struct {
	StatusCode int32         `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []model.Video `json:"video_list,omitempty"` // 用户发布的视频列表
}

type status int32

const (
	StatusOK status = iota
	StatusFailed
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusFailed:
		return "查询失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
