package feed

import (
	"net/http"
	"strconv"
	"tiktok/internal/service/feed"
	"tiktok/internal/transport/request"

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
	token := c.Query("token")

	// 包装请求
	args := &request.Feed{
		LatestTime: now,
		Token:      token,
	}

	// 调用服务
	reply, err := feed.Feed(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
