package list

import (
	"log"
	"tiktok/internal/services/model"
)

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ids, err := followerIDsByUserID(args.UserID)
	if err != nil {
		log.Println("relation.follower.list.List: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	if len(ids) == 0 {
		reply.UserList = []model.User{}
		return reply
	}

	users, err := usersByIDs(ids)
	if err != nil {
		log.Println("relation.follower.list.List: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	// 查自己关注列表，更新红心
	f, err := isFollowersOfUserID(ids, args.LoginID)
	if err != nil {
		log.Println("relation.follower.list.List: ", err)
	}
	if err == nil {
		for i := range users {
			if _, ok := f[users[i].ID]; ok {
				users[i].IsFollow = true
			}
		}
	}

	reply.UserList = users

	return reply
}

type Request struct {
	LoginID int64
	UserID  int64
}

type Response struct {
	StatusCode int32        `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg,omitempty"` // 返回状态描述
	UserList   []model.User `json:"user_list,omitempty"`
}

type status int32

const (
	StatusOK status = iota
	StatusFailed
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusFailed:
		return "查询失败"
	default:
		return "未知错误"
	}
}
