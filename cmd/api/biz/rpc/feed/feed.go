package feed

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
	Client      feedservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("feed")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitFeed()
}

func InitFeed() {
	c, err := feedservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func GetVideoList(ctx context.Context, req *tiktokapi.FeedRequest, userId int64) (list *feed.FeedResponse, err error) {
	var rpcReq = &feed.FeedRequest{
		LatestTime: req.LatestTime,
		UserId:     &userId,
	}
	list, err = Client.GetVideoList(ctx, rpcReq)
	return
}
