package publish

import (
	"log"
	"tiktok/internal/mapper"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
)

func List(args *in.PublishList) (*out.PublishList, error) {
	videos, err := mapper.VideosByUserID(args.UserID)
	if err != nil {
		log.Println("service.publish.List: ", err)
		return nil, err
	}

	reply := &out.PublishList{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
		VideoList:  videos,
	}
	return reply, nil
}
