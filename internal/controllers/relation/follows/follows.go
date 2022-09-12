package follows

import (
	"log"
	"sort"
	"tiktok/internal/model"
	"tiktok/internal/services/login"
	"tiktok/internal/services/relation"
	"tiktok/internal/services/user"
)

func Follows(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("Relation.Follows: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	ids, err := relation.FollowsByID(args.UserID)
	if err != nil {
		log.Println("Relation.Follows: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}
	if len(ids) == 0 {
		reply.UserList = []model.User{}
		return reply
	}

	users, err := user.ByIDs(ids)
	if err != nil {
		log.Println("Relation.Follows: ", err)
		reply.StatusCode = int32(StatusFailed)
		reply.StatusMsg = StatusFailed.msg()
		return reply
	}

	// 查自己关注列表，认为关注数<粉丝数
	f, err := relation.FollowsByID(args.LoginID)
	if err != nil {
		log.Println("Relation.Follows: ", err)
	}
	if err == nil {
		sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })
		for i := range users {
			t := sort.Search(len(f), func(j int) bool { return f[j] == users[i].ID })
			if t < len(f) {
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
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusFailed:
		return "查询失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
