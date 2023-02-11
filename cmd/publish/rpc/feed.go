package rpc

import (
	"context"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed/feedservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var (
	Client      feedservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("feed")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitUser()
}

func InitUser() {
	c, err := feedservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func GetVideoListById(ctx context.Context, userId int64, queryId int64) (videoInfo []*feed.VideoInfo, err error) {
	var rpcReq = &feed.GetByIDRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	var resp *feed.FeedResponse
	resp, err = Client.GetVideoListById(ctx, rpcReq)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get video by id: %v", err)
		return
	}
	videoInfo = resp.VideoList
	return
}
