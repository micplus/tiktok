package register

import (
	"log"
	"tiktok/internal/model"
	"tiktok/internal/pkg/auth"
	"tiktok/internal/services/login"
	"tiktok/internal/services/user"

	"time"
)

const maxLength = 32 // same as login

func Register(args *Request) *Response {
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
	// 检查重名
	cnt, err := login.CountByUsername(username)
	if err != nil {
		log.Println("User.Register: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	if cnt > 0 {
		reply.StatusCode = int32(StatusUsernameExists)
		reply.StatusMsg = StatusUsernameExists.msg()
		return reply
	}
	// 含盐加密
	salt := auth.MakeSalt()
	encrypted := auth.Encrypt(password, salt)

	// 创建用户
	now := time.Now().UnixMilli()
	u := &model.User{
		Name:       username,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	id, err := user.Insert(u)
	if err != nil {
		log.Println("User.Register: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	// 创建登录信息
	l := &model.UserLogin{
		UserID:     id,
		Username:   username,
		Password:   encrypted,
		Salt:       salt,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	_, err = login.Insert(l)
	if err != nil {
		log.Println("User.Register: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	// 生成token
	token, err := auth.ReleaseToken(u.ID)
	if err != nil {
		log.Println("User.Register: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	reply.UserID = id
	reply.Token = token

	// 缓存登录状态，流程外
	// go setCache(reply.UserID)

	return reply
}

// func setCache(id int64) {
// 	if err := login.SetCache(id); err != nil {
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
	StatusUsernameExists
	StatusFailed
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusTooLong:
		return "用户名、密码不能超过32个字符"
	case StatusUsernameExists:
		return "用户名已存在"
	case StatusFailed:
		return "注册失败"
	default:
		return "未知错误"
	}
}
