package feed

import "tiktok/internal/services/feed/feed"

type Feed int

func (*Feed) Feed(args *feed.Request, reply *feed.Response) error {
	r := feed.Feed(args)
	reply.StatusCode = r.StatusCode
	reply.StatusMsg = r.StatusMsg
	reply.NextTime = r.NextTime
	reply.VideoList = r.VideoList
	return nil
}
