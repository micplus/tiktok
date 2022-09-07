package middlewares

import (
	"net/http"
	"tiktok/internal/pkg/auth"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件鉴权，token失效则废弃当前请求
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		claims, ok := auth.ParseToken(token)
		if !ok {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}
		// 请求中的token有效，使用token携带的用户id
		c.Set("login_id", claims.UserID)
		c.Next()
	}
}
