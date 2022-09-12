package user

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/controllers/user/user"

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

	args := &user.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := &user.Response{}

	cli := remote.Client

	userCall := cli.Go(remote.User+".User", args, reply, nil)
	replyCall := <-userCall.Done
	if replyCall.Error != nil {
		log.Println("user.User: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
