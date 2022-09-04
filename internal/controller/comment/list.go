package comment

import "tiktok/internal/model"

type ListResponse struct {
	CommentList []model.Comment `json:"comment_list,omitempty"` // 评论列表
	StatusCode  int32           `json:"status_code"`            // 状态码，0-成功，其他值-失败
	StatusMsg   string          `json:"status_msg,omitempty"`   // 返回状态描述
}
