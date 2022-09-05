package response

import "tiktok/internal/model"

type Feed struct {
	NextTime   int64          `json:"next_time,omitempty"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	StatusCode int32          `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string         `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []*model.Video `json:"video_list,omitempty"` // 视频列表
}
