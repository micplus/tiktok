package comment

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/controllers/comment/action"

	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		log.Println("comment.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil {
		log.Println("comment.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if err != nil {
		commentID = 0
	}

	commentText := c.Query("comment_text")

	args := &action.Request{
		LoginID:     loginID,
		VideoID:     videoID,
		Type:        actionType,
		CommentText: commentText,
		CommentID:   commentID,
	}

	reply := &action.Response{}

	cli := remote.Client
	actionCall := cli.Go(remote.Comment+".Action", args, reply, nil)
	replyCall := <-actionCall.Done

	if replyCall.Error != nil {
		log.Println("comment.Action: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := action.Action(args)

	c.JSON(http.StatusOK, reply)
}
