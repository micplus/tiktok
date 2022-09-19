package action

import (
	"log"
	"tiktok/internal/model"
	"tiktok/internal/services/video"
	"time"
)

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

func Action(args *Request) *Response {
	reply := &Response{
		StatusCode: int32(StatusOK),
		StatusMsg:  StatusOK.msg(),
	}

	// // 检查登陆状态
	// ok, err := login.CheckCache(args.LoginID)
	// if err != nil || !ok {
	// 	log.Println("Publish.Action: ", err)
	// 	log.Println(ok)
	// 	reply.StatusCode = int32(StatusTokenExpired)
	// 	reply.StatusMsg = StatusTokenExpired.msg()
	// 	return reply
	// }

	now := time.Now().UnixMilli()
	v := &model.Video{
		UserID:     args.LoginID,
		Title:      args.Title,
		PlayURL:    args.PlayURL,
		CoverURL:   args.CoverURL,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	_, err := video.Insert(v)
	if err != nil {
		log.Println("Publish.Action: ", err)
		reply.StatusCode = int32(StatusUploadFailed)
		reply.StatusMsg = StatusUploadFailed.msg()
		return reply
	}

	return reply
}

type Request struct {
	LoginID  int64
	Title    string
	PlayURL  string
	CoverURL string
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
	StatusTokenExpired
)

func (s status) msg() string {
	switch s {
	case StatusOK:
		return "OK"
	case StatusVideoNotSupported:
		return "不支持的文件格式"
	case StatusUploadFailed:
		return "上传文件失败"
	case StatusTokenExpired:
		return "登录过期，请重新登录"
	default:
		return "未知错误"
	}
}
