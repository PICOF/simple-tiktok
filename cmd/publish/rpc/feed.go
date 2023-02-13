package rpc

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed"
	"github.com/PICOF/simple-tiktok/kitex_gen/feed/feedservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Client feedservice.Client

func init() {
	InitUser()
}

func InitUser() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.PublishServiceName),
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.PublishServiceName}),
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
