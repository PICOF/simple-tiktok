package feed

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed/feedservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Client feedservice.Client

func init() {
	InitFeed()
}

func InitFeed() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := feedservice.NewClient(
		constant.FeedServiceName,
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

func GetVideoList(ctx context.Context, req *tiktokapi.FeedRequest, userId int64) (list *feed.FeedResponse, err error) {
	var rpcReq = &feed.FeedRequest{
		LatestTime: req.LatestTime,
		UserId:     &userId,
	}
	list, err = Client.GetVideoList(ctx, rpcReq)
	return
}
