package list

import (
	"log"
	"tiktok/internal/services/model"
)

const ServiceName = "CommentList"

type Reg int

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	comments, err := commentsByVideoID(args.VideoID)
	if err != nil {
		log.Println("comment.list.List: ", err)
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
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
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusListFailed:
		return "获取评论列表失败"
	default:
		return "未知错误"
	}
}
