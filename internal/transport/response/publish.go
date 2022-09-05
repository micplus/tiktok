package response

import "tiktok/internal/model"

type PublishAction struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type PublishList struct {
	StatusCode int32          `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string         `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []*model.Video `json:"video_list,omitempty"` // 用户发布的视频列表
}
