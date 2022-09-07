package login

import (
	"log"
	"tiktok/internal/pkg/auth"
)

const maxLength = 32 // same as register

func Login(args *Request) (*Response, error) {
	username, password := args.Username, args.Password
	// 参数校验
	if len(username) > maxLength || len(password) > maxLength {
		return &Response{
			StatusCode: int32(statusTooLong),
			StatusMsg:  statusTooLong.msg(),
		}, nil
	}

	// 按名取值
	user, err := loginByUsername(username)
	if err != nil {
		log.Println("user.login.Login: ", err)
		return nil, err
	}
	if user == nil {
		return &Response{
			StatusCode: int32(statusNotMatch),
			StatusMsg:  statusNotMatch.msg(),
		}, nil
	}

	// 含盐加密
	encrypted := auth.Encrypt(password, user.Salt)
	// 校验密码
	if encrypted != user.Password {
		return &Response{
			StatusCode: int32(statusNotMatch),
			StatusMsg:  statusNotMatch.msg(),
		}, nil
	}
	// 签发token
	token, err := auth.ReleaseToken(user.UserID)
	if err != nil {
		return nil, err
	}

	reply := &Response{
		StatusCode: int32(statusOK),
		StatusMsg:  statusOK.msg(),
		UserID:     user.UserID,
		Token:      token,
	}

	return reply, nil
}

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
	statusOK status = iota
	statusTooLong
	statusNotMatch
)

func (s status) msg() string {
	switch s {
	case statusOK:
		return "OK"
	case statusTooLong:
		return "用户名、密码不能超过32个字符"
	case statusNotMatch:
		return "用户名或密码错误"
	default:
		return "未知错误"
	}
}
