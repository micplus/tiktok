package feed

import (
	"log"

	"tiktok/internal/services/model"
	"time"
)

func Feed(args *Request) (*Response, error) {
	now := time.Now().UnixMilli()
	if args.LatestTime != 0 {
		now = args.LatestTime
	}

	// 取视频
	videos, err := videosBeforeTime(now)
	printError(err)
	if len(videos) == 0 {
		return &Response{
			StatusCode: int32(StatusOK),
			StatusMsg:  StatusOK.msg(),
		}, nil
	}

	// 记录选出的ID（30条）
	ids := chosenIDs(videos)

	// 查视频点赞数
	favoriteCount, err := favoriteCountsByVideoIDs(ids)
	printError(err)
	if err == nil {
		// 更新videos中的点赞数信息
		for i := range videos {
			videos[i].FavoriteCount = favoriteCount[videos[i].ID]
		}
	}

	// 查评论数
	commentCount, err := commentCountsByVideoIDs(ids)
	printError(err)
	if err == nil {
		// 更新videos中的点赞数信息
		for i := range videos {
			videos[i].CommentCount = commentCount[videos[i].ID]
		}
	}

	// 更新当前用户看到的点赞信息
	if args.LoginID != 0 { // 已登录
		favorites, err := favoritesByUserID(args.LoginID)
		printError(err)
		f := make(map[int64]struct{})
		for _, id := range favorites {
			f[id] = struct{}{}
		}
		if err == nil {
			for i := range videos {
				if _, ok := f[videos[i].ID]; ok {
					videos[i].IsFavorite = true
				}
			}
		}
	}

	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
		NextTime:   videos[len(videos)-1].CreatedAt, // 倒序，最后一个之前即下次
		VideoList:  videos,
	}
	return reply, nil
}

func printError(err error) {
	if err != nil {
		log.Println("services.Feed: ", err)
	}
}

func chosenIDs(videos []model.Video) []int64 {
	// 选中的视频列表
	ids := make([]int64, len(videos))
	for i, video := range videos {
		ids[i] = video.ID
	}
	return ids
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
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	default:
		return "未知错误"
	}
}
