package publish

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"tiktok/config"
	"tiktok/internal/mapper"
	"tiktok/internal/model"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"

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
	videoDir = config.Detail.Static.Path + "/videos/"
	coverDir = config.Detail.Static.Path + "/images/"
)

const coverExt = ".jpeg"

func Action(args *in.PublishAction) (*out.PublishAction, error) {
	data := args.Data
	// 检查扩展名
	ext := filepath.Ext(data.Filename)
	if _, ok := supportedExts[ext]; !ok {
		return &out.PublishAction{
			StatusCode: int32(status.VideoNotSupported),
			StatusMsg:  status.VideoNotSupported.Message(),
		}, nil
	}
	// uuid生成随机文件名，不含扩展名
	name := uuid.NewString()
	// 设置文件路径
	videoPath := videoDir + name + ext
	// 保存文件
	file, err := data.Open()
	if err != nil {
		log.Println("service.publish.Action: ", err)
		return nil, err
	}
	defer file.Close()

	vf, err := os.Create(videoPath)
	if err != nil {
		log.Println("service.publish.Action: ", err)
		return nil, err
	}
	defer vf.Close()

	if _, err = io.Copy(vf, file); err != nil {
		log.Println("service.publish.Action: ", err)
		return nil, err
	}

	coverPath := coverDir + name + coverExt
	// 取1帧作封面，保存
	if err = generateCover(coverPath, videoPath, 1); err != nil {
		log.Println("service.publish.Action: ", err)
		return nil, err
	}

	video := &model.Video{
		UserID:   args.LoginID,
		Title:    args.Title,
		PlayURL:  videoPath,
		CoverURL: coverPath,
	}

	if err := mapper.CreateVideoWithUserID(video, args.LoginID); err != nil {
		log.Println("service.publish.Action: ", err)
		return nil, err
	}

	reply := &out.PublishAction{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
	}
	return reply, nil
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
