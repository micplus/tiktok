package login

import (
	"log"
	"tiktok/internal/pkg/auth"
	"tiktok/internal/services/login"
)

const maxLength = 32 // same as register

func Login(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	username, password := args.Username, args.Password
	// 参数校验
	if len(username) > maxLength || len(password) > maxLength {
		reply.StatusCode = int32(StatusTooLong)
		reply.StatusMsg = StatusTooLong.msg()
		return reply
	}

	// 按名取值
	user, err := login.ByUsername(username)
	if err != nil {
		log.Println("user.login.Login: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	if user == nil {
		reply.StatusCode = int32(StatusNotMatch)
		reply.StatusMsg = StatusNotMatch.msg()
		return reply
	}

	// 含盐加密
	encrypted := auth.Encrypt(password, user.Salt)
	// 校验密码
	if encrypted != user.Password {
		reply.StatusCode = int32(StatusNotMatch)
		reply.StatusMsg = StatusNotMatch.msg()
		return reply
	}
	// 签发token
	token, err := auth.ReleaseToken(user.UserID)
	if err != nil {
		log.Println("user.login.Login: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	reply.UserID = user.UserID
	reply.Token = token

	// 缓存登陆状态，不影响流程
	// if err := login.SetCache(reply.UserID); err != nil {
	// 	log.Println("User.Register: ", err)
	// }
	//go setCache(reply.UserID)

	return reply
}

// func setCache(id int64) {
// 	if err := login.RefreshCache(id); err != nil {
// 		log.Println("User.Register: ", err)
// 	}
// }

type Request struct {
	Username string
	Password string
}

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	Token      string `json:"token"`                // 用户鉴权token
	UserID     int64  `json:"user_id"`              // 用户id
}

type status int32

const (
	StatusOK status = iota
	StatusTooLong
	StatusNotMatch
	StatusFailed
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusTooLong:
		return "用户名、密码不能超过32个字符"
	case StatusNotMatch:
		return "用户名或密码错误"
	default:
		return "未知错误"
	}
}
