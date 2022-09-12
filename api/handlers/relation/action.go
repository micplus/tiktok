package relation

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/controllers/relation/action"

	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		log.Println("relation.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil {
		log.Println("relation.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &action.Request{
		LoginID: loginID,
		UserID:  userID,
		Type:    actionType,
	}

	reply := &action.Response{}

	cli := remote.Client
	actionCall := cli.Go(remote.Relation+".Action", args, reply, nil)
	replyCall := <-actionCall.Done

	if replyCall.Error != nil {
		log.Println("relation.Action: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := action.Action(args)

	c.JSON(http.StatusOK, reply)
}
