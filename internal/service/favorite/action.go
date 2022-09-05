package favorite

import (
	in "tiktok/internal/transport/request"
	out "tiktok/internal/transport/response"
)

// 点赞应当是一个实时任务，不仅我自己能看到点赞结果，别人点赞也应该很快被我看到
// 不断+1 +1，而且要加锁，开销大
// -> 点赞到缓存，每隔一段时间持久化到数据库

func Action(args *in.FavoriteAction) (*out.FavoriteAction, error) {
	// cache点赞
	return nil, nil
}
