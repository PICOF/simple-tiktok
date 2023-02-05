package operation

import (
	"context"
	"github.com/PICOF/simple-tiktok/dal"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"time"
)

func GetVideoList(ctx context.Context, userId int64, latestTime int64) (list []*feed.VideoInfo, nextTime *int64, err error) {
	var origin []TVideoInfo
	err = dal.DB.WithContext(ctx).Limit(30).Order("publish_time desc").Where("publish_time < ?", time.UnixMilli(latestTime).Format("2006-01-02 15:04:05")).Find(&origin).Error
	if err != nil {
		return
	}
	list, err = ConvertAllVideoInfo(ctx, userId, origin)
	if err != nil {
		return
	}
	if len(origin) == 0 {
		nextTime = nil
	} else {
		timestamp := origin[0].PublishTime.Unix()
		nextTime = &timestamp
	}
	return
}

func ConvertVideoInfo(ctx context.Context, userId int64, info TVideoInfo) (res *feed.VideoInfo, err error) {
	user, err := GetUserById(ctx, userId, info.AuthorId)
	if err != nil {
		return
	}
	isFavorite, err := IsFavorite(ctx, userId, info.Id)
	if err != nil {
		return
	}
	res = &feed.VideoInfo{
		Id:            info.Id,
		Author:        user,
		PlayUrl:       info.PlayUrl,
		CoverUrl:      info.CoverUrl,
		FavoriteCount: info.FavoriteCount,
		CommentCount:  info.CommentCount,
		IsFavorite:    isFavorite,
		Title:         info.Title,
	}
	return
}

func ConvertAllVideoInfo(ctx context.Context, userId int64, list []TVideoInfo) (res []*feed.VideoInfo, err error) {
	for _, t := range list {
		info, err := ConvertVideoInfo(ctx, userId, t)
		if err != nil {
			return nil, err
		}
		res = append(res, info)
	}
	return
}
