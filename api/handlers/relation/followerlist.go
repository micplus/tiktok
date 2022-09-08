package relation

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/internal/services/relation/follower/list"

	"github.com/gin-gonic/gin"
)

func FollowerList(c *gin.Context) {
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println("relation.FollowerList: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &list.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := list.List(args)

	c.JSON(http.StatusOK, reply)
}
