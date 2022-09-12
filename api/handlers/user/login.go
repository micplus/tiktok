package user

import (
	"log"
	"net/http"
	"tiktok/api/remote"
	"tiktok/internal/controllers/user/login"

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

	reply := &login.Response{}

	cli := remote.Client

	userCall := cli.Go(remote.User+".Login", args, reply, nil)
	replyCall := <-userCall.Done
	if replyCall.Error != nil {
		log.Println("user.Login: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
