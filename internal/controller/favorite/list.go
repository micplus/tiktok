package favorite

import "tiktok/internal/model"

type ListResponse struct {
	StatusCode int32          `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string         `json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []*model.Video `json:"video_list,omitempty"` // 用户点赞视频列表
}
