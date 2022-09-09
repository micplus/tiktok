package comment

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/services/comment/list"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Println("comment.List: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &list.Request{
		LoginID: loginID,
		VideoID: videoID,
	}

	reply := &list.Response{}

	cli := remote.Client

	listCall := cli.Go(remote.Comment+".List", args, reply, nil)
	replyCall := <-listCall.Done
	if replyCall.Error != nil {
		log.Println("comment.List: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
