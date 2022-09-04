package relation

import "tiktok/internal/model"

type FollowerListResponse struct {
	StatusCode int32        `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string       `json:"status_msg,omitempty"` // 返回状态描述
	UserList   []model.User `json:"user_list,omitempty"`  // 用户列表
}
