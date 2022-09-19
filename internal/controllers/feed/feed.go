package feed

import (
	"log"
	"sort"
	"sync"
	"tiktok/internal/model"
	"tiktok/internal/services/comment"
	"tiktok/internal/services/favorite"
	"tiktok/internal/services/user"
	"tiktok/internal/services/video"
	"time"
)

type Feed int

func (*Feed) Feed(args *Request, reply *Response) error {
	r := feed(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.NextTime = r.NextTime
	reply.VideoList = r.VideoList
	return nil
}

func feed(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	now := time.Now().UnixMilli()
	if args.LatestTime != 0 {
		now = args.LatestTime
	}

	// 取视频基本信息
	videos, err := video.Before(now)
	if err != nil {
		log.Println("Feed.Feed: ", err)
		reply.StatusCode = int32(StatusFetchError)
		reply.StatusMsg = StatusFetchError.msg()
		return reply
	}
	if len(videos) == 0 {
		return reply
	}
	reply.NextTime = videos[len(videos)-1].ModifiedAt // 倒序，最后一个之前即下次

	// 作者、点赞数、评论数、登录用户的点赞互不影响
	var wg sync.WaitGroup
	wg.Add(3)
	go setUsers(&videos, &wg)
	go setFavoriteCounts(&videos, &wg)
	go setCommentCounts(&videos, &wg)
	if args.LoginID > 0 {
		// ok, err := login.CheckCache(args.LoginID)
		// if err != nil || !ok {
		// 	log.Println("Feed.Feed: ", err)
		// }
		// if err == nil {
		wg.Add(1)
		go setIsFavorites(args.LoginID, &videos, &wg)
	}
	wg.Wait()

	reply.VideoList = videos

	return reply
}

// 取作者信息
// userIDs := make([]int64, len(videos))
// userToVideo := make(map[int64]*model.Video)
// for i := range videos {
// 	userIDs[i] = videos[i].UserID
// 	userToVideo[videos[i].UserID] = &videos[i]
// }
// users, err := user.ByIDs(userIDs)
// if err != nil {
// 	log.Println("Feed.Feed: ", err)
// }
// if err == nil {
// 	for _, u := range users {
// 		userToVideo[u.ID].User = u
// 	}
// }

// // 取点赞数量
// // 摘出已选中的ID减小查询开销
// idToVideo := make(map[int64]*model.Video)
// ids := make([]int64, len(videos))
// for i := range videos {
// 	ids[i] = videos[i].ID
// 	idToVideo[videos[i].ID] = &videos[i]
// }
// favs, err := favorite.CountFavoritedsByVideoIDs(ids)
// if err != nil {
// 	log.Println("Feed.Feed: ", err)
// }
// if err == nil {
// 	for i, id := range ids {
// 		idToVideo[id].FavoriteCount = favs[i]
// 	}
// }

// // 取评论数量
// idToCommentCount, err := comment.CountsByVideoIDs(ids)
// if err != nil {
// 	log.Println("Feed.Feed: ", err)
// }
// if err == nil {
// 	for _, id := range ids {
// 		idToVideo[id].CommentCount = idToCommentCount[id]
// 	}
// }
// // 取登录用户关系，认为点赞人数>用户点赞视频数
// if args.LoginID > 0 {
// 	ok, err := login.CheckCache(args.LoginID)
// 	if err != nil || !ok {
// 		log.Println("Feed.Feed: ", err)
// 		return reply
// 	}
// 	vids, err := favorite.FavoritesByUserID(args.LoginID)
// 	if err != nil {
// 		log.Println("Feed.Feed: ", err)
// 	}
// 	if err == nil {
// 		sort.Slice(vids, func(i, j int) bool { return vids[i] < vids[j] })
// 		for i := range videos {
// 			t := sort.Search(len(vids), func(j int) bool { return vids[j] == videos[i].ID })
// 			if t < len(videos) {
// 				videos[i].IsFavorite = true
// 			}
// 		}
// 	}
// }

func setUsers(videos *[]model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	vs := *videos
	userIDs := make([]int64, len(vs))
	for i := range vs {
		userIDs[i] = vs[i].UserID
	}
	users, err := user.ByIDs(userIDs)
	if err != nil {
		log.Println("Feed.Feed: ", err)
		return
	}
	idToUser := make(map[int64]*model.User)
	for i := range users {
		idToUser[users[i].ID] = &users[i]
	}
	for i := range vs {
		vs[i].User = *idToUser[vs[i].UserID]
	}
}

func setFavoriteCounts(videos *[]model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	vs := *videos
	idToVideo := make(map[int64]*model.Video)
	ids := make([]int64, len(vs))
	for i := range vs {
		ids[i] = vs[i].ID
		idToVideo[vs[i].ID] = &vs[i]
	}
	favs, err := favorite.CountFavoritedsByVideoIDs(ids)
	if err != nil {
		log.Println("Feed.Feed: ", err)
		return
	}
	for i, id := range ids {
		idToVideo[id].FavoriteCount = favs[i]
	}
}

func setCommentCounts(videos *[]model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	vs := *videos
	idToVideo := make(map[int64]*model.Video)
	ids := make([]int64, len(vs))
	for i := range vs {
		ids[i] = vs[i].ID
		idToVideo[vs[i].ID] = &vs[i]
	}
	idToCommentCount, err := comment.CountsByVideoIDs(ids)
	if err != nil {
		log.Println("Feed.Feed: ", err)
	}
	if err == nil {
		for _, id := range ids {
			idToVideo[id].CommentCount = idToCommentCount[id]
		}
	}
}

func setIsFavorites(userID int64, videos *[]model.Video, wg *sync.WaitGroup) {
	defer wg.Done()
	vids, err := favorite.FavoritesByUserID(userID)
	if err != nil {
		log.Println("Feed.Feed: ", err)
		return
	}
	sort.Slice(vids, func(i, j int) bool { return vids[i] < vids[j] })

	vs := *videos
	for i := range vs {
		t := sort.Search(len(vids), func(j int) bool { return vids[j] == vs[i].ID })
		if t < len(vs) {
			vs[i].IsFavorite = true
		}
	}
}

type Request struct {
	LatestTime int64
	LoginID    int64
}

type Response struct {
	StatusCode int32         `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg,omitempty"` // 返回状态描述
	NextTime   int64         `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList  []model.Video `json:"video_list,omitempty"` // 视频列表
}

type status int32

const (
	StatusOK status = iota
	StatusFetchError
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusFetchError:
		return "无法获取数据"
	default:
		return "未知错误"
	}
}
