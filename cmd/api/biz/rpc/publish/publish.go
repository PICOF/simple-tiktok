package publish

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish/publishservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
)

var (
	Client      publishservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("publish")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitPublish()
}

func InitPublish() {
	c, err := publishservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func GetPublishList(ctx context.Context, req *tiktokapi.PublishListRequest, userId int64) (list *publish.PublishListResponse, err error) {
	queryId, _ := strconv.ParseInt(req.GetUserID(), 10, 64)
	var rpcReq = &publish.PublishListRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	list, err = Client.GetPublishList(ctx, rpcReq)
	return
}

func Publish(ctx context.Context, req *tiktokapi.PublishRequest, userId int64) (resp *publish.PublishResponse, err error) {
	var rpcReq = &publish.PublishRequest{
		Data:   req.Data,
		UserId: userId,
		Title:  req.Title,
	}
	resp, err = Client.PublishAction(ctx, rpcReq)
	return
}
