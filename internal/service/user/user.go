package user

import (
	"tiktok/internal/mapper"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
)

func User(args *in.User) (*out.User, error) {
	// 鉴权由中间件代行
	user, err := mapper.UserByID(args.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return &out.User{
			StatusCode: int32(status.UserNotExists),
			StatusMsg:  status.UserNotExists.Message(),
		}, nil
	}

	reply := &out.User{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
		User:       user,
	}
	return reply, nil
}
