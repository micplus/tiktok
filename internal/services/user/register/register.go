package register

import (
	"log"
	"tiktok/internal/pkg/auth"
	"tiktok/internal/services/model"

	"time"
)

const (
	maxLength = 32
)

func Register(args *Request) (*Response, error) {
	username, password := args.Username, args.Password
	// 参数校验
	if len(username) > maxLength || len(password) > maxLength {
		return &Response{
			StatusCode: int32(statusTooLong),
			StatusMsg:  statusTooLong.msg(),
		}, nil
	}
	// 检查重名
	cnt, err := countUsername(username)
	if err != nil {
		log.Println("user.register.Register: ", err)
		return nil, err
	}
	if cnt > 0 {
		return &Response{
			StatusCode: int32(statusUsernameExists),
			StatusMsg:  statusUsernameExists.msg(),
		}, nil
	}
	// 含盐加密
	salt := auth.MakeSalt()
	encrypted := auth.Encrypt(password, salt)

	// 创建用户
	now := time.Now().UnixMilli()
	user := &model.User{
		Name:       username,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	id, err := createUser(user)
	if err != nil {
		return nil, err
	}

	// 创建登录信息
	login := &model.UserLogin{
		UserID:     id,
		Username:   username,
		Password:   encrypted,
		Salt:       salt,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	err = createUserLogin(login)
	if err != nil {
		return nil, err
	}

	// 生成token
	token, err := auth.ReleaseToken(user.ID)
	if err != nil {
		return nil, err
	}

	// 返回结果
	reply := &Response{
		StatusCode: int32(statusOK),
		StatusMsg:  statusOK.msg(),
		UserID:     id,
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
	statusUsernameExists
)

func (s status) msg() string {
	switch s {
	case statusOK:
		return "OK"
	case statusTooLong:
		return "用户名、密码不能超过32个字符"
	case statusUsernameExists:
		return "用户名已存在"
	default:
		return "未知错误"
	}
}
