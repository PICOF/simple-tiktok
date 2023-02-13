package rpc

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/user"
	"github.com/PICOF/simple-tiktok/kitex_gen/user/userservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var Client userservice.Client

func init() {
	InitUser()
}

func InitUser() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.RelationServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		constant.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constant.RelationServiceName}),
	)
	if err != nil {
		panic(err)
	}
	Client = c
}

func GetUserInfo(ctx context.Context, userId int64, queryId int64) (userInfo *user.UserInfo, err error) {
	var rpcReq = &user.UserInfoRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	var resp *user.UserInfoResponse
	resp, err = Client.GetUserInfo(ctx, rpcReq)
	if err != nil {
		klog.CtxErrorf(ctx, "Failed to get user info: %v", err)
		return
	}
	userInfo = resp.User
	return
}
