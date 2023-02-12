package relation

import (
	"context"
	"github.com/PICOF/simple-tiktok/cmd/api/biz/model/tiktokapi"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation"
	"github.com/PICOF/simple-tiktok/kitex_gen/relation/relationservice"
	"github.com/PICOF/simple-tiktok/pkg/config"
	"github.com/cloudwego/kitex/client"
	"github.com/spf13/viper"
	"strconv"
)

const followType = "1"

var (
	Client      relationservice.Client
	Config      *viper.Viper
	address     string
	serviceName string
)

func init() {
	Config = config.GetConfig("relation")
	address = Config.GetString("server.address")
	serviceName = Config.GetString("server.serviceName")
	InitRelation()
}

func InitRelation() {
	c, err := relationservice.NewClient(
		serviceName,
		client.WithHostPorts(address),
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
