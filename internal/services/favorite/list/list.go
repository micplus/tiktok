package list

import (
	"log"
	"tiktok/internal/services/model"
)

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// 查目标用户点赞列表
	ids, err := favoriteIDsByUserID(args.UserID)
	if err != nil {
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}
	if len(ids) == 0 {
		reply.VideoList = []model.Video{}
		return reply
	}
	videos, err := videosByIDs(ids)
	if err != nil {
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}

	// 查自己点赞列表，更新红心
	f, err := isFavoritesOfUserID(ids, args.LoginID)
	if err != nil {
		log.Println("favorite.action.List: ", err)
	}
	if err == nil {
		for i := range videos {
			if _, ok := f[videos[i].ID]; ok {
				videos[i].IsFavorite = true
			}
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
	VideoList  []model.Video `json:"video_list,omitempty"` // 用户点赞视频列表
}

type status int32

const (
	StatusOK status = iota
	StatusListFailed
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusListFailed:
		return "获取点赞列表失败"
	default:
		return "未知错误"
	}
}
