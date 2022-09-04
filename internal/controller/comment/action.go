package comment

import "tiktok/internal/model"

type ActionResponse struct {
	Comment    model.Comment `json:"comment,omitempty"`    // 评论成功返回评论内容，不需要重新拉取整个列表
	StatusCode int32         `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string        `json:"status_msg,omitempty"` // 返回状态描述
}
