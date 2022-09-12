package user

import (
	"log"
	"sort"
	"sync"
	"tiktok/internal/model"

	"tiktok/internal/services/login"
	"tiktok/internal/services/relation"
	"tiktok/internal/services/user"
)

func User(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// 检查登陆状态
	ok, err := login.CheckCache(args.LoginID)
	if err != nil || !ok {
		log.Println("User.User: ", err)
		reply.StatusCode = int32(StatusTokenExpired)
		reply.StatusMsg = StatusTokenExpired.msg()
		return reply
	}

	// 获取用户信息
	u, err := user.ByID(args.UserID)
	if err != nil {
		log.Println("User.User: ", err)
		reply.StatusCode = int32(StatusInternalErr)
		reply.StatusMsg = StatusInternalErr.msg()
		return reply
	}
	if u == nil {
		reply.StatusCode = int32(StatusNotExists)
		reply.StatusMsg = StatusNotExists.msg()
		return reply
	}

	// 三个操作修改的是&user的不同域
	var wg sync.WaitGroup
	wg.Add(3)
	go setFollowCount(args.UserID, u, &wg)
	go setFollowerCount(args.UserID, u, &wg)
	go setIsFollow(args.LoginID, args.UserID, u, &wg)
	wg.Wait()

	// // 获取关注数量
	// fwc, err := relation.CountFollowsByID(args.UserID)
	// if err != nil {
	// 	log.Println("User.User: ", err)
	// }
	// u.FollowCount = fwc
	// // 获取粉丝数量
	// frc, err := relation.CountFollowersByID(args.UserID)
	// if err != nil {
	// 	log.Println("User.User: ", err)
	// }
	// u.FollowerCount = frc
	// // 获取观察者与之关注关系
	// // 依常识，用户关注的人数比博主的粉丝数少，所以取用户的关注列表
	// follows, err := relation.FollowsByID(args.LoginID)
	// sort.Slice(follows, func(i, j int) bool { return follows[i] < follows[j] })
	// found := sort.Search(len(follows), func(i int) bool { return follows[i] == args.UserID })
	// if found < len(follows) {
	// 	u.IsFollow = true
	// }

	reply.User = u

	return reply
}

func setFollowCount(id int64, u *model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	fwc, err := relation.CountFollowsByID(id)
	if err != nil {
		log.Println("User.User: ", err)
	}
	u.FollowCount = fwc
}

func setFollowerCount(id int64, u *model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	fwc, err := relation.CountFollowersByID(id)
	if err != nil {
		log.Println("User.User: ", err)
	}
	u.FollowerCount = fwc
}

func setIsFollow(from, to int64, u *model.User, wg *sync.WaitGroup) {
	defer wg.Done()
	follows, err := relation.FollowsByID(from)
	if err != nil {
		log.Println("User.User: ", err)
		return
	}
	sort.Slice(follows, func(i, j int) bool { return follows[i] < follows[j] })
	found := sort.Search(len(follows), func(i int) bool { return follows[i] == to })
	if found < len(follows) {
		u.IsFollow = true
	}
}

type Request struct {
	LoginID int64
	UserID  int64
}

type Response struct {
	StatusCode int32       `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg,omitempty"` // 返回状态描述
	User       *model.User `json:"user"`                 // 用户信息
}

func logError(err error) {
	if err != nil {
		log.Println("user.User: ", err)
	}
}

const (
	StatusOK status = iota
	StatusNotExists
	StatusInternalErr
	StatusTokenExpired
)

type status int32

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusNotExists:
		return "要找的用户不存在"
	case StatusInternalErr:
		return "获取用户失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
