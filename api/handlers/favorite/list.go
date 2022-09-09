package favorite

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/api/remote"
	"tiktok/internal/services/favorite/list"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println("favorite.List: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &list.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := &list.Response{}

	cli := remote.Client

	listCall := cli.Go(remote.Favorite+".List", args, reply, nil)
	replyCall := <-listCall.Done
	if replyCall.Error != nil {
		log.Println("favorite.List: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
