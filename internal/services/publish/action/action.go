package action

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"tiktok/internal/services/model"
	"time"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var supportedExts = map[string]struct{}{
	".avi":  {},
	".flv":  {},
	".mp4":  {},
	".mpeg": {},
}

var (
	videoDir = "/home/abc/workspace/tiktok/public/videos/"
	coverDir = "/home/abc/workspace/tiktok/public/images/"
)

const coverExt = ".jpeg"

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// 检查扩展名
	ext := filepath.Ext(args.Filename)
	if _, ok := supportedExts[ext]; !ok {
		reply.StatusCode = int32(StatusVideoNotSupported)
		reply.StatusMsg = StatusVideoNotSupported.msg()
		return reply
	}
	// uuid生成随机文件名，不含扩展名
	name := uuid.NewString()
	// 设置文件路径
	playURL := videoDir + name + ext
	// 保存文件
	f, err := os.Create(playURL)
	if err != nil {
		log.Println("publish.action.Action: ", err)
		reply.StatusCode = int32(StatusUploadFailed)
		reply.StatusMsg = StatusUploadFailed.msg()
		return reply
	}

	if _, err = f.Write(args.Data); err != nil {
		log.Println("publish.action.Action: ", err)
		reply.StatusCode = int32(StatusUploadFailed)
		reply.StatusMsg = StatusUploadFailed.msg()
		f.Close()
		return reply
	}
	f.Close()

	coverURL := coverDir + name + coverExt
	// 取1帧作封面，保存
	if err = generateCover(coverURL, playURL, 1); err != nil {
		log.Println("publish.action.Action: ", err)
		reply.StatusCode = int32(StatusUploadFailed)
		reply.StatusMsg = StatusUploadFailed.msg()
		return reply
	}

	now := time.Now().UnixMilli()
	video := &model.Video{
		UserID:     args.LoginID,
		Title:      args.Title,
		PlayURL:    playURL,
		CoverURL:   coverURL,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	if err := createVideo(video); err != nil {
		log.Println("publish.action.Action: ", err)
		reply.StatusCode = int32(StatusUploadFailed)
		reply.StatusMsg = StatusUploadFailed.msg()
		return reply
	}

	return reply
}

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

type Request struct {
	LoginID  int64
	Title    string
	Filename string
	Data     []byte
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type status int32

const (
	StatusOK status = iota
	StatusVideoNotSupported
	StatusUploadFailed
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusVideoNotSupported:
		return "不支持的文件格式"
	case StatusUploadFailed:
		return "上传文件失败"
	default:
		return "未知错误"
	}
}
