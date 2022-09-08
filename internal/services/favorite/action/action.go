package action

import "log"

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	count, err := countUserFavorites(args.LoginID, args.VideoID)
	if err != nil {
		log.Println("favorite.action.Action: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	// 重复请求直接跳过
	if count != 0 && args.Type == OpFavorite {
		return reply
	}
	if count == 0 && args.Type == OpCancel {
		return reply
	}

	switch args.Type {
	case OpFavorite:
		err := createUserFavorite(args.LoginID, args.VideoID)
		if err != nil {
			log.Println("favorite.action.Action: ", err)
			reply.StatusCode = int32(StatusFavorateFailed)
			reply.StatusMsg = StatusFavorateFailed.msg()
			return reply
		}
	case OpCancel:
		err := deleteUserFavorite(args.LoginID, args.VideoID)
		if err != nil {
			log.Println("favorite.action.Action: ", err)
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
	default:
		return "未知错误"
	}
}
