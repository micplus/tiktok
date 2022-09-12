package user

import (
	"tiktok/internal/controllers/user/login"
	"tiktok/internal/controllers/user/register"
	"tiktok/internal/controllers/user/user"
)

type User int

func (*User) User(args *user.Request, reply *user.Response) error {
	r := user.User(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.User = r.User
	return nil
}

func (*User) Login(args *login.Request, reply *login.Response) error {
	r := login.Login(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserID = r.UserID
	reply.Token = r.Token
	return nil
}

func (*User) Register(args *register.Request, reply *register.Response) error {
	r := register.Register(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.UserID = r.UserID
	reply.Token = r.Token
	return nil
}
