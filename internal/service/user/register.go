package user

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"tiktok/internal/mapper"
	"tiktok/internal/middleware"
	"tiktok/internal/model"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
	"time"
)

const (
	maxLength  = 32
	saltLength = 4
)

func Register(args *in.Register) (*out.Register, error) {
	username, password := args.Username, args.Password
	// 参数校验
	if len(username) > maxLength || len(password) > maxLength {
		return &out.Register{
			StatusCode: int32(status.TooLong),
		}, nil
	}
	// 检查重名
	ok, err := mapper.CheckUniqueUsername(username)
	if err != nil {
		return nil, err
	}
	if !ok {
		return &out.Register{
			StatusCode: int32(status.UsernameExists),
		}, nil
	}
	// 生成盐值
	salt := makeSalt()
	// 含盐加密
	encrypted := encrypt(password, salt)
	// 存入数据库

	login := &model.UserLogin{
		Username: username,
		Password: encrypted,
		Salt:     salt,
	}
	user := &model.User{
		Login: login,
		Name:  username,
	}

	if err := mapper.CreateUser(user); err != nil {
		return nil, err
	}

	token, err := middleware.ReleaseToken(user.ID)
	if err != nil {
		return nil, err
	}

	reply := &out.Register{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
		UserID:     user.ID,
		Token:      token,
	}

	return reply, nil
}

func makeSalt() string {
	now := time.Now().Nanosecond() - time.Now().Second()
	nowBuf := []byte(strconv.FormatInt(int64(now), 10))

	hash := md5.New()
	hash.Write(nowBuf)

	salt := fmt.Sprintf("%x", hash.Sum(nil))
	salt = salt[:saltLength]
	return salt
}

func encrypt(password, salt string) string {
	buf := []byte(password + salt)
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
