package publish

import (
	"log"
	"net/http"
	"strconv"
	"tiktok/internal/services/publish/list"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// 中间件从token取id到key "login_id"
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)
	userID, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println("publish.List: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	args := &list.Request{
		LoginID: loginID,
		UserID:  userID,
	}

	reply := list.List(args)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	c.JSON(http.StatusOK, reply)
}