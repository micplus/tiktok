package action

import (
	"log"
	"tiktok/internal/services/favorite"
	"tiktok/internal/services/login"
)

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("Favorite.Action: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	switch args.Type {
	case OpFavorite:
		if err := favorite.Favorite(args.LoginID, args.VideoID); err != nil {
			log.Println("Favorite.Action: ", err)
			reply.StatusCode = int32(StatusFavorateFailed)
			reply.StatusMsg = StatusFavorateFailed.msg()
			return reply
		}
	case OpCancel:
		if err := favorite.Unfavorite(args.LoginID, args.VideoID); err != nil {
			log.Println("Favorite.Action: ", err)
			reply.StatusCode = int32(StatusCancelFailed)
			reply.StatusMsg = StatusCancelFailed.msg()
			return reply
		}
	default:
		reply.StatusCode = int32(StatusIllegalOperation)
		reply.StatusMsg = StatusIllegalOperation.msg()
		return reply
	}

	return reply
}

const (
	OpFavorite = 1
	OpCancel   = 2
)

type Request struct {
	LoginID int64
	VideoID int64
	Type    int64 // 1点赞，2取消
}

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type status int32

const (
	StatusOK status = iota
	StatusIllegalOperation
	StatusFavorateFailed
	StatusCancelFailed
	StatusFailed
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusIllegalOperation:
		return "不合法的操作"
	case StatusFavorateFailed:
		return "点赞失败"
	case StatusCancelFailed:
		return "取消失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
