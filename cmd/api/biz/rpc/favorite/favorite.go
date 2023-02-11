package favorite

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite/favoriteservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
)

var (
	Client      favoriteservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("favorite")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitFavorite()
}

func InitFavorite() {
	c, err := favoriteservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func Action(ctx context.Context, req *tiktokapi.LikeRequest, userId int64) (resp *favorite.LikeResponse, err error) {
	videoId, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return nil, err
	}
	var rpcReq = &favorite.LikeRequest{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: req.ActionType == "1",
	}
	resp, err = Client.LikeAction(ctx, rpcReq)
	return
}
func GetFavoriteList(ctx context.Context, req *tiktokapi.LikeListRequest, userId int64) (resp *favorite.LikeListResponse, err error) {
	queryId, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	var rpcReq = &favorite.LikeListRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	resp, err = Client.GetLikeList(ctx, rpcReq)
	return
}
