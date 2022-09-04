package user

import (
	"tiktok/internal/mapper"
	"tiktok/internal/middleware"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
)

func Login(args *in.Login) (*out.Login, error) {
	username, password := args.Username, args.Password
	// 参数校验
	if len(username) > maxLength || len(password) > maxLength {
		return &out.Login{
			StatusCode: int32(status.TooLong),
			StatusMsg:  status.TooLong.Message(),
		}, nil
	}
	// 按名取值
	user, err := mapper.LoginByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &out.Login{
			StatusCode: int32(status.UsernamePasswordNotMatch),
			StatusMsg:  status.UsernamePasswordNotMatch.Message(),
		}, nil
	}
	// 含盐加密
	encrypted := encrypt(password, user.Salt)
	// 校验密码
	if encrypted != user.Password {
		return &out.Login{
			StatusCode: int32(status.UsernamePasswordNotMatch),
			StatusMsg:  status.UsernamePasswordNotMatch.Message(),
		}, nil
	}
	// 签发token
	token, err := middleware.ReleaseToken(user.UserID)
	if err != nil {
		return nil, err
	}

	reply := &out.Login{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
		UserID:     user.UserID,
		Token:      token,
	}

	return reply, nil
}
