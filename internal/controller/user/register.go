package user

import (
	"net/http"
	"tiktok/internal/service/user"
	"tiktok/internal/transport/request"

	"github.com/gin-gonic/gin"
)

type RegisterResponse struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
	Token      string `json:"token"`                // 用户鉴权token
	UserID     int64  `json:"user_id"`              // 用户id
}

func Register(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &request.Register{
		Username: username,
		Password: password,
	}

	reply, err := user.Register(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
