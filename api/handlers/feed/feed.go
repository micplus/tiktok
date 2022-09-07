package feed

import (
	"net/http"
	"strconv"

	"tiktok/internal/services/feed"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	// 参数解析
	now := int64(0)
	if latestTime := c.Query("latest_time"); latestTime != "" {
		if lt, err := strconv.ParseInt(latestTime, 10, 64); err == nil {
			now = lt
		}
	}

	loginID := int64(0)
	if loginIDAny, ok := c.Get("login_id"); ok {
		loginID = loginIDAny.(int64)
	}

	// 包装请求
	args := &feed.Request{
		LatestTime: now,
		LoginID:    loginID,
	}

	// 调用服务
	reply, err := feed.Feed(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
