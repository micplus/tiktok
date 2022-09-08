package user

import (
	"net/http"
	"tiktok/internal/services/user/register"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &register.Request{
		Username: username,
		Password: password,
	}

	reply := register.Register(args)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	c.JSON(http.StatusOK, reply)
}
