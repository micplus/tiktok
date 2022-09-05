package publish

import (
	"log"
	"net/http"
	"tiktok/internal/service/publish"
	"tiktok/internal/transport/request"

	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// 中间件从token取id
	idany, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	id := idany.(int64)

	args := &request.PublishAction{
		Title:  title,
		UserID: id,
		Data:   data,
	}

	reply, err := publish.Action(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, reply)
}
