package relation

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/controllers/relation/followers"

	"github.com/gin-gonic/gin"
)

func Followers(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println("relation.Followers: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &followers.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := &followers.Response{}

	cli := remote.Client

	listCall := cli.Go(remote.Relation+".Followers", args, reply, nil)
	replyCall := <-listCall.Done
	if replyCall.Error != nil {
		log.Println("relation.Followers: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
