package user

import (
	"log"
	"tiktok/internal/services/model"
)

func User(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	user, err := userByID(args.UserID)
	logError(err)
	if user == nil {
		reply.StatusCode = int32(StatusNotExists)
		reply.StatusMsg = StatusNotExists.msg()
		return reply
	}

	fwc, err := countFollowsByID(args.UserID)
	logError(err)
	if err == nil {
		user.FollowCount = fwc
	}

	frc, err := countFollowersByID(args.UserID)
	logError(err)
	if err == nil {
		user.FollowerCount = frc
	}

	followed, err := isFollowByID(args.LoginID, args.UserID)
	logError(err)
	if err == nil {
		user.IsFollow = followed
	}

	reply.User = user

	return reply
}

type Request struct {
	LoginID int64
	UserID  int64
}

type Response struct {
	StatusCode int32       `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg,omitempty"` // 返回状态描述
	User       *model.User `json:"user"`                 // 用户信息
}

func logError(err error) {
	if err != nil {
		log.Println("user.User: ", err)
	}
}

const (
	StatusOK status = iota
	StatusNotExists
)

type status int32

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusNotExists:
		return "要找的用户不存在"
	default:
		return "未知错误"
	}
}
