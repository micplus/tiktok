package favorite

import (
	"tiktok/internal/services/favorite/action"
	"tiktok/internal/services/favorite/list"
)

type Favorite int

func (*Favorite) Action(args *action.Request, reply *action.Response) error {
	r := action.Action(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	return nil
}

func (*Favorite) List(args *list.Request, reply *list.Response) error {
	r := list.List(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.VideoList = r.VideoList
	return nil
}
