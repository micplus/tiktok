package user

import (
	"net/http"
	"strconv"
	"tiktok/internal/service/user"
	"tiktok/internal/transport/request"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &request.User{
		LoginID: loginID,
		UserID:  userID,
	}

	reply, err := user.User(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
