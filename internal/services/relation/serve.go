package relation

import (
	"tiktok/internal/services/relation/action"
	fwl "tiktok/internal/services/relation/follow/list"
	frl "tiktok/internal/services/relation/follower/list"
)

type Relation int

func (*Relation) Action(args *action.Request, reply *action.Response) error {
	r := action.Action(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	return nil
}

func (*Relation) FollowList(args *fwl.Request, reply *fwl.Response) error {
	r := fwl.List(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserList = r.UserList
	return nil
}

func (*Relation) FollowerList(args *frl.Request, reply *frl.Response) error {
	r := frl.List(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserList = r.UserList
	return nil
}
