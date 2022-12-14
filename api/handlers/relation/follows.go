package relation

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/controllers/relation/follows"

	"github.com/gin-gonic/gin"
)

func Follows(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println("relation.Follows: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &follows.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := &follows.Response{}

	cli := remote.Client

	listCall := cli.Go(remote.Relation+".Follows", args, reply, nil)
	replyCall := <-listCall.Done
	if replyCall.Error != nil {
		log.Println("relation.Follows: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
