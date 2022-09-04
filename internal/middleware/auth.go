package middleware

import (
	"net/http"
	"tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var key = []byte("Micplus-tiktok")

type Claims struct {
	UserID int64
	jwt.StandardClaims
}

func ReleaseToken(id int64) (string, error) {
	expire := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "tiktok",
			Subject:   "user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return ts, nil
}

func ParseToken(token string) (*Claims, bool) {
	t, _ := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, true
	}
	return nil, false
}

// 中间件鉴权，token失效则废弃当前请求
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusOK, response.Common{
				StatusCode: int32(status.NoLogin),
				StatusMsg:  status.NoLogin.Message(),
			})
			c.Abort()
			return
		}
		claims, ok := ParseToken(token)
		if !ok {
			c.JSON(http.StatusOK, response.Common{
				StatusCode: int32(status.NoLogin),
				StatusMsg:  status.NoLogin.Message(),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusOK, response.Common{
				StatusCode: int32(status.TokenExpired),
				StatusMsg:  status.TokenExpired.Message(),
			})
			c.Abort()
			return
		}
		// 请求中的token有效，使用token携带的用户id
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
