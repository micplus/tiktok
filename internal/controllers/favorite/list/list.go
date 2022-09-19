package list

import (
	"log"
	"sort"
	"sync"
	"tiktok/internal/model"
	"tiktok/internal/services/favorite"
	"tiktok/internal/services/video"
)

func List(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// ok, err := login.CheckCache(args.LoginID)
	// if err != nil || !ok {
	// 	log.Println("Favorite.List: ", err)
	// 	reply.StatusCode = int32(StatusTokenExpired)
	// 	reply.StatusMsg = StatusTokenExpired.msg()
	// 	return reply
	// }

	// 查目标用户点赞列表
	ids, err := favorite.FavoritesByUserID(args.UserID)
	if err != nil {
		log.Println("Favorite.List: ", err)
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}
	if len(ids) == 0 {
		reply.VideoList = []model.Video{}
		return reply
	}
	videos, err := video.ByIDs(ids)
	if err != nil {
		log.Println("Favorite.List: ", err)
		reply.StatusCode = int32(StatusListFailed)
		reply.StatusMsg = StatusListFailed.msg()
		return reply
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go setFavoriteCounts(&videos, ids, &wg)
	go setIsFavorites(args.LoginID, &videos, ids, &wg)
	wg.Wait()

	reply.VideoList = videos

	return reply
}

// // 查视频点赞数
// favoriteCount, err := favorite.CountFavoritedsByVideoIDs(ids)
// if err != nil {
// 	log.Println("Favorite.List: ", err)
// }
// if err == nil {
// 	// 更新videos中的点赞数信息
// 	for i := range videos {
// 		videos[i].FavoriteCount = favoriteCount[videos[i].ID]
// 	}
// }

// // 查自己点赞列表，更新红心
// favs, err := favorite.FavoritesByUserID(args.LoginID)
// if err != nil {
// 	log.Println("Favorite.List ", err)
// }
// if err == nil {
// 	sort.Slice(favs, func(i, j int) bool { return favs[i] < favs[j] })
// 	for i := range videos {
// 		t := sort.Search(len(favs), func(j int) bool { return favs[j] == videos[i].ID })
// 		if t < len(favs) {
// 			videos[i].IsFavorite = true
// 		}
// 	}
// }

func setFavoriteCounts(videos *[]model.Video, ids []int64, wg *sync.WaitGroup) {
	defer wg.Done()
	favoriteCount, err := favorite.CountFavoritedsByVideoIDs(ids)
	if err != nil {
		log.Println("Favorite.List: ", err)
		return
	}
	vs := *videos
	// 更新videos中的点赞数信息
	for i := range vs {
		vs[i].FavoriteCount = favoriteCount[i]
	}
}

func setIsFavorites(userID int64, videos *[]model.Video, ids []int64, wg *sync.WaitGroup) {
	defer wg.Done()
	favs, err := favorite.FavoritesByUserID(userID)
	if err != nil {
		log.Println("Favorite.List ", err)
		return
	}
	sort.Slice(favs, func(i, j int) bool { return favs[i] < favs[j] })

	vs := *videos
	for i := range vs {
		t := sort.Search(len(favs), func(j int) bool { return favs[j] == vs[i].ID })
		if t < len(favs) {
			vs[i].IsFavorite = true
		}
	}
}

type Request struct {
	LoginID int64
	UserID  int64
}

type Response struct {
	StatusCode int32         `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []model.Video `json:"video_list,omitempty"` // 用户点赞视频列表
}

type status int32

const (
	StatusOK status = iota
	StatusListFailed
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusListFailed:
		return "获取点赞列表失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
