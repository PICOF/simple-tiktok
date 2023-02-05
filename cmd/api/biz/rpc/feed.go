package rpc

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed/feedservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
)

var (
	feedClient  feedservice.Client
	feedConfig  *viper.Viper
	address     string
	serviceName string
)

func init() {
	feedConfig = config.GetConfig("feed")
	address = feedConfig.GetString("server.address")
	serviceName = feedConfig.GetString("server.serviceName")
}

func initFeed() {
	c, err := feedservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

func GetVideoList(ctx context.Context, req *tiktokapi.FeedRequest) (list *feed.FeedResponse, err error) {
	var rpcReq = &feed.FeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	}
	list, err = feedClient.GetVideoList(ctx, rpcReq)
	return
}
