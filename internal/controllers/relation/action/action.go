package action

import (
	"log"
	"tiktok/internal/services/login"
	"tiktok/internal/services/relation"
)

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("Relation.Action: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	if args.LoginID == args.UserID {
		reply.StatusCode = int32(StatusIllegalOperation)
		reply.StatusMsg = StatusIllegalOperation.msg()
		return reply
	}

	switch args.Type {
	case OpFollow:
		if err := relation.Follow(args.LoginID, args.UserID); err != nil {
			log.Println("Relation.Action: ", err)
			reply.StatusCode = int32(StatusFollowFailed)
			reply.StatusMsg = StatusFollowFailed.msg()
			return reply
		}
	case OpCancel:
		if err := relation.Unfollow(args.LoginID, args.UserID); err != nil {
			log.Println("Relation.Action: ", err)
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
	OpFollow = 1
	OpCancel = 2
)

type Request struct {
	LoginID int64
	UserID  int64
	Type    int64
}

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type status int32

const (
	StatusOK status = iota
	StatusIllegalOperation
	StatusFollowFailed
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
	case StatusFollowFailed:
		return "关注失败"
	case StatusCancelFailed:
		return "取消失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
