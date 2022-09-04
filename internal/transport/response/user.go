package response

import "tiktok/internal/model"

type User struct {
	StatusCode int32       `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string      `json:"status_msg,omitempty"` // 返回状态描述
	User       *model.User `json:"user"`                 // 用户信息
}

type Register struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	Token      string `json:"token"`                // 用户鉴权token
	UserID     int64  `json:"user_id"`              // 用户id
}

type Login struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	Token      string `json:"token"`                // 用户鉴权token
	UserID     int64  `json:"user_id"`              // 用户id
}
