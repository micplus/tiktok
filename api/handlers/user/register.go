package user

import (
	"log"
	"net/http"
	"tiktok/api/remote"
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

	reply := &register.Response{}

	cli := remote.Client

	userCall := cli.Go(remote.User+".Register", args, reply, nil)
	replyCall := <-userCall.Done
	if replyCall.Error != nil {
		log.Println("user.Register: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
