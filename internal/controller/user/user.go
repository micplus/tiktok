package user

import (
	"net/http"
	"strconv"
	"tiktok/internal/service/user"
	"tiktok/internal/transport/request"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}
	token := c.Query("token")

	args := &request.User{
		ID:    id,
		Token: token,
	}

	reply, err := user.User(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
