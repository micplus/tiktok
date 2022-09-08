package publish

import (
	"io"
	"log"
	"net/http"
	"tiktok/internal/services/publish/action"

	"github.com/gin-gonic/gin"
)

func Action(c *gin.Context) {
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// 读出数据
	filename := data.Filename
	file, err := data.Open()
	if err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		file.Close()
		return
	}
	file.Close()

	// 中间件从token取id
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		file.Close()
		return
	}
	loginID := loginIDAny.(int64)

	args := &action.Request{
		Title:    title,
		LoginID:  loginID,
		Filename: filename,
		Data:     bytes,
	}

	reply := action.Action(args)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, nil)
	//	return
	//}
	c.JSON(http.StatusOK, reply)
}
