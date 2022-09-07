package list

import (
	"log"
	"tiktok/internal/services/model"
)

func List(args *Request) (*Response, error) {
	user, err := userByID(args.UserID)
	if err != nil {
		log.Println("publish.list.List: ", err)
		return nil, err
	}

	videos, err := videosByUserID(args.UserID)
	if err != nil {
		log.Println("publish.list.List: ", err)
		return nil, err
	}
	// 设置User，并取出所有ID
	ids := make([]int64, len(videos))
	for i := range videos {
		videos[i].User = *user
		ids[i] = videos[i].ID
	}

	// 设置点赞数
	favoriteCount, err := favoriteCountsByVideoIDs(ids)
	if err != nil {
		log.Println(err)
	}
	if err == nil {
		for i := range videos {
			videos[i].FavoriteCount = favoriteCount[videos[i].ID]
		}
	}

	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
		VideoList:  videos,
	}
	return reply, nil
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
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	default:
		return "未知错误"
	}
}
