package feedaction

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/feed/rpc"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
)

func ConvertVideoInfo(ctx context.Context, userId int64, info operation.TVideoInfo) (res *feed.VideoInfo, err error) {
	author, err := rpc.GetUserInfo(ctx, userId, info)
	if err != nil {
		return
	}
	isFavorite, err := operation.IsFavorite(ctx, userId, info.Id)
	if err != nil {
		return
	}
	res = &feed.VideoInfo{
		Id:            info.Id,
		Author:        author,
		PlayUrl:       info.PlayUrl,
		CoverUrl:      info.CoverUrl,
		FavoriteCount: info.FavoriteCount,
		CommentCount:  info.CommentCount,
		IsFavorite:    isFavorite,
		Title:         info.Title,
	}
	return
}

func ConvertAllVideoInfo(ctx context.Context, userId int64, list []operation.TVideoInfo) (res []*feed.VideoInfo, err error) {
	for _, t := range list {
		info, err := ConvertVideoInfo(ctx, userId, t)
		if err != nil {
			return nil, err
		}
		res = append(res, info)
	}
	return
}
