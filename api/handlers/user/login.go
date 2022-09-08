package user

import (
	"net/http"
	"tiktok/internal/services/user/login"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	args := &login.Request{
		Username: username,
		Password: password,
	}

	reply := login.Login(args)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	c.JSON(http.StatusOK, reply)
}
