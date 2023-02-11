package publishaction

import (
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish"
)

func PackListResponse(list []*feed.VideoInfo, status constant.Status) (response *publish.PublishListResponse) {
	code, msg := status.GetInfo()
	response = &publish.PublishListResponse{
		StatusCode: code,
		StatusMsg:  &msg,
		VideoList:  list,
	}
	return
}
