package publish

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"tiktok/api/config"
	"tiktok/api/remote"
	"tiktok/internal/controllers/publish/action"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Action(c *gin.Context) {
	// 中间件从token取id
	loginIDAny, ok := c.Get("login_id")
	if !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	loginID := loginIDAny.(int64)

	title := c.PostForm("title")

	file, err := c.FormFile("data")
	if err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	ext := filepath.Ext(file.Filename)
	if _, ok := supportedExts[ext]; !ok {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	// uuid生成随机文件名，不含扩展名
	name := uuid.NewString()
	// 设置文件路径
	playURL := videoDir + name + ext

	if err = c.SaveUploadedFile(file, config.StaticDir+playURL); err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	coverURL := coverDir + name + coverExt
	// 取1帧作封面，保存
	if err = generateCover(config.StaticDir+coverURL, config.StaticDir+playURL, 1); err != nil {
		log.Println("publish.Action: ", err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	// TODO: 文件上传到云端存储

	args := &action.Request{
		Title:    title,
		LoginID:  loginID,
		PlayURL:  config.StaticAddr + "/static" + playURL,
		CoverURL: config.StaticAddr + "/static" + coverURL,
	}

	log.Println(args.PlayURL)
	log.Println(args.CoverURL)

	reply := &action.Response{}

	cli := remote.Client
	actionCall := cli.Go(remote.Publish+".Action", args, reply, nil)
	replyCall := <-actionCall.Done

	if replyCall.Error != nil {
		log.Println("publish.Action: ", replyCall.Error)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	// reply := action.Action(args)

	c.JSON(http.StatusOK, reply)
}

var supportedExts = map[string]struct{}{
	".avi":  {},
	".flv":  {},
	".mp4":  {},
	".mpeg": {},
}

const (
	coverExt = ".jpeg"

	videoDir = "/videos/"
	coverDir = "/covers/"
)

func generateCover(coverName, videoName string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		return err
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		return err
	}

	if err = imaging.Save(img, coverName); err != nil {
		return err
	}
	return nil
}
