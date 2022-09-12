package publish

import (
	"tiktok/internal/controllers/publish/action"
	"tiktok/internal/controllers/publish/list"
)

type Publish int

func (*Publish) Action(args *action.Request, reply *action.Response) error {
	r := action.Action(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	return nil
}

func (*Publish) List(args *list.Request, reply *list.Response) error {
	r := list.List(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.VideoList = r.VideoList
	return nil
}
