package feed

import (
	"tiktok/internal/mapper"
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
	"tiktok/internal/transport/status"
	"time"
)

func Feed(args *in.Feed) (*out.Feed, error) {
	now := time.Now().UnixMilli()
	if args.LatestTime != 0 {
		now = args.LatestTime
	}

	videos, err := mapper.VideosByTimeUserID(now, args.LoginID)

	// videos, err := mapper.VideosByTime(now)
	// if err != nil {
	// 	log.Println("service.feed.feed: ", err)
	// 	return nil, err
	// }

	if len(videos) == 0 {
		return &out.Feed{
			StatusCode: int32(status.OK),
			StatusMsg:  status.OK.Message(),
		}, nil
	}

	reply := &out.Feed{
		StatusCode: int32(status.OK),
		StatusMsg:  status.OK.Message(),
		NextTime:   videos[len(videos)-1].CreatedAt,
		VideoList:  videos,
	}

	return reply, nil
}
