package publish

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish"
	"github.com/PICOF/simple-tiktok/kitex_gen/publish/publishservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"strconv"
)

var Client publishservice.Client

func init() {
	InitPublish()
}

func InitPublish() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := publishservice.NewClient(
		constant.PublishServiceName,
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
