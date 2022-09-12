package comment

import (
	"tiktok/internal/controllers/comment/action"
	"tiktok/internal/controllers/comment/list"
)

type Comment int

func (*Comment) Action(args *action.Request, reply *action.Response) error {
	r := action.Action(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.Comment = r.Comment
	return nil
}

func (*Comment) List(args *list.Request, reply *list.Response) error {
	r := list.List(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.CommentList = r.CommentList
	return nil
}
