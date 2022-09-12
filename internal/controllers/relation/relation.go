package relation

import (
	"tiktok/internal/controllers/relation/action"
	"tiktok/internal/controllers/relation/followers"
	"tiktok/internal/controllers/relation/follows"
)

type Relation int

func (*Relation) Action(args *action.Request, reply *action.Response) error {
	r := action.Action(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	return nil
}

func (*Relation) Follows(args *follows.Request, reply *follows.Response) error {
	r := follows.Follows(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserList = r.UserList
	return nil
}

func (*Relation) Followers(args *followers.Request, reply *followers.Response) error {
	r := followers.Followers(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserList = r.UserList
	return nil
}
