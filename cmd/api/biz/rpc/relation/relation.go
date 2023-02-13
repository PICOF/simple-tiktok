package relation

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/cmd/constant"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation/relationservice"
	"github.com/PICOF/simple-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"strconv"
)

const followType = "1"

var Client relationservice.Client

func init() {
	InitRelation()
}

func InitRelation() {
	r, err := etcd.NewEtcdResolver(constant.ETCDAddress)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(constant.ServerServiceName),
		provider.WithExportEndpoint(constant.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := relationservice.NewClient(
		constant.RelationServiceName,
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

func FollowAction(ctx context.Context, req *tiktokapi.FollowRequest, userId int64) (list *relation.FollowResponse, err error) {
	toUserId, _ := strconv.ParseInt(req.GetToUserID(), 10, 64)
	var rpcReq = &relation.FollowRequest{
		UserId:     userId,
		ToUserId:   toUserId,
		ActionType: req.GetActionType() == followType,
	}
	list, err = Client.FollowAction(ctx, rpcReq)
	return
}
func GetFollowList(ctx context.Context, req *tiktokapi.RelationListRequest, userId int64) (list *relation.RelationListResponse, err error) {
	queryId, _ := strconv.ParseInt(req.GetUserID(), 10, 64)
	var rpcReq = &relation.RelationListRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	list, err = Client.GetFollowList(ctx, rpcReq)
	return
}
func GetFollowerList(ctx context.Context, req *tiktokapi.RelationListRequest, userId int64) (list *relation.RelationListResponse, err error) {
	queryId, _ := strconv.ParseInt(req.GetUserID(), 10, 64)
	var rpcReq = &relation.RelationListRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	list, err = Client.GetFollowerList(ctx, rpcReq)
	return
}
func GetFriendList(ctx context.Context, req *tiktokapi.RelationListRequest, userId int64) (list *relation.RelationListResponse, err error) {
	queryId, _ := strconv.ParseInt(req.GetUserID(), 10, 64)
	var rpcReq = &relation.RelationListRequest{
		UserId:  userId,
		QueryId: queryId,
	}
	list, err = Client.GetFriendList(ctx, rpcReq)
	return
}
