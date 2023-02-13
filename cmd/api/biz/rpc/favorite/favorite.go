package favorite

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite"
	"github.com/PICOF/simple-tiktok/kitex_gen/favorite/favoriteservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"strconv"
)

var Client favoriteservice.Client

func init() {
	InitFavorite()
}

func InitFavorite() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := favoriteservice.NewClient(
		constant.FavoriteServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.ServerServiceName}),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func Action(ctx context.Context, req *tiktokapi.LikeRequest, userId int64) (resp *favorite.LikeResponse, err error) {
	videoId, _ := strconv.ParseInt(req.VideoID, 10, 64)
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
