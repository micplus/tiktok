package list

import (
	"log"
	"tiktok/internal/model"
	"tiktok/internal/services/comment"
	"tiktok/internal/services/login"
	"tiktok/internal/services/user"
)

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("Comment.List: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	comments, err := comment.ByVideoID(args.VideoID)
	if err != nil {
		log.Println("Comment.List: ", err)
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}

	// 设置每条评论对应的用户信息
	userIDs := make([]int64, len(comments))
	for i, c := range comments {
		userIDs[i] = c.UserID
	}
	users, err := user.ByIDs(userIDs)
	if err != nil {
		log.Println("Comment.List: ", err)
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}
	idToUser := make(map[int64]*model.User)
	for i, u := range users {
		idToUser[u.ID] = &users[i]
	}
	for i := range comments {
		comments[i].User = *idToUser[comments[i].UserID]
	}

	reply.CommentList = comments

	return reply
}

type Request struct {
	LoginID int64
	VideoID int64
}

type Response struct {
	StatusCode  int32           `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string          `json:"status_msg,omitempty"`   // 返回状态描述
	CommentList []model.Comment `json:"comment_list,omitempty"` // 评论列表
}

type status int32

const (
	StatusOK status = iota
	StatusListFailed
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusListFailed:
		return "获取评论列表失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
