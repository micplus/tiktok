package action

import (
	"fmt"
	"log"
	"tiktok/internal/services/model"
	"time"
)

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	if args.Type == OpComment && len(args.CommentText) == 0 {
		reply.StatusCode = int32(StatusNoText)
		reply.StatusMsg = StatusNoText.msg()
		return reply
	}

	switch args.Type {
	case OpComment:
		now := time.Now()
		month, day := now.Month(), now.Day()
		date := fmt.Sprintf("%d-%d", month, day)

		c := model.Comment{
			Content:    args.CommentText,
			CreateDate: date,
			UserID:     args.LoginID,
			VideoID:    args.VideoID,
			CreatedAt:  now.UnixMilli(),
			ModifiedAt: now.UnixMilli(),
		}

		if err := createComment(&c); err != nil {
			log.Println("comment.action.Action: ", err)
			reply.StatusCode = int32(StatusCommentFailed)
			reply.StatusMsg = StatusCommentFailed.msg()
			return reply
		}

		user, err := userByID(args.LoginID)
		if err != nil {
			log.Println("comment.action.Action: ", err)
		}
		if err == nil {
			c.User = *user
		}

		reply.Comment = c

	case OpDelete:
		if err := deleteCommentByID(args.CommentID); err != nil {
			log.Println("comment.action.Action: ", err)
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
	default:
		return "未知错误"
	}
}
