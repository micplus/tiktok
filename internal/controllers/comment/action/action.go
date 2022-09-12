package action

import (
	"fmt"
	"log"
	"sync"
	"tiktok/internal/model"
	"tiktok/internal/services/comment"
	"tiktok/internal/services/login"
	"tiktok/internal/services/user"
	"time"
)

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("Comment.Action: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	if args.Type == OpComment && len(args.CommentText) == 0 {
		reply.StatusCode = int32(StatusNoText)
		reply.StatusMsg = StatusNoText.msg()
		return reply
	}

	switch args.Type {
	case OpComment:
		u := new(model.User)
		var wg sync.WaitGroup
		wg.Add(1)
		go userByID(args.LoginID, u, &wg)

		now := time.Now()
		month, day := now.Month(), now.Day()
		date := fmt.Sprintf("%02d-%02d", month, day)

		c := model.Comment{
			Content:    args.CommentText,
			CreateDate: date,
			UserID:     args.LoginID,
			VideoID:    args.VideoID,
			CreatedAt:  now.UnixMilli(),
			ModifiedAt: now.UnixMilli(),
		}

		_, err := comment.Insert(&c)
		if err != nil {
			log.Println("Comment.Action: ", err)
			reply.StatusCode = int32(StatusCommentFailed)
			reply.StatusMsg = StatusCommentFailed.msg()
			return reply
		}

		wg.Wait()
		if u != nil {
			c.User = *u
		}

		reply.Comment = c
	case OpDelete:
		if err := comment.DeleteByID(args.CommentID); err != nil {
			log.Println("Comment.Action: ", err)
			reply.StatusCode = int32(StatusDeleteFailed)
			reply.StatusMsg = StatusDeleteFailed.msg()
			return reply
		}
	default:
		reply.StatusCode = int32(StatusIllegalOperation)
		reply.StatusMsg = StatusIllegalOperation.msg()
		return reply
	}

	return reply
}

func userByID(id int64, reply *model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := user.ByID(id)
	if err != nil {
		log.Println("Comment.Action: ", err)
	}
	*reply = *r
}

const (
	OpComment = 1
	OpDelete  = 2
)

type Request struct {
	LoginID     int64
	VideoID     int64
	Type        int64
	CommentText string
	CommentID   int64
}

type Response struct {
	StatusCode int32         `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg,omitempty"` // 返回状态描述
	Comment    model.Comment `json:"comment,omitempty"`    // 评论成功返回评论内容，不需要重新拉取整个列表
}

type status int32

const (
	StatusOK status = iota
	StatusNoText
	StatusIllegalOperation
	StatusCommentFailed
	StatusDeleteFailed
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusNoText:
		return "评论内容为空"
	case StatusIllegalOperation:
		return "不合法的操作"
	case StatusCommentFailed:
		return "评论失败"
	case StatusDeleteFailed:
		return "删除失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
