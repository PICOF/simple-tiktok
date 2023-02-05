package service

import (
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
)

func PackResponse(list []*feed.VideoInfo, status constant.Status, nextTime *int64) (response *feed.FeedResponse) {
	code, msg := status.GetInfo()
	response = &feed.FeedResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		NextTime:   nextTime,
		VideoList:  list,
	}
	return
}
