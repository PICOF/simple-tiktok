package service

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/dal/operation"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"strconv"
	"time"
)

func GetVideoList(ctx context.Context, request *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	var latestTime int64
	var userId int64
	if request.LatestTime == nil {
		latestTime = time.Now().UnixNano()
	} else {
		latestTime, err = strconv.ParseInt(*request.LatestTime, 10, 64)
		if err != nil {
			return
		}
	}
	if request.Token != nil {
		userId, err = ParseToken(*request.Token)
		if err != nil {
			resp = PackResponse(nil, constant.Failed, nil)
			return
		}
	}
	list, nextTime, err := operation.GetVideoList(ctx, userId, latestTime)
	if err != nil {
		resp = PackResponse(nil, constant.Failed, nil)
		return
	}
	resp = PackResponse(list, constant.Success, nextTime)
	return
}
